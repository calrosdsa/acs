package pg

import (
	_reporte "acs/domain/repository/reporte"

	
	"context"
	"database/sql"

	// "encoding/json"
	"log"

)

type repoReporte struct {
	Conn *sql.DB
}

func NewRepoReporte(db *sql.DB)_reporte.ReporteRepo{
	return &repoReporte{
		Conn: db,
	}
}

func (p *repoReporte)GetReporteEmpleado(ctx context.Context)(res []_reporte.Data,horario []_reporte.Horario,err error){
	var query string

	// query = `SELECT dd::date as date,
	// (SELECT ARRAY_AGG(date) FROM marcacion where date::date=dd::date) as times,
	// (SELECT ARRAY_AGG(type_marcacion) FROM marcacion where date::date=dd::date) as types
	// FROM generate_series
	// 		( '2024-01-01'::timestamp 
	// 		, '2024-01-06'::timestamp
	// 		, '1 day'::interval) dd;`
	log.Println("GETTING DATA")
    query = `select cast(DATEADD(DAY, value, '2024-01-01') as date) as date,
	(select STRING_AGG(convert(varchar(25), created_on, 120), ',') from marcacion where
	CAST(created_on as date) = cast(DATEADD(DAY, value, '2024-01-01') as date)) AS times,
	(select STRING_AGG(CAST(type_marcacion AS VARCHAR), ',') from marcacion where
	CAST(created_on as date) = cast(DATEADD(DAY, value, '2024-01-01') as date)) AS types
	from GENERATE_SERIES(0, 5);`
	res,err = p.fetchData(ctx,query)
	if err != nil {
		log.Println(err)
	}
	log.Println(res)
	query = `SELECT start_time,end_time,day from horario where profile_id = 1;`
	horario,err = p.fetchHorario(ctx,query)
	log.Println("HORARIOS",horario)
	return
}

func (p *repoReporte) fetchData(ctx context.Context, query string, args ...interface{}) (res []_reporte.Data, err error) {
	rows, err := p.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Println(errRow)
		}
	}()
	res = make([]_reporte.Data, 0)
	for rows.Next() {
		t := _reporte.Data{}
		err = rows.Scan(
			&t.Date,
			&t.TimesString,
			&t.TypesString,
		)
		res = append(res, t)
	}
	return res, nil
}

func (p *repoReporte) fetchHorario(ctx context.Context, query string, args ...interface{}) (res []_reporte.Horario, err error) {
	rows, err := p.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Println(errRow)
		}
	}()
	res = make([]_reporte.Horario, 0)
	for rows.Next() {
		t := _reporte.Horario{}
		err = rows.Scan(
			&t.StartTime,
			&t.EndTime,
			&t.Day,
		)
		res = append(res, t)
	}
	return res, nil
}
