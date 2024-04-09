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

type daoTrn受払製造実績 struct {
	dm        *DaoTrnManager
	trn       *sql.Tx
	WbForInit Wb受払製造実績
}

func (d *daoTrn受払製造実績) init() (err error) {
	d.dm.dt受払製造実績, err = d.SelectW(d.WbForInit)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	d.dm.mapIDvsDr受払製造実績 = a.ToMap(d.dm.dt受払製造実績, func(e *Dto受払製造実績) Id {
		return e.FldNo
	})
	return
}
func (d *daoTrn受払製造実績) Reset() {
	list := make([]*Dto受払製造実績, 0, len(d.dm.dt受払製造実績))
	for _, dr := range d.dm.dt受払製造実績 {
		if dr.rowState == Deleted {
			dr.rowState = Detached
			continue
		} else {
			dr.rowState = UnChanged
			list = append(list, dr)
		}
	}
	d.dm.dt受払製造実績 = list
}
func (d daoTrn受払製造実績) Dt() ([]*Dto受払製造実績, error) {
	if len(d.dm.dt受払製造実績) == 0 {
		err := d.init()
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return nil, err
		}
	}
	return d.dm.dt受払製造実績, nil
}
func (d daoTrn受払製造実績) GetBy(id Id) (dr *Dto受払製造実績, err error) {
	if len(d.dm.mapIDvsDr受払製造実績) == 0 {
		err = d.init()
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
	}
	var ok bool
	dr, ok = d.dm.mapIDvsDr受払製造実績[id]
	if !ok {
		err = xerrors.Errorf("受払_製造実績が見つかりません。No=%d: %w", id, types.ErrNotFound)
		return
	}
	return
}
func (d daoTrn受払製造実績) SelectAll() ([]*Dto受払製造実績, error) {
	sql := fmt.Sprintf(sqlSelect受払製造実績, "")
	rows, err := d.trn.Query(sql)
	defer rows.Close()
	if err != nil {
		return nil, xerrors.Errorf("sql=%s: %w", sql, err)
	}
	var dt []*Dto受払製造実績
	for rows.Next() {
		var dr Dto受払製造実績
		err = rows.Scan(&dr.FldNo, &dr.Fld製造数量, &dr.Fld製造単位ID, &dr.Fld製造指図ID)
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}
		dr.rowState = Detached
		dr.Ub = NewUb受払製造実績()
		dt = append(dt, &dr)
	}
	return dt, rows.Err()
}
func (d daoTrn受払製造実績) SelectW(wb Wb受払製造実績) ([]*Dto受払製造実績, error) {
	where := wb.build()
	prms, exists := where.Params()
	if exists {
		sql := fmt.Sprintf(sqlSelect受払製造実績, where.String())
		rows, err := d.trn.Query(sql, prms...)
		defer rows.Close()
		if err != nil {
			return nil, xerrors.Errorf("sql=%s, args=%v: %w", sql, prms, err)
		}
		var dt []*Dto受払製造実績
		for rows.Next() {
			var dr Dto受払製造実績
			err = rows.Scan(&dr.FldNo, &dr.Fld製造数量, &dr.Fld製造単位ID, &dr.Fld製造指図ID)
			if err != nil {
				return nil, xerrors.Errorf(": %w", err)
			}
			dr.rowState = Detached
			dr.Ub = NewUb受払製造実績()
			dt = append(dt, &dr)
		}
		return dt, rows.Err()
	} else {
		return d.SelectAll()
	}
}
func (d daoTrn受払製造実績) Count() (cnt int64, err error) {
	return d.CountW(NewWb受払製造実績())
}
func (d daoTrn受払製造実績) CountW(wb Wb受払製造実績) (cnt int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	if exists {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelect受払製造実績ForAggregation, "count(\"No\")", where.String()), prms...).Scan(&cnt)
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
		return
	} else {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelect受払製造実績ForAggregation, "count(\"No\")", "")).Scan(&cnt)
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
		return
	}
}

// Min は受払_製造実績のfld最小値を返します。
func (d daoTrn受払製造実績) Min(fld fld受払製造実績) (min int64, err error) {
	return d.MinW(fld, NewWb受払製造実績())
}
func (d daoTrn受払製造実績) MinW(fld fld受払製造実績, wb Wb受払製造実績) (min int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	var x sql.NullInt64
	if exists {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelect受払製造実績ForAggregation, fmt.Sprintf("min(%s)", strconv.Quote(string(fld))), where.String()), prms...).Scan(&x)
	} else {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelect受払製造実績ForAggregation, fmt.Sprintf("min(%s)", strconv.Quote(string(fld))), where.String())).Scan(&x)
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

// Max は受払_製造実績のfld最大値を返します。
func (d daoTrn受払製造実績) Max(fld fld受払製造実績) (max int64, err error) {
	return d.MaxW(fld, NewWb受払製造実績())
}
func (d daoTrn受払製造実績) MaxW(fld fld受払製造実績, wb Wb受払製造実績) (max int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	var x sql.NullInt64
	if exists {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelect受払製造実績ForAggregation, fmt.Sprintf("max(%s)", strconv.Quote(string(fld))), where.String()), prms...).Scan(&x)
	} else {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelect受払製造実績ForAggregation, fmt.Sprintf("max(%s)", strconv.Quote(string(fld))), where.String())).Scan(&x)
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
func (d daoTrn受払製造実績) Insert(dr *Dto受払製造実績) (err error) {
	_, err = d.trn.Exec(sqlInsert受払製造実績, dr.FldNo, dr.Fld製造数量, dr.Fld製造単位ID, dr.Fld製造指図ID)
	if err != nil {
		err = xerrors.Errorf(": %w", err)
		return
	}
	dr.rowState = Added
	d.dm.dt受払製造実績 = append(d.dm.dt受払製造実績, dr)
	d.dm.mapIDvsDr受払製造実績[dr.FldNo] = dr
	return
}
func (d daoTrn受払製造実績) MultiInsert(dt []*Dto受払製造実績) (err error) {
	cs := a.Chunk(dt, 1000)
	for _, c := range cs {
		vals := make([]string, len(c))
		args := make([]interface{}, len(c)*4)
		for i, dr := range c {
			vals[i] = fmt.Sprintf(sqlValue4, 4*i+1, 4*i+2, 4*i+3, 4*i+4)
			args[4*i] = dr.FldNo
			args[4*i+1] = dr.Fld製造数量
			args[4*i+2] = dr.Fld製造単位ID
			args[4*i+3] = dr.Fld製造指図ID
			dr.rowState = Added
		}
		_, err = d.trn.Exec(fmt.Sprintf(sqlMultiInsert受払製造実績, strings.Join(vals, ",")), args...)
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
	}
	return
}
func (d daoTrn受払製造実績) UpdateBy(dr *Dto受払製造実績) (cnt int64, err error) {
	if dr.Ub.Count() == 0 {
		dr.rowState = UnChanged
		return
	}
	s, w, execArgs := dr.Ub.build(newWb受払製造実績WithPrimaryKeys(dr.FldNo))
	sql := fmt.Sprintf(sqlUpdate受払製造実績, s, w)
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
func (d daoTrn受払製造実績) UpdateW(ub *ub受払製造実績, wb Wb受払製造実績) (cnt int64, err error) {
	s, w, execArgs := ub.build(wb)
	sql := fmt.Sprintf(sqlUpdate受払製造実績, s, w)
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
func (d daoTrn受払製造実績) DeleteBy(dr *Dto受払製造実績) (cnt int64, err error) {
	where := newWb受払製造実績WithPrimaryKeys(dr.FldNo).build()
	prms, exists := where.Params()
	if !exists {
		err = xerrors.Errorf("主キーがありません。: %#v", *dr)
		return
	}
	sql := fmt.Sprintf(sqlDelete受払製造実績, where.String())
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
func (d daoTrn受払製造実績) DeleteW(wb Wb受払製造実績) (cnt int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	if exists {
		sql := fmt.Sprintf(sqlDelete受払製造実績, where.String())
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
		result, errExec := d.trn.Exec(sqlTruncate受払製造実績)
		if errExec != nil {
			err = xerrors.Errorf("sql=%s: %w", sqlTruncate受払製造実績, errExec)
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
