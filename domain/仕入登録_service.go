package domain

import (
	"fmt"
	"techbookfest16-sample/domain/objects"
	"techbookfest16-sample/domain/types"
	"time"
)

type Srv仕入登録 struct {
	rm  objects.RepManager
	cmd objects.Cmd受払仕入
}

func NewSrv仕入登録(rm objects.RepManager, cmd objects.Cmd受払仕入) *Srv仕入登録 {
	return &Srv仕入登録{
		rm:  rm,
		cmd: cmd,
	}
}

// Exec登録 は仕入を新規登録します。
func (s *Srv仕入登録) Exec登録(
	//アップロード履歴 types.No,
	登録日時 time.Time,
	計上月 time.Time,
	受払区分 objects.Enum受払区分,
	赤伝フラグ bool,
	品目コード types.Code品目,
	基準数量 types.Inventory,
	仕入数量 types.Quantity,
	仕入金額 types.Amount,
	仕入単価 types.Price,
) error {
	rep品目 := s.rm.NewRep品目()
	仕入品, notFound := rep品目.Get仕入品By(品目コード)
	if notFound != nil {
		return notFound
	}

	if 基準数量.Unit() != 仕入品.Get基準単位.Getコード {
		return fmt.Errorf("仕入時の基準数量の単位コードは%sであるべきです（品目コード=%s、単位コード=%s）。", 仕入品.Get基準単位.Getコード, 品目コード, 基準数量.Unit())
	}

	// 到達しない分岐
	// if 仕入数量.Unit() != 仕入単価.PerUnit() {
	// 	return fmt.Errorf("仕入数量と仕入単価の単位が一致しません（品目コード=%s、単位コード=%s）。", 品目コード, 仕入数量.Unit())
	// }

	e仕入 := &objects.Ent受払仕入{
		Ent受払: &objects.Ent受払{
			//GetNo:    0,
			Get登録日時:  登録日時,
			Get計上月:   計上月,
			Get受払区分:  受払区分,
			Get赤伝フラグ: 赤伝フラグ,
			Get品目:    仕入品.Ent品目,
			Get基準数量:  基準数量,
		},
		Get仕入数量: 仕入数量,
		Get仕入金額: 仕入金額,
		Get仕入単価: 仕入単価,
	}

	err := e仕入.Validate()
	if err != nil {
		return fmt.Errorf("validate error: %w, %w", err, types.ErrArg)
	}
	err = s.cmd.Entry(e仕入)
	if err != nil {
		return fmt.Errorf("entry error: %w, %w", err, types.ErrArg)
	}
	return nil
}
