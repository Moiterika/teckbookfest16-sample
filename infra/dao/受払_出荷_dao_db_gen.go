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

type daoDb受払出荷 struct {
	dm        *DaoDbManager
	db        *sql.DB
	WbForInit Wb受払出荷
}

func (d *daoDb受払出荷) init() (err error) {
	d.dm.dt受払出荷, err = d.SelectW(d.WbForInit)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	d.dm.mapIDvsDr受払出荷 = a.ToMap(d.dm.dt受払出荷, func(e *Dto受払出荷) Id {
		return e.FldNo
	})
	return
}
func (d *daoDb受払出荷) Reset() {
	list := make([]*Dto受払出荷, 0, len(d.dm.dt受払出荷))
	for _, dr := range d.dm.dt受払出荷 {
		if dr.rowState == Deleted {
			dr.rowState = Detached
			continue
		} else {
			dr.rowState = UnChanged
			list = append(list, dr)
		}
	}
	d.dm.dt受払出荷 = list
}
func (d daoDb受払出荷) Dt() ([]*Dto受払出荷, error) {
	if len(d.dm.dt受払出荷) == 0 {
		err := d.init()
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return nil, err
		}
	}
	return d.dm.dt受払出荷, nil
}
func (d daoDb受払出荷) GetBy(id Id) (dr *Dto受払出荷, err error) {
	if len(d.dm.mapIDvsDr受払出荷) == 0 {
		err = d.init()
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
	}
	var ok bool
	dr, ok = d.dm.mapIDvsDr受払出荷[id]
	if !ok {
		err = xerrors.Errorf("受払_出荷が見つかりません。No=%d: %w", id, NotFoundError)
		return
	}
	return
}
func (d daoDb受払出荷) SelectAll() ([]*Dto受払出荷, error) {
	sql := fmt.Sprintf(sqlSelect受払出荷, "")
	rows, err := d.db.Query(sql)
	defer rows.Close()
	if err != nil {
		return nil, xerrors.Errorf("sql=%s: %w", sql, err)
	}
	var dt []*Dto受払出荷
	for rows.Next() {
		var dr Dto受払出荷
		err = rows.Scan(&dr.FldNo, &dr.Fld出荷数量, &dr.Fld出荷単位ID)
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}
		dr.rowState = Detached
		dr.Ub = NewUb受払出荷()
		dt = append(dt, &dr)
	}
	return dt, rows.Err()
}
func (d daoDb受払出荷) SelectW(wb Wb受払出荷) ([]*Dto受払出荷, error) {
	where := wb.build()
	prms, exists := where.Params()
	if exists {
		sql := fmt.Sprintf(sqlSelect受払出荷, where.String())
		rows, err := d.db.Query(sql, prms...)
		defer rows.Close()
		if err != nil {
			return nil, xerrors.Errorf("sql=%s, args=%v: %w", sql, prms, err)
		}
		var dt []*Dto受払出荷
		for rows.Next() {
			var dr Dto受払出荷
			err = rows.Scan(&dr.FldNo, &dr.Fld出荷数量, &dr.Fld出荷単位ID)
			if err != nil {
				return nil, xerrors.Errorf(": %w", err)
			}
			dr.rowState = Detached
			dr.Ub = NewUb受払出荷()
			dt = append(dt, &dr)
		}
		return dt, rows.Err()
	} else {
		return d.SelectAll()
	}
}
func (d daoDb受払出荷) Count() (cnt int64, err error) {
	return d.CountW(NewWb受払出荷())
}
func (d daoDb受払出荷) CountW(wb Wb受払出荷) (cnt int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	if exists {
		err = d.db.QueryRow(fmt.Sprintf(sqlSelect受払出荷ForAggregation, "count(\"No\")", where.String()), prms...).Scan(&cnt)
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
		return
	} else {
		err = d.db.QueryRow(fmt.Sprintf(sqlSelect受払出荷ForAggregation, "count(\"No\")", "")).Scan(&cnt)
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
		return
	}
}

// Min は受払_出荷のfld最小値を返します。
func (d daoDb受払出荷) Min(fld fld受払出荷) (min int64, err error) {
	return d.MinW(fld, NewWb受払出荷())
}
func (d daoDb受払出荷) MinW(fld fld受払出荷, wb Wb受払出荷) (min int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	var x sql.NullInt64
	if exists {
		err = d.db.QueryRow(fmt.Sprintf(sqlSelect受払出荷ForAggregation, fmt.Sprintf("min(%s)", strconv.Quote(string(fld))), where.String()), prms...).Scan(&x)
	} else {
		err = d.db.QueryRow(fmt.Sprintf(sqlSelect受払出荷ForAggregation, fmt.Sprintf("min(%s)", strconv.Quote(string(fld))), where.String())).Scan(&x)
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

// Max は受払_出荷のfld最大値を返します。
func (d daoDb受払出荷) Max(fld fld受払出荷) (max int64, err error) {
	return d.MaxW(fld, NewWb受払出荷())
}
func (d daoDb受払出荷) MaxW(fld fld受払出荷, wb Wb受払出荷) (max int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	var x sql.NullInt64
	if exists {
		err = d.db.QueryRow(fmt.Sprintf(sqlSelect受払出荷ForAggregation, fmt.Sprintf("max(%s)", strconv.Quote(string(fld))), where.String()), prms...).Scan(&x)
	} else {
		err = d.db.QueryRow(fmt.Sprintf(sqlSelect受払出荷ForAggregation, fmt.Sprintf("max(%s)", strconv.Quote(string(fld))), where.String())).Scan(&x)
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
func (d daoDb受払出荷) Insert(dr *Dto受払出荷) (err error) {
	_, err = d.db.Exec(sqlInsert受払出荷, dr.FldNo, dr.Fld出荷数量, dr.Fld出荷単位ID)
	if err != nil {
		err = xerrors.Errorf(": %w", err)
		return
	}
	dr.rowState = Added
	return
}
func (d daoDb受払出荷) MultiInsert(dt []*Dto受払出荷) (err error) {
	cs := a.Chunk(dt, 1000)
	for _, c := range cs {
		vals := make([]string, len(c))
		args := make([]interface{}, len(c)*3)
		for i, dr := range c {
			vals[i] = fmt.Sprintf(sqlValue3, 3*i+1, 3*i+2, 3*i+3)
			args[3*i] = dr.FldNo
			args[3*i+1] = dr.Fld出荷数量
			args[3*i+2] = dr.Fld出荷単位ID
			dr.rowState = Added
		}
		_, err = d.db.Exec(fmt.Sprintf(sqlMultiInsert受払出荷, strings.Join(vals, ",")), args...)
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
	}
	return
}
func (d daoDb受払出荷) UpdateBy(dr *Dto受払出荷) (cnt int64, err error) {
	if dr.Ub.Count() == 0 {
		dr.rowState = UnChanged
		return
	}
	s, w, execArgs := dr.Ub.build(newWb受払出荷WithPrimaryKeys(dr.FldNo))
	sql := fmt.Sprintf(sqlUpdate受払出荷, s, w)
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
func (d daoDb受払出荷) UpdateW(ub *ub受払出荷, wb Wb受払出荷) (cnt int64, err error) {
	s, w, execArgs := ub.build(wb)
	sql := fmt.Sprintf(sqlUpdate受払出荷, s, w)
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
func (d daoDb受払出荷) DeleteBy(dr *Dto受払出荷) (cnt int64, err error) {
	where := newWb受払出荷WithPrimaryKeys(dr.FldNo).build()
	prms, exists := where.Params()
	if !exists {
		err = xerrors.Errorf("主キーがありません。: %#v", *dr)
		return
	}
	sql := fmt.Sprintf(sqlDelete受払出荷, where.String())
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
func (d daoDb受払出荷) DeleteW(wb Wb受払出荷) (cnt int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	if exists {
		sql := fmt.Sprintf(sqlDelete受払出荷, where.String())
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
		result, errExec := d.db.Exec(sqlTruncate受払出荷)
		if errExec != nil {
			err = xerrors.Errorf("sql=%s: %w", sqlTruncate受払出荷, errExec)
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
