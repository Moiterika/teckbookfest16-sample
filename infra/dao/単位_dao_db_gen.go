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

type daoDb単位 struct {
	dm        *DaoDbManager
	db        *sql.DB
	WbForInit Wb単位
}

func (d *daoDb単位) init() (err error) {
	d.dm.dt単位, err = d.SelectW(d.WbForInit)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	d.dm.mapIDvsDr単位 = a.ToMap(d.dm.dt単位, func(e *Dto単位) Id {
		return e.FldID
	})
	d.dm.mapコードvsDr単位 = a.ToMap(d.dm.dt単位, func(e *Dto単位) types.Code単位 {
		return e.Fldコード
	})
	return
}
func (d *daoDb単位) Reset() {
	list := make([]*Dto単位, 0, len(d.dm.dt単位))
	for _, dr := range d.dm.dt単位 {
		if dr.rowState == Deleted {
			dr.rowState = Detached
			continue
		} else {
			dr.rowState = UnChanged
			list = append(list, dr)
		}
	}
	d.dm.dt単位 = list
	d.dm.mapコードvsDr単位 = a.ToMap(d.dm.dt単位, func(e *Dto単位) types.Code単位 {
		return e.Fldコード
	})
}
func (d daoDb単位) Dt() ([]*Dto単位, error) {
	if len(d.dm.dt単位) == 0 {
		err := d.init()
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return nil, err
		}
	}
	return d.dm.dt単位, nil
}
func (d daoDb単位) GetBy(id Id) (dr *Dto単位, err error) {
	if len(d.dm.mapIDvsDr単位) == 0 {
		err = d.init()
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
	}
	var ok bool
	dr, ok = d.dm.mapIDvsDr単位[id]
	if !ok {
		err = xerrors.Errorf("単位が見つかりません。ID=%d: %w", id, NotFoundError)
		return
	}
	return
}
func (d daoDb単位) GetByCode(コード types.Code単位) (dr *Dto単位, err error) {
	if len(d.dm.mapコードvsDr単位) == 0 {
		err = d.init()
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
	}
	var ok bool
	dr, ok = d.dm.mapコードvsDr単位[コード]
	if !ok {
		err = xerrors.Errorf("単位が見つかりません。コード=%s: %w", コード, NotFoundError)
		return
	}
	return
}
func (d daoDb単位) SelectAll() ([]*Dto単位, error) {
	sql := fmt.Sprintf(sqlSelect単位, "")
	rows, err := d.db.Query(sql)
	defer rows.Close()
	if err != nil {
		return nil, xerrors.Errorf("sql=%s: %w", sql, err)
	}
	var dt []*Dto単位
	for rows.Next() {
		var dr Dto単位
		err = rows.Scan(&dr.FldID, &dr.Fldコード, &dr.Fld名称)
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}
		dr.rowState = Detached
		dr.Ub = NewUb単位()
		dt = append(dt, &dr)
	}
	return dt, rows.Err()
}
func (d daoDb単位) SelectW(wb Wb単位) ([]*Dto単位, error) {
	where := wb.build()
	prms, exists := where.Params()
	if exists {
		sql := fmt.Sprintf(sqlSelect単位, where.String())
		rows, err := d.db.Query(sql, prms...)
		defer rows.Close()
		if err != nil {
			return nil, xerrors.Errorf("sql=%s, args=%v: %w", sql, prms, err)
		}
		var dt []*Dto単位
		for rows.Next() {
			var dr Dto単位
			err = rows.Scan(&dr.FldID, &dr.Fldコード, &dr.Fld名称)
			if err != nil {
				return nil, xerrors.Errorf(": %w", err)
			}
			dr.rowState = Detached
			dr.Ub = NewUb単位()
			dt = append(dt, &dr)
		}
		return dt, rows.Err()
	} else {
		return d.SelectAll()
	}
}
func (d daoDb単位) Count() (cnt int64, err error) {
	return d.CountW(NewWb単位())
}
func (d daoDb単位) CountW(wb Wb単位) (cnt int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	if exists {
		err = d.db.QueryRow(fmt.Sprintf(sqlSelect単位ForAggregation, "count(\"ID\")", where.String()), prms...).Scan(&cnt)
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
		return
	} else {
		err = d.db.QueryRow(fmt.Sprintf(sqlSelect単位ForAggregation, "count(\"ID\")", "")).Scan(&cnt)
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
		return
	}
}

// Min は単位のfld最小値を返します。
func (d daoDb単位) Min(fld fld単位) (min int64, err error) {
	return d.MinW(fld, NewWb単位())
}
func (d daoDb単位) MinW(fld fld単位, wb Wb単位) (min int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	var x sql.NullInt64
	if exists {
		err = d.db.QueryRow(fmt.Sprintf(sqlSelect単位ForAggregation, fmt.Sprintf("min(%s)", strconv.Quote(string(fld))), where.String()), prms...).Scan(&x)
	} else {
		err = d.db.QueryRow(fmt.Sprintf(sqlSelect単位ForAggregation, fmt.Sprintf("min(%s)", strconv.Quote(string(fld))), where.String())).Scan(&x)
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

// Max は単位のfld最大値を返します。
func (d daoDb単位) Max(fld fld単位) (max int64, err error) {
	return d.MaxW(fld, NewWb単位())
}
func (d daoDb単位) MaxW(fld fld単位, wb Wb単位) (max int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	var x sql.NullInt64
	if exists {
		err = d.db.QueryRow(fmt.Sprintf(sqlSelect単位ForAggregation, fmt.Sprintf("max(%s)", strconv.Quote(string(fld))), where.String()), prms...).Scan(&x)
	} else {
		err = d.db.QueryRow(fmt.Sprintf(sqlSelect単位ForAggregation, fmt.Sprintf("max(%s)", strconv.Quote(string(fld))), where.String())).Scan(&x)
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
func (d daoDb単位) Insert(dr *Dto単位) (err error) {
	err = d.db.QueryRow(sqlInsert単位, dr.Fldコード, dr.Fld名称).Scan(&dr.FldID)
	if err != nil {
		err = xerrors.Errorf(": %w", err)
		return
	}
	dr.rowState = Added
	d.dm.dt単位 = append(d.dm.dt単位, dr)
	d.dm.mapIDvsDr単位[dr.FldID] = dr
	d.dm.mapコードvsDr単位[dr.Fldコード] = dr

	return
}
func (d daoDb単位) MultiInsert(dt []*Dto単位) (err error) {
	cs := a.Chunk(dt, 1000)
	for _, c := range cs {
		vals := make([]string, len(c))
		args := make([]interface{}, len(c)*2)
		for i, dr := range c {
			vals[i] = fmt.Sprintf(sqlValue2, 2*i+1, 2*i+2)
			args[2*i] = dr.Fldコード
			args[2*i+1] = dr.Fld名称
			dr.rowState = Added
		}
		_, err = d.db.Exec(fmt.Sprintf(sqlMultiInsert単位, strings.Join(vals, ",")), args...)
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
	}
	return
}
func (d daoDb単位) UpdateBy(dr *Dto単位) (cnt int64, err error) {
	if dr.Ub.Count() == 0 {
		dr.rowState = UnChanged
		return
	}
	s, w, execArgs := dr.Ub.build(newWb単位WithPrimaryKeys(dr.FldID))
	sql := fmt.Sprintf(sqlUpdate単位, s, w)
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
func (d daoDb単位) UpdateW(ub *ub単位, wb Wb単位) (cnt int64, err error) {
	s, w, execArgs := ub.build(wb)
	sql := fmt.Sprintf(sqlUpdate単位, s, w)
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
func (d daoDb単位) DeleteBy(dr *Dto単位) (cnt int64, err error) {
	where := newWb単位WithPrimaryKeys(dr.FldID).build()
	prms, exists := where.Params()
	if !exists {
		err = xerrors.Errorf("主キーがありません。: %#v", *dr)
		return
	}
	sql := fmt.Sprintf(sqlDelete単位, where.String())
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
func (d daoDb単位) DeleteW(wb Wb単位) (cnt int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	if exists {
		sql := fmt.Sprintf(sqlDelete単位, where.String())
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
		result, errExec := d.db.Exec(sqlTruncate単位)
		if errExec != nil {
			err = xerrors.Errorf("sql=%s: %w", sqlTruncate単位, errExec)
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
