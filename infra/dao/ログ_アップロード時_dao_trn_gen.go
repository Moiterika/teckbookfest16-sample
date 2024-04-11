// Code generated by xlsx2go.go; DO NOT EDIT.

package dao

import (
	"database/sql"
	"fmt"
	a "github.com/Moiterika/a"
	xerrors "golang.org/x/xerrors"
	"strconv"
	"strings"
	types "techbookfest16-sample/domain/types"
)

type daoTrnログアップロード時 struct {
	dm        *DaoTrnManager
	trn       *sql.Tx
	WbForInit Wbログアップロード時
}

func (d *daoTrnログアップロード時) init() (err error) {
	d.dm.dtログアップロード時, err = d.SelectW(d.WbForInit)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	d.dm.mapIDvsDrログアップロード時 = a.ToMap(d.dm.dtログアップロード時, func(e *Dtoログアップロード時) Id {
		return e.FldNo
	})
	return
}
func (d *daoTrnログアップロード時) Reset() {
	list := make([]*Dtoログアップロード時, 0, len(d.dm.dtログアップロード時))
	for _, dr := range d.dm.dtログアップロード時 {
		if dr.rowState == Deleted {
			dr.rowState = Detached
			continue
		} else {
			dr.rowState = UnChanged
			list = append(list, dr)
		}
	}
	d.dm.dtログアップロード時 = list
}
func (d daoTrnログアップロード時) Dt() ([]*Dtoログアップロード時, error) {
	if len(d.dm.dtログアップロード時) == 0 {
		err := d.init()
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return nil, err
		}
	}
	return d.dm.dtログアップロード時, nil
}
func (d daoTrnログアップロード時) GetBy(id Id) (dr *Dtoログアップロード時, err error) {
	if len(d.dm.mapIDvsDrログアップロード時) == 0 {
		err = d.init()
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
	}
	var ok bool
	dr, ok = d.dm.mapIDvsDrログアップロード時[id]
	if !ok {
		err = xerrors.Errorf("ログ_アップロード時が見つかりません。No=%d: %w", id, types.ErrNotFound)
		return
	}
	return
}
func (d daoTrnログアップロード時) SelectAll() ([]*Dtoログアップロード時, error) {
	sql := fmt.Sprintf(sqlSelectログアップロード時, "")
	rows, err := d.trn.Query(sql)
	if err != nil {
		return nil, xerrors.Errorf("sql=%s: %w", sql, err)
	}
	defer rows.Close()
	var dt []*Dtoログアップロード時
	for rows.Next() {
		var dr Dtoログアップロード時
		err = rows.Scan(&dr.FldNo, &dr.Fldアップロード履歴ID)
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}
		dr.rowState = Detached
		dr.Ub = NewUbログアップロード時()
		dt = append(dt, &dr)
	}
	return dt, rows.Err()
}
func (d daoTrnログアップロード時) SelectW(wb Wbログアップロード時) ([]*Dtoログアップロード時, error) {
	where := wb.build()
	prms, exists := where.Params()
	if exists {
		sql := fmt.Sprintf(sqlSelectログアップロード時, where.String())
		rows, err := d.trn.Query(sql, prms...)
		if err != nil {
			return nil, xerrors.Errorf("sql=%s, args=%v: %w", sql, prms, err)
		}
		defer rows.Close()
		var dt []*Dtoログアップロード時
		for rows.Next() {
			var dr Dtoログアップロード時
			err = rows.Scan(&dr.FldNo, &dr.Fldアップロード履歴ID)
			if err != nil {
				return nil, xerrors.Errorf(": %w", err)
			}
			dr.rowState = Detached
			dr.Ub = NewUbログアップロード時()
			dt = append(dt, &dr)
		}
		return dt, rows.Err()
	} else {
		return d.SelectAll()
	}
}
func (d daoTrnログアップロード時) Count() (cnt int64, err error) {
	return d.CountW(NewWbログアップロード時())
}
func (d daoTrnログアップロード時) CountW(wb Wbログアップロード時) (cnt int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	if exists {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelectログアップロード時ForAggregation, "count(\"No\")", where.String()), prms...).Scan(&cnt)
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
		return
	} else {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelectログアップロード時ForAggregation, "count(\"No\")", "")).Scan(&cnt)
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
		return
	}
}

// Min はログ_アップロード時のfld最小値を返します。
func (d daoTrnログアップロード時) Min(fld fldログアップロード時) (min int64, err error) {
	return d.MinW(fld, NewWbログアップロード時())
}
func (d daoTrnログアップロード時) MinW(fld fldログアップロード時, wb Wbログアップロード時) (min int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	var x sql.NullInt64
	if exists {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelectログアップロード時ForAggregation, fmt.Sprintf("min(%s)", strconv.Quote(string(fld))), where.String()), prms...).Scan(&x)
	} else {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelectログアップロード時ForAggregation, fmt.Sprintf("min(%s)", strconv.Quote(string(fld))), where.String())).Scan(&x)
	}
	if err != nil {
		err = xerrors.Errorf(": %w", err)
		return
	}
	if !x.Valid {
		err = xerrors.Errorf(": %w", types.ErrNotFound)
		return
	}
	min = x.Int64
	return
}

// Max はログ_アップロード時のfld最大値を返します。
func (d daoTrnログアップロード時) Max(fld fldログアップロード時) (max int64, err error) {
	return d.MaxW(fld, NewWbログアップロード時())
}
func (d daoTrnログアップロード時) MaxW(fld fldログアップロード時, wb Wbログアップロード時) (max int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	var x sql.NullInt64
	if exists {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelectログアップロード時ForAggregation, fmt.Sprintf("max(%s)", strconv.Quote(string(fld))), where.String()), prms...).Scan(&x)
	} else {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelectログアップロード時ForAggregation, fmt.Sprintf("max(%s)", strconv.Quote(string(fld))), where.String())).Scan(&x)
	}
	if err != nil {
		err = xerrors.Errorf(": %w", err)
		return
	}
	if !x.Valid {
		err = xerrors.Errorf(": %w", types.ErrNotFound)
		return
	}
	max = x.Int64
	return
}
func (d daoTrnログアップロード時) Insert(dr *Dtoログアップロード時) (err error) {
	_, err = d.trn.Exec(sqlInsertログアップロード時, dr.FldNo, dr.Fldアップロード履歴ID)
	if err != nil {
		err = xerrors.Errorf(": %w", err)
		return
	}
	dr.rowState = Added
	d.dm.dtログアップロード時 = append(d.dm.dtログアップロード時, dr)
	d.dm.mapIDvsDrログアップロード時[dr.FldNo] = dr
	return
}
func (d daoTrnログアップロード時) MultiInsert(dt []*Dtoログアップロード時) (err error) {
	cs := a.Chunk(dt, 1000)
	for _, c := range cs {
		vals := make([]string, len(c))
		args := make([]interface{}, len(c)*2)
		for i, dr := range c {
			vals[i] = fmt.Sprintf(sqlValue2, 2*i+1, 2*i+2)
			args[2*i] = dr.FldNo
			args[2*i+1] = dr.Fldアップロード履歴ID
			dr.rowState = Added
		}
		_, err = d.trn.Exec(fmt.Sprintf(sqlMultiInsertログアップロード時, strings.Join(vals, ",")), args...)
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
	}
	return
}
func (d daoTrnログアップロード時) UpdateBy(dr *Dtoログアップロード時) (cnt int64, err error) {
	if dr.Ub.Count() == 0 {
		dr.rowState = UnChanged
		return
	}
	s, w, execArgs := dr.Ub.build(newWbログアップロード時WithPrimaryKeys(dr.FldNo))
	sql := fmt.Sprintf(sqlUpdateログアップロード時, s, w)
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
func (d daoTrnログアップロード時) UpdateW(ub *ubログアップロード時, wb Wbログアップロード時) (cnt int64, err error) {
	s, w, execArgs := ub.build(wb)
	sql := fmt.Sprintf(sqlUpdateログアップロード時, s, w)
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
func (d daoTrnログアップロード時) DeleteBy(dr *Dtoログアップロード時) (cnt int64, err error) {
	where := newWbログアップロード時WithPrimaryKeys(dr.FldNo).build()
	prms, exists := where.Params()
	if !exists {
		err = xerrors.Errorf("主キーがありません。: %#v", *dr)
		return
	}
	sql := fmt.Sprintf(sqlDeleteログアップロード時, where.String())
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
func (d daoTrnログアップロード時) DeleteW(wb Wbログアップロード時) (cnt int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	if exists {
		sql := fmt.Sprintf(sqlDeleteログアップロード時, where.String())
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
		result, errExec := d.trn.Exec(sqlTruncateログアップロード時)
		if errExec != nil {
			err = xerrors.Errorf("sql=%s: %w", sqlTruncateログアップロード時, errExec)
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
