// Code generated by xlsx2go.go; DO NOT EDIT.

package dao

import (
	"database/sql"
	"fmt"
	a "github.com/Moiterika/a"
	xerrors "golang.org/x/xerrors"
	"strconv"
	"strings"
	objects "techbookfest16-sample/domain/objects"
)

type daoTrn生産用品目区分 struct {
	dm        *DaoTrnManager
	trn       *sql.Tx
	WbForInit Wb生産用品目区分
}

func (d *daoTrn生産用品目区分) init() (err error) {
	d.dm.dt生産用品目区分, err = d.SelectW(d.WbForInit)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	d.dm.mapIDvsDr生産用品目区分 = a.ToMap(d.dm.dt生産用品目区分, func(e *Dto生産用品目区分) Id {
		return e.FldID
	})
	d.dm.mapコードvsDr生産用品目区分 = a.ToMap(d.dm.dt生産用品目区分, func(e *Dto生産用品目区分) objects.Code生産用品目区分 {
		return e.Fldコード
	})
	return
}
func (d *daoTrn生産用品目区分) Reset() {
	list := make([]*Dto生産用品目区分, 0, len(d.dm.dt生産用品目区分))
	for _, dr := range d.dm.dt生産用品目区分 {
		if dr.rowState == Deleted {
			dr.rowState = Detached
			continue
		} else {
			dr.rowState = UnChanged
			list = append(list, dr)
		}
	}
	d.dm.dt生産用品目区分 = list
	d.dm.mapコードvsDr生産用品目区分 = a.ToMap(d.dm.dt生産用品目区分, func(e *Dto生産用品目区分) objects.Code生産用品目区分 {
		return e.Fldコード
	})
}
func (d daoTrn生産用品目区分) Dt() ([]*Dto生産用品目区分, error) {
	if len(d.dm.dt生産用品目区分) == 0 {
		err := d.init()
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return nil, err
		}
	}
	return d.dm.dt生産用品目区分, nil
}
func (d daoTrn生産用品目区分) GetBy(id Id) (dr *Dto生産用品目区分, err error) {
	if len(d.dm.mapIDvsDr生産用品目区分) == 0 {
		err = d.init()
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
	}
	var ok bool
	dr, ok = d.dm.mapIDvsDr生産用品目区分[id]
	if !ok {
		err = xerrors.Errorf("生産用品目区分が見つかりません。ID=%d: %w", id, NotFoundError)
		return
	}
	return
}
func (d daoTrn生産用品目区分) GetByCode(コード objects.Code生産用品目区分) (dr *Dto生産用品目区分, err error) {
	if len(d.dm.mapコードvsDr生産用品目区分) == 0 {
		err = d.init()
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
	}
	var ok bool
	dr, ok = d.dm.mapコードvsDr生産用品目区分[コード]
	if !ok {
		err = xerrors.Errorf("生産用品目区分が見つかりません。コード=%s: %w", コード, NotFoundError)
		return
	}
	return
}
func (d daoTrn生産用品目区分) SelectAll() ([]*Dto生産用品目区分, error) {
	sql := fmt.Sprintf(sqlSelect生産用品目区分, "")
	rows, err := d.trn.Query(sql)
	defer rows.Close()
	if err != nil {
		return nil, xerrors.Errorf("sql=%s: %w", sql, err)
	}
	var dt []*Dto生産用品目区分
	for rows.Next() {
		var dr Dto生産用品目区分
		err = rows.Scan(&dr.FldID, &dr.Fldコード, &dr.Fld名称, &dr.Fld何かのフラグ1, &dr.Fld何かのフラグ2)
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}
		dr.rowState = Detached
		dr.Ub = NewUb生産用品目区分()
		dt = append(dt, &dr)
	}
	return dt, rows.Err()
}
func (d daoTrn生産用品目区分) SelectW(wb Wb生産用品目区分) ([]*Dto生産用品目区分, error) {
	where := wb.build()
	prms, exists := where.Params()
	if exists {
		sql := fmt.Sprintf(sqlSelect生産用品目区分, where.String())
		rows, err := d.trn.Query(sql, prms...)
		defer rows.Close()
		if err != nil {
			return nil, xerrors.Errorf("sql=%s, args=%v: %w", sql, prms, err)
		}
		var dt []*Dto生産用品目区分
		for rows.Next() {
			var dr Dto生産用品目区分
			err = rows.Scan(&dr.FldID, &dr.Fldコード, &dr.Fld名称, &dr.Fld何かのフラグ1, &dr.Fld何かのフラグ2)
			if err != nil {
				return nil, xerrors.Errorf(": %w", err)
			}
			dr.rowState = Detached
			dr.Ub = NewUb生産用品目区分()
			dt = append(dt, &dr)
		}
		return dt, rows.Err()
	} else {
		return d.SelectAll()
	}
}
func (d daoTrn生産用品目区分) Count() (cnt int64, err error) {
	return d.CountW(NewWb生産用品目区分())
}
func (d daoTrn生産用品目区分) CountW(wb Wb生産用品目区分) (cnt int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	if exists {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelect生産用品目区分ForAggregation, "count(\"ID\")", where.String()), prms...).Scan(&cnt)
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
		return
	} else {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelect生産用品目区分ForAggregation, "count(\"ID\")", "")).Scan(&cnt)
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
		return
	}
}

// Min は生産用品目区分のfld最小値を返します。
func (d daoTrn生産用品目区分) Min(fld fld生産用品目区分) (min int64, err error) {
	return d.MinW(fld, NewWb生産用品目区分())
}
func (d daoTrn生産用品目区分) MinW(fld fld生産用品目区分, wb Wb生産用品目区分) (min int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	var x sql.NullInt64
	if exists {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelect生産用品目区分ForAggregation, fmt.Sprintf("min(%s)", strconv.Quote(string(fld))), where.String()), prms...).Scan(&x)
	} else {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelect生産用品目区分ForAggregation, fmt.Sprintf("min(%s)", strconv.Quote(string(fld))), where.String())).Scan(&x)
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

// Max は生産用品目区分のfld最大値を返します。
func (d daoTrn生産用品目区分) Max(fld fld生産用品目区分) (max int64, err error) {
	return d.MaxW(fld, NewWb生産用品目区分())
}
func (d daoTrn生産用品目区分) MaxW(fld fld生産用品目区分, wb Wb生産用品目区分) (max int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	var x sql.NullInt64
	if exists {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelect生産用品目区分ForAggregation, fmt.Sprintf("max(%s)", strconv.Quote(string(fld))), where.String()), prms...).Scan(&x)
	} else {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelect生産用品目区分ForAggregation, fmt.Sprintf("max(%s)", strconv.Quote(string(fld))), where.String())).Scan(&x)
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
func (d daoTrn生産用品目区分) Insert(dr *Dto生産用品目区分) (id Id, err error) {
	err = d.trn.QueryRow(sqlInsert生産用品目区分, dr.Fldコード, dr.Fld名称, dr.Fld何かのフラグ1, dr.Fld何かのフラグ2).Scan(&id)
	if err != nil {
		err = xerrors.Errorf(": %w", err)
		return
	}
	dr.rowState = Added
	return
}
func (d daoTrn生産用品目区分) MultiInsert(dt []*Dto生産用品目区分) (err error) {
	cs := a.Chunk(dt, 1000)
	for _, c := range cs {
		vals := make([]string, len(c))
		args := make([]interface{}, len(c)*4)
		for i, dr := range c {
			vals[i] = fmt.Sprintf(sqlValue4, 4*i+1, 4*i+2, 4*i+3, 4*i+4)
			args[4*i] = dr.Fldコード
			args[4*i+1] = dr.Fld名称
			args[4*i+2] = dr.Fld何かのフラグ1
			args[4*i+3] = dr.Fld何かのフラグ2
			dr.rowState = Added
		}
		_, err = d.trn.Exec(fmt.Sprintf(sqlMultiInsert生産用品目区分, strings.Join(vals, ",")), args...)
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
	}
	return
}
func (d daoTrn生産用品目区分) UpdateBy(dr *Dto生産用品目区分) (cnt int64, err error) {
	if dr.Ub.Count() == 0 {
		dr.rowState = UnChanged
		return
	}
	s, w, execArgs := dr.Ub.build(newWb生産用品目区分WithPrimaryKeys(dr.FldID))
	sql := fmt.Sprintf(sqlUpdate生産用品目区分, s, w)
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
func (d daoTrn生産用品目区分) UpdateW(ub *ub生産用品目区分, wb Wb生産用品目区分) (cnt int64, err error) {
	s, w, execArgs := ub.build(wb)
	sql := fmt.Sprintf(sqlUpdate生産用品目区分, s, w)
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
func (d daoTrn生産用品目区分) DeleteBy(dr *Dto生産用品目区分) (cnt int64, err error) {
	where := newWb生産用品目区分WithPrimaryKeys(dr.FldID).build()
	prms, exists := where.Params()
	if !exists {
		err = xerrors.Errorf("主キーがありません。: %#v", *dr)
		return
	}
	sql := fmt.Sprintf(sqlDelete生産用品目区分, where.String())
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
func (d daoTrn生産用品目区分) DeleteW(wb Wb生産用品目区分) (cnt int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	if exists {
		sql := fmt.Sprintf(sqlDelete生産用品目区分, where.String())
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
		result, errExec := d.trn.Exec(sqlTruncate生産用品目区分)
		if errExec != nil {
			err = xerrors.Errorf("sql=%s: %w", sqlTruncate生産用品目区分, errExec)
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
