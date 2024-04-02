// Code generated by xlsx2go.go; DO NOT EDIT.

package dao

import "database/sql"

type Dtogorpmigrations struct {
	Fldid         string
	Fldapplied_at sql.NullTime

	rowState DataRowState
	Ub       *ubgorpmigrations `json:"-"`
}

func (d Dtogorpmigrations) TableName() string {
	return "gorp_migrations"
}
func (d Dtogorpmigrations) RowState() DataRowState {
	return d.rowState
}

// Import はDtogorpmigrations型に主キー以外を上書きする。
func (d *Dtogorpmigrations) Import(applied_at sql.NullTime) {
	// 項目がすべて一致していたら、何もしない
	if d.Fldapplied_at == applied_at {
		return
	}
	if d.Fldapplied_at != applied_at {
		d.Fldapplied_at = applied_at
		d.Ub.Set(Tblgorpmigrations().Fldapplied_at(), applied_at)
	}

}
