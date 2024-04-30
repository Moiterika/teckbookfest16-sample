package apl

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"techbookfest16-sample/domain"
	"techbookfest16-sample/domain/types"
	"techbookfest16-sample/infra"
	"techbookfest16-sample/infra/dao"
	"time"
)

func (mhs *myHttpServer) UseCase仕入(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL.Path, r.URL.RawPath)
	if r.Method == http.MethodGet {
		// TODO Transactionではなく、Dbを渡して参照だけすること（毎回、トランザクション貼りたくない。以前のコミットではできてたので、それを参照）
		trn, err := mhs.defaultDb.Begin()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer trn.Commit()

		wb受払 := dao.NewWb受払()
		wb品目 := dao.NewWb品目()
		// TODO wb仕入品を入れ子にできるかは後で確かめる

		if r.URL.Path == "" {
			// クエリパラメータ年月を取得
			var 計上年月 time.Time
			年月 := r.URL.Query().Get("年月")
			if 年月 != "" {
				計上年月, err = time.Parse("200601", 年月)
				if err != nil {
					http.Error(w, "年月はyyyyMMの書式にしてください。", http.StatusBadRequest)
					return
				}
				// 検索条件にセット
				wb受払.And(dao.Tbl受払().Fld計上月(), dao.OpEqu, 計上年月)
				wb品目.Exists(dao.NewEb品目join受払().And(dao.Tbl受払().Fld計上月(), dao.OpEqu, 計上年月))
			}
			q := infra.NewQryTrn受払仕入(
				infra.NewRepManagerWithTrn(
					trn,
					infra.Wb受払(wb受払),
					infra.Wb品目(wb品目),
					infra.Wb製造品(dao.NewNothingWb品目製造品()),
				),
			)
			list, err := q.List()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			jsonb, err := json.Marshal(list)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(jsonb)
			return
		}

		// 受払Noで1件取得して返却
		受払No, err := mhs.GetNo(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// 検索条件にセット
		wb受払.And(dao.Tbl受払().FldNo(), dao.OpEqu, 受払No)
		wb品目.Exists(dao.NewEb品目join受払().And(dao.Tbl受払().FldNo(), dao.OpEqu, 受払No))
		// TODO wb仕入品を入れ子にできるかは後で確かめる

		q := infra.NewQryTrn受払仕入(
			infra.NewRepManagerWithTrn(
				trn,
				infra.Wb受払(wb受払),
				infra.Wb品目(wb品目),
			),
		)
		e, err := q.GetBy(受払No)
		if err != nil {
			if errors.Is(err, types.ErrNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			} else {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
		jsonb, err := json.Marshal(e)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonb)
	} else if r.Method == http.MethodPost {
		trn, err := mhs.defaultDb.Begin()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer func() {
			if err != nil {
				trn.Rollback()
				return
			}
			trn.Commit()
		}()
		var 仕入 Dto仕入
		err = json.NewDecoder(r.Body).Decode(&仕入)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// ここでは1件だけ登録する仕様なので、DBの検索条件をセット
		// TODO 別のユースケースとして、複数の仕入をバッチ登録する場合、DBの検索条件をセットせず、品目リポジトリに全件取得した方が効率が良い
		wb品目 := dao.NewWb品目().And(dao.Tbl品目().Fldコード(), dao.OpEqu, 仕入.Get品目コード)
		wb仕入品 := dao.NewWb品目仕入品().Exists(dao.NewEb品目仕入品join品目().And(dao.Tbl品目().Fldコード(), dao.OpEqu, 仕入.Get品目コード))

		rm := infra.NewRepManagerWithTrn(trn,
			infra.Wb品目(wb品目),
			infra.Wb仕入品(wb仕入品),
			infra.Wb製造品(dao.NewNothingWb品目製造品()))
		s := domain.NewSrv仕入登録(rm, infra.NewCmdTrn受払(rm))
		err = s.Exec登録(
			仕入.Get登録日時,  // 登録日時 time.Time,
			仕入.Get計上月,   // 計上月 time.Time,
			仕入.Get受払区分,  // 受払区分 objects.Enum受払区分,
			仕入.Get赤伝フラグ, // 赤伝フラグ bool,
			仕入.Get品目コード, // 品目コード types.Code品目,
			仕入.Get基準数量,  // 基準数量 types.Inventory,
			仕入.Get仕入数量,  // 仕入数量 types.Quantity,
			仕入.Get仕入金額,  // 仕入金額 types.Amount,
			仕入.Get仕入単価,  // 仕入単価 types.Price,
		)
		if err != nil {
			if errors.Is(err, types.ErrArg) {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, fmt.Sprintf("HTTP method=%sは許可されていません。", r.Method), http.StatusMethodNotAllowed)
	}
}
