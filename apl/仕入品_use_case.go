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

func (mhs *myHttpServer) UseCase仕入品(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL.Path, r.URL.RawPath)
	if r.Method == http.MethodGet {
		// TODO Transactionではなく、Dbを渡して参照だけすること（毎回、トランザクション貼りたくない。1つ前のコミットではできてたので、それを参照）
		trn, err := mhs.defaultDb.Begin()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer trn.Commit()
		rm := infra.NewRepManagerWithTrn(trn)

		// 全件返却
		if r.URL.Path == "" {
			list, err := rm.NewRep品目().Get仕入品一覧()
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

		// コードで1件取得して返却
		品目コード := mhs.GetCode(r)
		e, err := rm.NewRep品目().Get仕入品By(types.Code品目(品目コード))
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
	} else if r.Method == http.MethodPatch {
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

		var dr仕入品 Dto仕入品
		err = json.NewDecoder(r.Body).Decode(&dr仕入品)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		s := domain.NewSrv仕入品登録(infra.NewRepManagerWithTrn(trn))
		err = s.Exec登録(
			1, // アップロード履歴は今回の範囲外なので常に1とする
			dr仕入品.Getコード,
			dr仕入品.Get名称,
			dr仕入品.Get基準単位コード,
			dr仕入品.Get生産用品目区分コード,
			dr仕入品.Get標準単価,
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
