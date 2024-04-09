package apl

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"techbookfest16-sample/domain"
	"techbookfest16-sample/domain/types"
	"techbookfest16-sample/infra"
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

		q := infra.NewQryTrn受払仕入(infra.NewRepManagerWithTrn(trn))

		// 全件返却
		// TODO 件数が多すぎる可能性があるので、ページネーションで1リクエストあたりの返却件数を区切るか、計上年月を必須にして返却件数を制限すべき
		// TODO その代わりに全件返却の仕組みもどこかに必要ではある
		if r.URL.Path == "" {
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
		return
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

		s := domain.NewSrv仕入登録(infra.NewCmdTrn受払(infra.NewRepManagerWithTrn(trn)))
		仕入数量, err := 仕入.仕入数量()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = s.Exec登録(
			仕入.Get登録日時,  // 登録日時 time.Time,
			仕入.Get計上月,   // 計上月 time.Time,
			仕入.Get受払区分,  // 受払区分 objects.Enum受払区分,
			仕入.Get赤伝フラグ, // 赤伝フラグ bool,
			仕入.Get品目コード, // 品目コード types.Code品目,
			仕入.基準数量(),   // 基準数量 types.Inventory,
			仕入数量,        // 仕入数量 types.Quantity,
			仕入.仕入金額(),   // 仕入金額 types.Amount,
			仕入.仕入単価(),   // 仕入単価 types.Price,
		)
		if err != nil {
			if errors.Is(err, types.ErrArg) {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	http.Error(w, fmt.Sprintf("HTTP method=%sは許可されていません。", r.Method), http.StatusMethodNotAllowed)
}
