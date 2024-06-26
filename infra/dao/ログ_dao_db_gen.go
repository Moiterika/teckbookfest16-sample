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

type daoDbログ struct {
	dm        *DaoDbManager
	db        *sql.DB
	WbForInit Wbログ
}

func (d *daoDbログ) init() (err error) {
	d.dm.dtログ, err = d.SelectW(d.WbForInit)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	d.dm.mapIDvsDrログ = a.ToMap(d.dm.dtログ, func(e *Dtoログ) Id {
		return e.FldNo
	})
	return
}
func (d *daoDbログ) Reset() {
	list := make([]*Dtoログ, 0, len(d.dm.dtログ))
	for _, dr := range d.dm.dtログ {
		if dr.rowState == Deleted {
			dr.rowState = Detached
			continue
		} else {
			dr.rowState = UnChanged
			list = append(list, dr)
		}
	}
	d.dm.dtログ = list
}
func (d daoDbログ) Dt() ([]*Dtoログ, error) {
	if len(d.dm.dtログ) == 0 {
		err := d.init()
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return nil, err
		}
	}
	return d.dm.dtログ, nil
}
func (d daoDbログ) GetBy(id Id) (dr *Dtoログ, err error) {
	if len(d.dm.mapIDvsDrログ) == 0 {
		err = d.init()
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
	}
	var ok bool
	dr, ok = d.dm.mapIDvsDrログ[id]
	if !ok {
		err = xerrors.Errorf("ログが見つかりません。No=%d: %w", id, types.ErrNotFound)
		return
	}
	return
}
func (d daoDbログ) SelectAll() ([]*Dtoログ, error) {
	sql := fmt.Sprintf(sqlSelectログ, "")
	rows, err := d.db.Query(sql)
	if err != nil {
		return nil, xerrors.Errorf("sql=%s: %w", sql, err)
	}
	defer rows.Close()
	var dt []*Dtoログ
	for rows.Next() {
		var dr Dtoログ
		err = rows.Scan(&dr.FldNo, &dr.Fld登録日時, &dr.Fld区分, &dr.Fld内容)
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}
		dr.rowState = Detached
		dr.Ub = NewUbログ()
		dt = append(dt, &dr)
	}
	return dt, rows.Err()
}
func (d daoDbログ) SelectW(wb Wbログ) ([]*Dtoログ, error) {
	where := wb.build()
	prms, exists := where.Params()
	if exists {
		sql := fmt.Sprintf(sqlSelectログ, where.String())
		rows, err := d.db.Query(sql, prms...)
		if err != nil {
			return nil, xerrors.Errorf("sql=%s, args=%v: %w", sql, prms, err)
		}
		defer rows.Close()
		var dt []*Dtoログ
		for rows.Next() {
			var dr Dtoログ
			err = rows.Scan(&dr.FldNo, &dr.Fld登録日時, &dr.Fld区分, &dr.Fld内容)
			if err != nil {
				return nil, xerrors.Errorf(": %w", err)
			}
			dr.rowState = Detached
			dr.Ub = NewUbログ()
			dt = append(dt, &dr)
		}
		return dt, rows.Err()
	} else {
		return d.SelectAll()
	}
}
func (d daoDbログ) Count() (cnt int64, err error) {
	return d.CountW(NewWbログ())
}
func (d daoDbログ) CountW(wb Wbログ) (cnt int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	if exists {
		err = d.db.QueryRow(fmt.Sprintf(sqlSelectログForAggregation, "count(\"No\")", where.String()), prms...).Scan(&cnt)
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
		return
	} else {
		err = d.db.QueryRow(fmt.Sprintf(sqlSelectログForAggregation, "count(\"No\")", "")).Scan(&cnt)
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
		return
	}
}

// Min はログのfld最小値を返します。
func (d daoDbログ) Min(fld fldログ) (min int64, err error) {
	return d.MinW(fld, NewWbログ())
}
func (d daoDbログ) MinW(fld fldログ, wb Wbログ) (min int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	var x sql.NullInt64
	if exists {
		err = d.db.QueryRow(fmt.Sprintf(sqlSelectログForAggregation, fmt.Sprintf("min(%s)", strconv.Quote(string(fld))), where.String()), prms...).Scan(&x)
	} else {
		err = d.db.QueryRow(fmt.Sprintf(sqlSelectログForAggregation, fmt.Sprintf("min(%s)", strconv.Quote(string(fld))), where.String())).Scan(&x)
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

// Max はログのfld最大値を返します。
func (d daoDbログ) Max(fld fldログ) (max int64, err error) {
	return d.MaxW(fld, NewWbログ())
}
func (d daoDbログ) MaxW(fld fldログ, wb Wbログ) (max int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	var x sql.NullInt64
	if exists {
		err = d.db.QueryRow(fmt.Sprintf(sqlSelectログForAggregation, fmt.Sprintf("max(%s)", strconv.Quote(string(fld))), where.String()), prms...).Scan(&x)
	} else {
		err = d.db.QueryRow(fmt.Sprintf(sqlSelectログForAggregation, fmt.Sprintf("max(%s)", strconv.Quote(string(fld))), where.String())).Scan(&x)
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
func (d daoDbログ) Insert(dr *Dtoログ) (err error) {
	err = d.db.QueryRow(sqlInsertログ, dr.Fld登録日時, dr.Fld区分, dr.Fld内容).Scan(&dr.FldNo)
	if err != nil {
		err = xerrors.Errorf(": %w", err)
		return
	}
	dr.rowState = Added
	d.dm.dtログ = append(d.dm.dtログ, dr)
	d.dm.mapIDvsDrログ[dr.FldNo] = dr
	return
}
func (d daoDbログ) MultiInsert(dt []*Dtoログ) (err error) {
	cs := a.Chunk(dt, 1000)
	for _, c := range cs {
		vals := make([]string, len(c))
		args := make([]interface{}, len(c)*3)
		for i, dr := range c {
			vals[i] = fmt.Sprintf(sqlValue3, 3*i+1, 3*i+2, 3*i+3)
			args[3*i] = dr.Fld登録日時
			args[3*i+1] = dr.Fld区分
			args[3*i+2] = dr.Fld内容
			dr.rowState = Added
		}
		_, err = d.db.Exec(fmt.Sprintf(sqlMultiInsertログ, strings.Join(vals, ",")), args...)
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
	}
	return
}
func (d daoDbログ) UpdateBy(dr *Dtoログ) (cnt int64, err error) {
	if dr.Ub.Count() == 0 {
		dr.rowState = UnChanged
		return
	}
	s, w, execArgs := dr.Ub.build(newWbログWithPrimaryKeys(dr.FldNo))
	sql := fmt.Sprintf(sqlUpdateログ, s, w)
	result, err := d.db.Exec(sql, execArgs...)
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
func (d daoDbログ) UpdateW(ub *ubログ, wb Wbログ) (cnt int64, err error) {
	s, w, execArgs := ub.build(wb)
	sql := fmt.Sprintf(sqlUpdateログ, s, w)
	result, err := d.db.Exec(sql, execArgs...)
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
func (d daoDbログ) DeleteBy(dr *Dtoログ) (cnt int64, err error) {
	where := newWbログWithPrimaryKeys(dr.FldNo).build()
	prms, exists := where.Params()
	if !exists {
		err = xerrors.Errorf("主キーがありません。: %#v", *dr)
		return
	}
	sql := fmt.Sprintf(sqlDeleteログ, where.String())
	result, errExec := d.db.Exec(sql, prms...)
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
func (d daoDbログ) DeleteW(wb Wbログ) (cnt int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	if exists {
		sql := fmt.Sprintf(sqlDeleteログ, where.String())
		result, errExec := d.db.Exec(sql, prms...)
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
		result, errExec := d.db.Exec(sqlTruncateログ)
		if errExec != nil {
			err = xerrors.Errorf("sql=%s: %w", sqlTruncateログ, errExec)
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
