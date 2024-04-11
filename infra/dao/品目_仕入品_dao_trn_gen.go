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

type daoTrn品目仕入品 struct {
	dm        *DaoTrnManager
	trn       *sql.Tx
	WbForInit Wb品目仕入品
}

func (d *daoTrn品目仕入品) init() (err error) {
	d.dm.dt品目仕入品, err = d.SelectW(d.WbForInit)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	d.dm.mapIDvsDr品目仕入品 = a.ToMap(d.dm.dt品目仕入品, func(e *Dto品目仕入品) Id {
		return e.FldID
	})
	return
}
func (d *daoTrn品目仕入品) Reset() {
	list := make([]*Dto品目仕入品, 0, len(d.dm.dt品目仕入品))
	for _, dr := range d.dm.dt品目仕入品 {
		if dr.rowState == Deleted {
			dr.rowState = Detached
			continue
		} else {
			dr.rowState = UnChanged
			list = append(list, dr)
		}
	}
	d.dm.dt品目仕入品 = list
}
func (d daoTrn品目仕入品) Dt() ([]*Dto品目仕入品, error) {
	if len(d.dm.dt品目仕入品) == 0 {
		err := d.init()
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return nil, err
		}
	}
	return d.dm.dt品目仕入品, nil
}
func (d daoTrn品目仕入品) GetBy(id Id) (dr *Dto品目仕入品, err error) {
	if len(d.dm.mapIDvsDr品目仕入品) == 0 {
		err = d.init()
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
	}
	var ok bool
	dr, ok = d.dm.mapIDvsDr品目仕入品[id]
	if !ok {
		err = xerrors.Errorf("品目_仕入品が見つかりません。ID=%d: %w", id, types.ErrNotFound)
		return
	}
	return
}
func (d daoTrn品目仕入品) SelectAll() ([]*Dto品目仕入品, error) {
	sql := fmt.Sprintf(sqlSelect品目仕入品, "")
	rows, err := d.trn.Query(sql)
	if err != nil {
		return nil, xerrors.Errorf("sql=%s: %w", sql, err)
	}
	defer rows.Close()
	var dt []*Dto品目仕入品
	for rows.Next() {
		var dr Dto品目仕入品
		err = rows.Scan(&dr.FldID, &dr.Fld標準単価, &dr.Fld標準単価通貨ID, &dr.Fld標準単価単位ID)
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}
		dr.rowState = Detached
		dr.Ub = NewUb品目仕入品()
		dt = append(dt, &dr)
	}
	return dt, rows.Err()
}
func (d daoTrn品目仕入品) SelectW(wb Wb品目仕入品) ([]*Dto品目仕入品, error) {
	where := wb.build()
	prms, exists := where.Params()
	if exists {
		sql := fmt.Sprintf(sqlSelect品目仕入品, where.String())
		rows, err := d.trn.Query(sql, prms...)
		if err != nil {
			return nil, xerrors.Errorf("sql=%s, args=%v: %w", sql, prms, err)
		}
		defer rows.Close()
		var dt []*Dto品目仕入品
		for rows.Next() {
			var dr Dto品目仕入品
			err = rows.Scan(&dr.FldID, &dr.Fld標準単価, &dr.Fld標準単価通貨ID, &dr.Fld標準単価単位ID)
			if err != nil {
				return nil, xerrors.Errorf(": %w", err)
			}
			dr.rowState = Detached
			dr.Ub = NewUb品目仕入品()
			dt = append(dt, &dr)
		}
		return dt, rows.Err()
	} else {
		return d.SelectAll()
	}
}
func (d daoTrn品目仕入品) Count() (cnt int64, err error) {
	return d.CountW(NewWb品目仕入品())
}
func (d daoTrn品目仕入品) CountW(wb Wb品目仕入品) (cnt int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	if exists {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelect品目仕入品ForAggregation, "count(\"ID\")", where.String()), prms...).Scan(&cnt)
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
		return
	} else {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelect品目仕入品ForAggregation, "count(\"ID\")", "")).Scan(&cnt)
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
		return
	}
}

// Min は品目_仕入品のfld最小値を返します。
func (d daoTrn品目仕入品) Min(fld fld品目仕入品) (min int64, err error) {
	return d.MinW(fld, NewWb品目仕入品())
}
func (d daoTrn品目仕入品) MinW(fld fld品目仕入品, wb Wb品目仕入品) (min int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	var x sql.NullInt64
	if exists {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelect品目仕入品ForAggregation, fmt.Sprintf("min(%s)", strconv.Quote(string(fld))), where.String()), prms...).Scan(&x)
	} else {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelect品目仕入品ForAggregation, fmt.Sprintf("min(%s)", strconv.Quote(string(fld))), where.String())).Scan(&x)
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

// Max は品目_仕入品のfld最大値を返します。
func (d daoTrn品目仕入品) Max(fld fld品目仕入品) (max int64, err error) {
	return d.MaxW(fld, NewWb品目仕入品())
}
func (d daoTrn品目仕入品) MaxW(fld fld品目仕入品, wb Wb品目仕入品) (max int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	var x sql.NullInt64
	if exists {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelect品目仕入品ForAggregation, fmt.Sprintf("max(%s)", strconv.Quote(string(fld))), where.String()), prms...).Scan(&x)
	} else {
		err = d.trn.QueryRow(fmt.Sprintf(sqlSelect品目仕入品ForAggregation, fmt.Sprintf("max(%s)", strconv.Quote(string(fld))), where.String())).Scan(&x)
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
func (d daoTrn品目仕入品) Insert(dr *Dto品目仕入品) (err error) {
	_, err = d.trn.Exec(sqlInsert品目仕入品, dr.FldID, dr.Fld標準単価, dr.Fld標準単価通貨ID, dr.Fld標準単価単位ID)
	if err != nil {
		err = xerrors.Errorf(": %w", err)
		return
	}
	dr.rowState = Added
	d.dm.dt品目仕入品 = append(d.dm.dt品目仕入品, dr)
	d.dm.mapIDvsDr品目仕入品[dr.FldID] = dr
	return
}
func (d daoTrn品目仕入品) MultiInsert(dt []*Dto品目仕入品) (err error) {
	cs := a.Chunk(dt, 1000)
	for _, c := range cs {
		vals := make([]string, len(c))
		args := make([]interface{}, len(c)*4)
		for i, dr := range c {
			vals[i] = fmt.Sprintf(sqlValue4, 4*i+1, 4*i+2, 4*i+3, 4*i+4)
			args[4*i] = dr.FldID
			args[4*i+1] = dr.Fld標準単価
			args[4*i+2] = dr.Fld標準単価通貨ID
			args[4*i+3] = dr.Fld標準単価単位ID
			dr.rowState = Added
		}
		_, err = d.trn.Exec(fmt.Sprintf(sqlMultiInsert品目仕入品, strings.Join(vals, ",")), args...)
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
	}
	return
}
func (d daoTrn品目仕入品) UpdateBy(dr *Dto品目仕入品) (cnt int64, err error) {
	if dr.Ub.Count() == 0 {
		dr.rowState = UnChanged
		return
	}
	s, w, execArgs := dr.Ub.build(newWb品目仕入品WithPrimaryKeys(dr.FldID))
	sql := fmt.Sprintf(sqlUpdate品目仕入品, s, w)
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
func (d daoTrn品目仕入品) UpdateW(ub *ub品目仕入品, wb Wb品目仕入品) (cnt int64, err error) {
	s, w, execArgs := ub.build(wb)
	sql := fmt.Sprintf(sqlUpdate品目仕入品, s, w)
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
func (d daoTrn品目仕入品) DeleteBy(dr *Dto品目仕入品) (cnt int64, err error) {
	where := newWb品目仕入品WithPrimaryKeys(dr.FldID).build()
	prms, exists := where.Params()
	if !exists {
		err = xerrors.Errorf("主キーがありません。: %#v", *dr)
		return
	}
	sql := fmt.Sprintf(sqlDelete品目仕入品, where.String())
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
func (d daoTrn品目仕入品) DeleteW(wb Wb品目仕入品) (cnt int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	if exists {
		sql := fmt.Sprintf(sqlDelete品目仕入品, where.String())
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
		result, errExec := d.trn.Exec(sqlTruncate品目仕入品)
		if errExec != nil {
			err = xerrors.Errorf("sql=%s: %w", sqlTruncate品目仕入品, errExec)
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
