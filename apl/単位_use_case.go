package apl

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"techbookfest16-sample/domain"
	"techbookfest16-sample/domain/objects"
	"techbookfest16-sample/domain/types"
	"techbookfest16-sample/infra"
)

func (mhs *myHttpServer) UseCase単位(w http.ResponseWriter, r *http.Request) {
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
			list, err := rm.NewRep単位().List()
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

		// 単位コードで1件取得して返却
		単位コード := mhs.GetCode(r)
		e, err := rm.NewRep単位().GetBy(types.Code単位(単位コード))
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

		var 単位 objects.Ent単位
		err = json.NewDecoder(r.Body).Decode(&単位)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		s := domain.NewSrv単位登録(infra.NewRepManagerWithTrn(trn))
		err = s.Exec登録(1, &単位) // アップロード履歴は今回の範囲外なので常に1とする
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
