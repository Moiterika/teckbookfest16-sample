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

type daoTrn品目 struct {
	dm        *DaoTrnManager
	trn       *sql.Tx
	WbForInit Wb品目
}

func (d *daoTrn品目) init() (err error) {
	d.dm.dt品目, err = d.SelectW(d.WbForInit)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	d.dm.mapIDvsDr品目 = a.ToMap(d.dm.dt品目, func(e *Dto品目) Id {
		return e.FldID
	})
	d.dm.mapコードvsDr品目 = a.ToMap(d.dm.dt品目, func(e *Dto品目) objects.Code品目 {
		return e.Fldコード
	})
	return
}
func (d *daoTrn品目) Reset() {
	list := make([]*Dto品目, 0, len(d.dm.dt品目))
	for _, dr := range d.dm.dt品目 {
		if dr.rowState == Deleted {
			dr.rowState = Detached
			continue
		} else {
			dr.rowState = UnChanged
			list = append(list, dr)
		}
	}
	d.dm.dt品目 = list
	d.dm.mapコードvsDr品目 = a.ToMap(d.dm.dt品目, func(e *Dto品目) objects.Code品目 {
		return e.Fldコード
	})
}
func (d daoTrn品目) Dt() ([]*Dto品目, error) {
	if len(d.dm.dt品目) == 0 {
		err := d.init()
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return nil, err
		}
	}
	return d.dm.dt品目, nil
}
func (d daoTrn品目) GetBy(id Id) (dr *Dto品目, err error) {
	if len(d.dm.mapIDvsDr品目) == 0 {
		err = d.init()
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
	}
	var ok bool
	dr, ok = d.dm.mapIDvsDr品目[id]
	if !ok {
		err = xerrors.Errorf("品目が見つかりません。ID=%d: %w", id, NotFoundError)
		return
	}
	return
}
func (d daoTrn品目) GetByCode(コード objects.Code品目) (dr *Dto品目, err error) {
	if len(d.dm.mapコードvsDr品目) == 0 {
		err = d.init()
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
	}
	var ok bool
	dr, ok = d.dm.mapコードvsDr品目[コード]
	if !ok {
		err = xerrors.Errorf("品目が見つかりません。コード=%s: %w", コード, NotFoundError)
		return
	}
	return
}
func (d daoTrn品目) SelectAll() ([]*Dto品目, error) {
	sql := fmt.Sprintf(sqlSelect品目, "")
	rows, err := d.trn.Query(sql)
	defer rows.Close()
	if err != nil {
		return nil, xerrors.Errorf("sql=%s: %w", sql, err)
	}
	var dt []*Dto品目
	for rows.Next() {
		var dr Dto品目
		err = rows.Scan(&dr.FldID, &dr.Fldコード, &dr.Fld名称, &dr.Fld基準単位ID, &dr.Fld生産用品目区分ID)
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}
		dr.rowState = Detached
		dr.Ub = NewUb品目()
		dt = append(dt, &dr)
	}
	return dt, rows.Err()
}
func (d daoTrn品目) SelectW(wb Wb品目) ([]*Dto品目, error) {
	where := wb.build()
	prms, exists := where.Params()
	if exists {
		sql := fmt.Sprintf(sqlSelect品目, where.String())
		rows, err := d.trn.Query(sql, prms...)
		defer rows.Close()
		if err != nil {
			return nil, xerrors.Errorf("sql=%s, args=%v: %w", sql, prms, err)
		}
		var dt []*Dto品目
		for rows.Next() {
			var dr Dto品目
			err = rows.Scan(&dr.FldID, &dr.Fldコード, &dr.Fld名称, &dr.Fld基準単位ID, &dr.Fld生産用品目区分ID)
			if err != nil {
				return nil, xerrors.Errorf(": %w", err)
			}
			dr.rowState = Detached
			dr.Ub = NewUb品目()
			dt = append(dt, &dr)
		}
		return dt, rows.Err()
	} else {
		return d.SelectAll()
	}
}
func (d daoTrn品目) Count() (cnt int64, err error) {
	return d.CountW(NewWb品目())
}
func (d daoTrn品目) CountW(wb Wb品目) (cnt int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	if exists {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelect品目ForAggregation, "count(\"ID\")", where.String()), prms...).Scan(&cnt)
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
		return
	} else {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelect品目ForAggregation, "count(\"ID\")", "")).Scan(&cnt)
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
		return
	}
}

// Min は品目のfld最小値を返します。
func (d daoTrn品目) Min(fld fld品目) (min int64, err error) {
	return d.MinW(fld, NewWb品目())
}
func (d daoTrn品目) MinW(fld fld品目, wb Wb品目) (min int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	var x sql.NullInt64
	if exists {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelect品目ForAggregation, fmt.Sprintf("min(%s)", strconv.Quote(string(fld))), where.String()), prms...).Scan(&x)
	} else {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelect品目ForAggregation, fmt.Sprintf("min(%s)", strconv.Quote(string(fld))), where.String())).Scan(&x)
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

// Max は品目のfld最大値を返します。
func (d daoTrn品目) Max(fld fld品目) (max int64, err error) {
	return d.MaxW(fld, NewWb品目())
}
func (d daoTrn品目) MaxW(fld fld品目, wb Wb品目) (max int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	var x sql.NullInt64
	if exists {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelect品目ForAggregation, fmt.Sprintf("max(%s)", strconv.Quote(string(fld))), where.String()), prms...).Scan(&x)
	} else {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelect品目ForAggregation, fmt.Sprintf("max(%s)", strconv.Quote(string(fld))), where.String())).Scan(&x)
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
func (d daoTrn品目) Insert(dr *Dto品目) (id Id, err error) {
	err = d.trn.QueryRow(sqlInsert品目, dr.Fldコード, dr.Fld名称, dr.Fld基準単位ID, dr.Fld生産用品目区分ID).Scan(&id)
	if err != nil {
		err = xerrors.Errorf(": %w", err)
		return
	}
	dr.rowState = Added
	return
}
func (d daoTrn品目) MultiInsert(dt []*Dto品目) (err error) {
	cs := a.Chunk(dt, 1000)
	for _, c := range cs {
		vals := make([]string, len(c))
		args := make([]interface{}, len(c)*4)
		for i, dr := range c {
			vals[i] = fmt.Sprintf(sqlValue4, 4*i+1, 4*i+2, 4*i+3, 4*i+4)
			args[4*i] = dr.Fldコード
			args[4*i+1] = dr.Fld名称
			args[4*i+2] = dr.Fld基準単位ID
			args[4*i+3] = dr.Fld生産用品目区分ID
			dr.rowState = Added
		}
		_, err = d.trn.Exec(fmt.Sprintf(sqlMultiInsert品目, strings.Join(vals, ",")), args...)
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
	}
	return
}
func (d daoTrn品目) UpdateBy(dr *Dto品目) (cnt int64, err error) {
	if dr.Ub.Count() == 0 {
		dr.rowState = UnChanged
		return
	}
	s, w, execArgs := dr.Ub.build(newWb品目WithPrimaryKeys(dr.FldID))
	sql := fmt.Sprintf(sqlUpdate品目, s, w)
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
func (d daoTrn品目) UpdateW(ub *ub品目, wb Wb品目) (cnt int64, err error) {
	s, w, execArgs := ub.build(wb)
	sql := fmt.Sprintf(sqlUpdate品目, s, w)
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
func (d daoTrn品目) DeleteBy(dr *Dto品目) (cnt int64, err error) {
	where := newWb品目WithPrimaryKeys(dr.FldID).build()
	prms, exists := where.Params()
	if !exists {
		err = xerrors.Errorf("主キーがありません。: %#v", *dr)
		return
	}
	sql := fmt.Sprintf(sqlDelete品目, where.String())
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
func (d daoTrn品目) DeleteW(wb Wb品目) (cnt int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	if exists {
		sql := fmt.Sprintf(sqlDelete品目, where.String())
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
		result, errExec := d.trn.Exec(sqlTruncate品目)
		if errExec != nil {
			err = xerrors.Errorf("sql=%s: %w", sqlTruncate品目, errExec)
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
