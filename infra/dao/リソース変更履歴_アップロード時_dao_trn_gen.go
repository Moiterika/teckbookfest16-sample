// Code generated by xlsx2go.go; DO NOT EDIT.

package dao

import (
	"database/sql"
	"fmt"
	a "github.com/Moiterika/a"
	xerrors "golang.org/x/xerrors"
	"strconv"
	"strings"
)

type daoTrnリソース変更履歴アップロード時 struct {
	dm        *DaoTrnManager
	trn       *sql.Tx
	WbForInit Wbリソース変更履歴アップロード時
}

func (d *daoTrnリソース変更履歴アップロード時) init() (err error) {
	d.dm.dtリソース変更履歴アップロード時, err = d.SelectW(d.WbForInit)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	d.dm.mapIDvsDrリソース変更履歴アップロード時 = a.ToMap(d.dm.dtリソース変更履歴アップロード時, func(e *Dtoリソース変更履歴アップロード時) Id {
		return e.FldID
	})
	return
}
func (d *daoTrnリソース変更履歴アップロード時) Reset() {
	list := make([]*Dtoリソース変更履歴アップロード時, 0, len(d.dm.dtリソース変更履歴アップロード時))
	for _, dr := range d.dm.dtリソース変更履歴アップロード時 {
		if dr.rowState == Deleted {
			dr.rowState = Detached
			continue
		} else {
			dr.rowState = UnChanged
			list = append(list, dr)
		}
	}
	d.dm.dtリソース変更履歴アップロード時 = list
}
func (d daoTrnリソース変更履歴アップロード時) Dt() ([]*Dtoリソース変更履歴アップロード時, error) {
	if len(d.dm.dtリソース変更履歴アップロード時) == 0 {
		err := d.init()
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return nil, err
		}
	}
	return d.dm.dtリソース変更履歴アップロード時, nil
}
func (d daoTrnリソース変更履歴アップロード時) GetBy(id Id) (dr *Dtoリソース変更履歴アップロード時, err error) {
	if len(d.dm.mapIDvsDrリソース変更履歴アップロード時) == 0 {
		err = d.init()
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
	}
	var ok bool
	dr, ok = d.dm.mapIDvsDrリソース変更履歴アップロード時[id]
	if !ok {
		err = xerrors.Errorf("リソース変更履歴_アップロード時が見つかりません。ID=%d: %w", id, NotFoundError)
		return
	}
	return
}
func (d daoTrnリソース変更履歴アップロード時) SelectAll() ([]*Dtoリソース変更履歴アップロード時, error) {
	sql := fmt.Sprintf(sqlSelectリソース変更履歴アップロード時, "")
	rows, err := d.trn.Query(sql)
	defer rows.Close()
	if err != nil {
		return nil, xerrors.Errorf("sql=%s: %w", sql, err)
	}
	var dt []*Dtoリソース変更履歴アップロード時
	for rows.Next() {
		var dr Dtoリソース変更履歴アップロード時
		err = rows.Scan(&dr.FldID, &dr.Fldアップロード履歴ID)
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}
		dr.rowState = Detached
		dr.Ub = NewUbリソース変更履歴アップロード時()
		dt = append(dt, &dr)
	}
	return dt, rows.Err()
}
func (d daoTrnリソース変更履歴アップロード時) SelectW(wb Wbリソース変更履歴アップロード時) ([]*Dtoリソース変更履歴アップロード時, error) {
	where := wb.build()
	prms, exists := where.Params()
	if exists {
		sql := fmt.Sprintf(sqlSelectリソース変更履歴アップロード時, where.String())
		rows, err := d.trn.Query(sql, prms...)
		defer rows.Close()
		if err != nil {
			return nil, xerrors.Errorf("sql=%s, args=%v: %w", sql, prms, err)
		}
		var dt []*Dtoリソース変更履歴アップロード時
		for rows.Next() {
			var dr Dtoリソース変更履歴アップロード時
			err = rows.Scan(&dr.FldID, &dr.Fldアップロード履歴ID)
			if err != nil {
				return nil, xerrors.Errorf(": %w", err)
			}
			dr.rowState = Detached
			dr.Ub = NewUbリソース変更履歴アップロード時()
			dt = append(dt, &dr)
		}
		return dt, rows.Err()
	} else {
		return d.SelectAll()
	}
}
func (d daoTrnリソース変更履歴アップロード時) Count() (cnt int64, err error) {
	return d.CountW(NewWbリソース変更履歴アップロード時())
}
func (d daoTrnリソース変更履歴アップロード時) CountW(wb Wbリソース変更履歴アップロード時) (cnt int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	if exists {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelectリソース変更履歴アップロード時ForAggregation, "count(\"ID\")", where.String()), prms...).Scan(&cnt)
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
		return
	} else {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelectリソース変更履歴アップロード時ForAggregation, "count(\"ID\")", "")).Scan(&cnt)
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
		return
	}
}

// Min はリソース変更履歴_アップロード時のfld最小値を返します。
func (d daoTrnリソース変更履歴アップロード時) Min(fld fldリソース変更履歴アップロード時) (min int64, err error) {
	return d.MinW(fld, NewWbリソース変更履歴アップロード時())
}
func (d daoTrnリソース変更履歴アップロード時) MinW(fld fldリソース変更履歴アップロード時, wb Wbリソース変更履歴アップロード時) (min int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	var x sql.NullInt64
	if exists {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelectリソース変更履歴アップロード時ForAggregation, fmt.Sprintf("min(%s)", strconv.Quote(string(fld))), where.String()), prms...).Scan(&x)
	} else {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelectリソース変更履歴アップロード時ForAggregation, fmt.Sprintf("min(%s)", strconv.Quote(string(fld))), where.String())).Scan(&x)
	}
	if err != nil {
		err = xerrors.Errorf(": %w", err)
		return
	}
	if !x.Valid {
		err = xerrors.Errorf(": %w", NotFoundError)
		return
	}
	min = x.Int64
	return
}

// Max はリソース変更履歴_アップロード時のfld最大値を返します。
func (d daoTrnリソース変更履歴アップロード時) Max(fld fldリソース変更履歴アップロード時) (max int64, err error) {
	return d.MaxW(fld, NewWbリソース変更履歴アップロード時())
}
func (d daoTrnリソース変更履歴アップロード時) MaxW(fld fldリソース変更履歴アップロード時, wb Wbリソース変更履歴アップロード時) (max int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	var x sql.NullInt64
	if exists {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelectリソース変更履歴アップロード時ForAggregation, fmt.Sprintf("max(%s)", strconv.Quote(string(fld))), where.String()), prms...).Scan(&x)
	} else {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelectリソース変更履歴アップロード時ForAggregation, fmt.Sprintf("max(%s)", strconv.Quote(string(fld))), where.String())).Scan(&x)
	}
	if err != nil {
		err = xerrors.Errorf(": %w", err)
		return
	}
	if !x.Valid {
		err = xerrors.Errorf(": %w", NotFoundError)
		return
	}
	max = x.Int64
	return
}
func (d daoTrnリソース変更履歴アップロード時) Insert(dr *Dtoリソース変更履歴アップロード時) (err error) {
	_, err = d.trn.Exec(sqlInsertリソース変更履歴アップロード時, dr.FldID, dr.Fldアップロード履歴ID)
	if err != nil {
		err = xerrors.Errorf(": %w", err)
		return
	}
	dr.rowState = Added
	return
}
func (d daoTrnリソース変更履歴アップロード時) MultiInsert(dt []*Dtoリソース変更履歴アップロード時) (err error) {
	cs := a.Chunk(dt, 1000)
	for _, c := range cs {
		vals := make([]string, len(c))
		args := make([]interface{}, len(c)*2)
		for i, dr := range c {
			vals[i] = fmt.Sprintf(sqlValue2, 2*i+1, 2*i+2)
			args[2*i] = dr.FldID
			args[2*i+1] = dr.Fldアップロード履歴ID
			dr.rowState = Added
		}
		_, err = d.trn.Exec(fmt.Sprintf(sqlMultiInsertリソース変更履歴アップロード時, strings.Join(vals, ",")), args...)
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
	}
	return
}
func (d daoTrnリソース変更履歴アップロード時) UpdateBy(dr *Dtoリソース変更履歴アップロード時) (cnt int64, err error) {
	if dr.Ub.Count() == 0 {
		dr.rowState = UnChanged
		return
	}
	s, w, execArgs := dr.Ub.build(newWbリソース変更履歴アップロード時WithPrimaryKeys(dr.FldID))
	sql := fmt.Sprintf(sqlUpdateリソース変更履歴アップロード時, s, w)
	result, err := d.trn.Exec(sql, execArgs...)
	if err != nil {
		err = xerrors.Errorf("sql=%s, args=%v: %w", sql, execArgs, err)
		return
	}
	cnt, err = result.RowsAffected()
	if err != nil {
		err = xerrors.Errorf(": %w", err)
		return
	}
	dr.rowState = Modified
	return
}
func (d daoTrnリソース変更履歴アップロード時) UpdateW(ub *ubリソース変更履歴アップロード時, wb Wbリソース変更履歴アップロード時) (cnt int64, err error) {
	s, w, execArgs := ub.build(wb)
	sql := fmt.Sprintf(sqlUpdateリソース変更履歴アップロード時, s, w)
	result, err := d.trn.Exec(sql, execArgs...)
	if err != nil {
		err = xerrors.Errorf("sql=%s, args=%v: %w", sql, execArgs, err)
		return
	}
	cnt, err = result.RowsAffected()
	if err != nil {
		err = xerrors.Errorf(": %w", err)
		return
	}
	return
}
func (d daoTrnリソース変更履歴アップロード時) DeleteBy(dr *Dtoリソース変更履歴アップロード時) (cnt int64, err error) {
	where := newWbリソース変更履歴アップロード時WithPrimaryKeys(dr.FldID).build()
	prms, exists := where.Params()
	if !exists {
		err = xerrors.Errorf("主キーがありません。: %#v", *dr)
		return
	}
	sql := fmt.Sprintf(sqlDeleteリソース変更履歴アップロード時, where.String())
	result, errExec := d.trn.Exec(sql, prms...)
	if errExec != nil {
		err = xerrors.Errorf("sql=%s, args=%v: %w", sql, prms, errExec)
		return
	}
	cnt, err = result.RowsAffected()
	if err != nil {
		err = xerrors.Errorf(": %w", err)
		return
	}
	dr.rowState = Deleted
	return
}
func (d daoTrnリソース変更履歴アップロード時) DeleteW(wb Wbリソース変更履歴アップロード時) (cnt int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	if exists {
		sql := fmt.Sprintf(sqlDeleteリソース変更履歴アップロード時, where.String())
		result, errExec := d.trn.Exec(sql, prms...)
		if errExec != nil {
			err = xerrors.Errorf("sql=%s, args=%v: %w", sql, prms, errExec)
			return
		}
		cnt, err = result.RowsAffected()
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
	} else {
		result, errExec := d.trn.Exec(sqlTruncateリソース変更履歴アップロード時)
		if errExec != nil {
			err = xerrors.Errorf("sql=%s: %w", sqlTruncateリソース変更履歴アップロード時, errExec)
			return
		}
		cnt, err = result.RowsAffected()
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
	}
	return
}
