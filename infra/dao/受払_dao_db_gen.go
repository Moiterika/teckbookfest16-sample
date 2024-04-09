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

type daoDb受払 struct {
	dm        *DaoDbManager
	db        *sql.DB
	WbForInit Wb受払
}

func (d *daoDb受払) init() (err error) {
	d.dm.dt受払, err = d.SelectW(d.WbForInit)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	d.dm.mapIDvsDr受払 = a.ToMap(d.dm.dt受払, func(e *Dto受払) Id {
		return e.FldNo
	})
	return
}
func (d *daoDb受払) Reset() {
	list := make([]*Dto受払, 0, len(d.dm.dt受払))
	for _, dr := range d.dm.dt受払 {
		if dr.rowState == Deleted {
			dr.rowState = Detached
			continue
		} else {
			dr.rowState = UnChanged
			list = append(list, dr)
		}
	}
	d.dm.dt受払 = list
}
func (d daoDb受払) Dt() ([]*Dto受払, error) {
	if len(d.dm.dt受払) == 0 {
		err := d.init()
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return nil, err
		}
	}
	return d.dm.dt受払, nil
}
func (d daoDb受払) GetBy(id Id) (dr *Dto受払, err error) {
	if len(d.dm.mapIDvsDr受払) == 0 {
		err = d.init()
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
	}
	var ok bool
	dr, ok = d.dm.mapIDvsDr受払[id]
	if !ok {
		err = xerrors.Errorf("受払が見つかりません。No=%d: %w", id, types.ErrNotFound)
		return
	}
	return
}
func (d daoDb受払) SelectAll() ([]*Dto受払, error) {
	sql := fmt.Sprintf(sqlSelect受払, "")
	rows, err := d.db.Query(sql)
	defer rows.Close()
	if err != nil {
		return nil, xerrors.Errorf("sql=%s: %w", sql, err)
	}
	var dt []*Dto受払
	for rows.Next() {
		var dr Dto受払
		err = rows.Scan(&dr.FldNo, &dr.Fld登録日時, &dr.Fld計上月, &dr.Fld受払区分, &dr.Fld赤伝フラグ, &dr.Fld品目ID, &dr.Fld基準数量, &dr.Fld基準単位ID)
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}
		dr.rowState = Detached
		dr.Ub = NewUb受払()
		dt = append(dt, &dr)
	}
	return dt, rows.Err()
}
func (d daoDb受払) SelectW(wb Wb受払) ([]*Dto受払, error) {
	where := wb.build()
	prms, exists := where.Params()
	if exists {
		sql := fmt.Sprintf(sqlSelect受払, where.String())
		rows, err := d.db.Query(sql, prms...)
		defer rows.Close()
		if err != nil {
			return nil, xerrors.Errorf("sql=%s, args=%v: %w", sql, prms, err)
		}
		var dt []*Dto受払
		for rows.Next() {
			var dr Dto受払
			err = rows.Scan(&dr.FldNo, &dr.Fld登録日時, &dr.Fld計上月, &dr.Fld受払区分, &dr.Fld赤伝フラグ, &dr.Fld品目ID, &dr.Fld基準数量, &dr.Fld基準単位ID)
			if err != nil {
				return nil, xerrors.Errorf(": %w", err)
			}
			dr.rowState = Detached
			dr.Ub = NewUb受払()
			dt = append(dt, &dr)
		}
		return dt, rows.Err()
	} else {
		return d.SelectAll()
	}
}
func (d daoDb受払) Count() (cnt int64, err error) {
	return d.CountW(NewWb受払())
}
func (d daoDb受払) CountW(wb Wb受払) (cnt int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	if exists {
		err = d.db.QueryRow(fmt.Sprintf(sqlSelect受払ForAggregation, "count(\"No\")", where.String()), prms...).Scan(&cnt)
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
		return
	} else {
		err = d.db.QueryRow(fmt.Sprintf(sqlSelect受払ForAggregation, "count(\"No\")", "")).Scan(&cnt)
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
		return
	}
}

// Min は受払のfld最小値を返します。
func (d daoDb受払) Min(fld fld受払) (min int64, err error) {
	return d.MinW(fld, NewWb受払())
}
func (d daoDb受払) MinW(fld fld受払, wb Wb受払) (min int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	var x sql.NullInt64
	if exists {
		err = d.db.QueryRow(fmt.Sprintf(sqlSelect受払ForAggregation, fmt.Sprintf("min(%s)", strconv.Quote(string(fld))), where.String()), prms...).Scan(&x)
	} else {
		err = d.db.QueryRow(fmt.Sprintf(sqlSelect受払ForAggregation, fmt.Sprintf("min(%s)", strconv.Quote(string(fld))), where.String())).Scan(&x)
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

// Max は受払のfld最大値を返します。
func (d daoDb受払) Max(fld fld受払) (max int64, err error) {
	return d.MaxW(fld, NewWb受払())
}
func (d daoDb受払) MaxW(fld fld受払, wb Wb受払) (max int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	var x sql.NullInt64
	if exists {
		err = d.db.QueryRow(fmt.Sprintf(sqlSelect受払ForAggregation, fmt.Sprintf("max(%s)", strconv.Quote(string(fld))), where.String()), prms...).Scan(&x)
	} else {
		err = d.db.QueryRow(fmt.Sprintf(sqlSelect受払ForAggregation, fmt.Sprintf("max(%s)", strconv.Quote(string(fld))), where.String())).Scan(&x)
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
func (d daoDb受払) Insert(dr *Dto受払) (err error) {
	err = d.db.QueryRow(sqlInsert受払, dr.Fld登録日時, dr.Fld計上月, dr.Fld受払区分, dr.Fld赤伝フラグ, dr.Fld品目ID, dr.Fld基準数量, dr.Fld基準単位ID).Scan(&dr.FldNo)
	if err != nil {
		err = xerrors.Errorf(": %w", err)
		return
	}
	dr.rowState = Added
	d.dm.dt受払 = append(d.dm.dt受払, dr)
	d.dm.mapIDvsDr受払[dr.FldNo] = dr
	return
}
func (d daoDb受払) MultiInsert(dt []*Dto受払) (err error) {
	cs := a.Chunk(dt, 1000)
	for _, c := range cs {
		vals := make([]string, len(c))
		args := make([]interface{}, len(c)*7)
		for i, dr := range c {
			vals[i] = fmt.Sprintf(sqlValue7, 7*i+1, 7*i+2, 7*i+3, 7*i+4, 7*i+5, 7*i+6, 7*i+7)
			args[7*i] = dr.Fld登録日時
			args[7*i+1] = dr.Fld計上月
			args[7*i+2] = dr.Fld受払区分
			args[7*i+3] = dr.Fld赤伝フラグ
			args[7*i+4] = dr.Fld品目ID
			args[7*i+5] = dr.Fld基準数量
			args[7*i+6] = dr.Fld基準単位ID
			dr.rowState = Added
		}
		_, err = d.db.Exec(fmt.Sprintf(sqlMultiInsert受払, strings.Join(vals, ",")), args...)
		if err != nil {
			err = xerrors.Errorf(": %w", err)
			return
		}
	}
	return
}
func (d daoDb受払) UpdateBy(dr *Dto受払) (cnt int64, err error) {
	if dr.Ub.Count() == 0 {
		dr.rowState = UnChanged
		return
	}
	s, w, execArgs := dr.Ub.build(newWb受払WithPrimaryKeys(dr.FldNo))
	sql := fmt.Sprintf(sqlUpdate受払, s, w)
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
func (d daoDb受払) UpdateW(ub *ub受払, wb Wb受払) (cnt int64, err error) {
	s, w, execArgs := ub.build(wb)
	sql := fmt.Sprintf(sqlUpdate受払, s, w)
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
func (d daoDb受払) DeleteBy(dr *Dto受払) (cnt int64, err error) {
	where := newWb受払WithPrimaryKeys(dr.FldNo).build()
	prms, exists := where.Params()
	if !exists {
		err = xerrors.Errorf("主キーがありません。: %#v", *dr)
		return
	}
	sql := fmt.Sprintf(sqlDelete受払, where.String())
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
func (d daoDb受払) DeleteW(wb Wb受払) (cnt int64, err error) {
	where := wb.build()
	prms, exists := where.Params()
	if exists {
		sql := fmt.Sprintf(sqlDelete受払, where.String())
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
		result, errExec := d.db.Exec(sqlTruncate受払)
		if errExec != nil {
			err = xerrors.Errorf("sql=%s: %w", sqlTruncate受払, errExec)
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
