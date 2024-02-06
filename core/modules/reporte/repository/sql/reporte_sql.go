package pg

import (
	_r "acs/domain/repository"

	"context"
	"database/sql"

	// "encoding/json"
	"log"
)

type repoReporte struct {
	Conn *sql.DB
}

func NewRepoReporte(db *sql.DB) _r.ReporteRepo {
	return &repoReporte{
		Conn: db,
	}
}

func (m *repoReporte) GetReportEmploye(ctx context.Context, d _r.ReporteRequest) (res []_r.Asistencia, err error) {
	var query string
	query = `select id,asistenciaDate,marcaciones,horario,hrsTotales,hrsTrabajadas,hrsTrabajadasEnHorario,retraso,
	countMarcaciones,countTurnos
	from TAsistencia where cardholderGuid = @p1`
	res, err = m.fetchAsistenciasUser(ctx, query, d.CHGuid)
	if err != nil {
		return
	}
	return
}

func (p *repoReporte) GetReporteEmpleado(ctx context.Context) (res []_r.Data, horario []_r.Horario, err error) {
	var query string
	log.Println("GETTING DATA")
	query = `select 
	 cast(DATEADD(DAY, value, '2024-02-01') as date) as date,
	(select STRING_AGG(convert(varchar(25), fecha, 120), ',') from TMarcacionAsistencia where
	CAST(fecha as date) = cast(DATEADD(DAY, value, '2024-02-01') as date)) AS times,
	(select STRING_AGG(CAST(typeMarcacion AS VARCHAR), ',') from TMarcacionAsistencia where
	CAST(fecha as date) = cast(DATEADD(DAY, value, '2024-02-01') as date)) AS types,
	(select TOP 1 convert(varchar(25), fecha, 120) from TMarcacionAsistencia where
	    CAST(fecha as date) = cast(DATEADD(DAY, value -1, '2024-02-01') as date)
		order by fecha desc
	    ) AS firstM,
    (select TOP 1 convert(varchar(25), fecha, 120) from TMarcacionAsistencia where
	    CAST(fecha as date) = cast(DATEADD(DAY, value + 1, '2024-02-01') as date)
		order by fecha 
	    ) AS lastM,
     (select TOP 1 CAST(typeMarcacion AS VARCHAR) from TMarcacionAsistencia where
	    CAST(fecha as date) = cast(DATEADD(DAY, value - 1, '2024-02-01') as date)
		order by fecha  desc
	    ) AS firstT,
	(select TOP 1 CAST(typeMarcacion AS VARCHAR) from TMarcacionAsistencia where
	    CAST(fecha as date) = cast(DATEADD(DAY, value + 1, '2024-02-01') as date)
		order by fecha 
	    ) AS lastT
	from GENERATE_SERIES(0, 5);`
	res, err = p.fetchData(ctx, query)
	if err != nil {
		log.Println(err)
	}
	log.Println(res)
	query = `SELECT horaEntrada,horaSalida,diaNumber from THorarioPerfil where idPerfil = 1;`
	horario, err = p.fetchHorario(ctx, query)
	log.Println("HORARIOS", horario)
	return
}

func (p *repoReporte) fetchAsistenciasUser(ctx context.Context, query string, args ...interface{}) (res []_r.Asistencia, err error) {
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
	res = make([]_r.Asistencia, 0)
	for rows.Next() {
		t := _r.Asistencia{}
		err = rows.Scan(
			&t.Id,
			&t.AsistenciaDate,
			&t.Marcaciones,
			&t.Horario,
			&t.HrsTotales,
			&t.HrsTrabajadas,
			&t.HrsTrabajadasEnHorario,
			&t.Retraso,
			&t.CountMarcaciones,
			&t.CountTurnos,
		)
		res = append(res, t)
	}
	return res, nil
}

func (p *repoReporte) fetchData(ctx context.Context, query string, args ...interface{}) (res []_r.Data, err error) {
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
	res = make([]_r.Data, 0)
	for rows.Next() {
		t := _r.Data{}
		err = rows.Scan(
			&t.Date,
			&t.TimesString,
			&t.TypesString,
			&t.FirstM,
			&t.LastM,
			&t.FirstT,
			&t.LastT,
		)
		res = append(res, t)
	}
	return res, nil
}

func (p *repoReporte) fetchHorario(ctx context.Context, query string, args ...interface{}) (res []_r.Horario, err error) {
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
	res = make([]_r.Horario, 0)
	for rows.Next() {
		t := _r.Horario{}
		err = rows.Scan(
			&t.StartTime,
			&t.EndTime,
			&t.Day,
		)
		res = append(res, t)
	}
	return res, nil
}
