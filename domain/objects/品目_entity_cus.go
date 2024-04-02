package objects

import "golang.org/x/xerrors"

func (e *Ent品目) Validate() error {
	if e.Get仕入品 == nil && e.Get製造品 == nil {
		err := xerrors.Errorf("仕入品・製造品は必須項目です。どちらか、あるいは、両方を設定してください。品目コード=%s", e.Getコード)
		return err
	}
	return nil
}
