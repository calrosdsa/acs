package sql

import (
	_r "acs/domain/repository"
	"context"
	"database/sql"
	"log"
)

type asistenciaRepository struct {
	Conn *sql.DB
}

func NewRepository(conn *sql.DB) _r.AsistenciaRepository {
	return &asistenciaRepository{
		Conn: conn,
	}
}
func (m *asistenciaRepository) GetAsistencia(ctx context.Context, chguid string, fecha string) (res _r.Asistencia, err error) {
	return
}

func (m *asistenciaRepository) GetAsistenciasUser(ctx context.Context, chGuid string, page, size int) (res []_r.Asistencia,
	count int, err error) {
	var query string
	query = `select id,asistenciaDate,marcaciones,horario,hrsTotales,hrsTrabajadas,hrsTrabajadasEnHorario,retraso
	from TAsistencia where cardholderGuid = @p1`
	res, err = m.fetchAsistenciasUser(ctx, query, chGuid)
	if err != nil {
		return
	}
	query = `select count(*) from TAsistencia where cardholderGuid = @p1`
	err = m.Conn.QueryRowContext(ctx, query, chGuid).Scan(&count)
	return
}

func (m *asistenciaRepository) CreateAsistencia(ctx context.Context, d _r.Asistencia) (err error) {
	log.Println(d)
	query := `insert into TAsistencia(asistenciaDate,cardholderGuid,retraso,retraso2,hrsTotales,hrsTrabajadas,hrsTrabajadasEnHorario,
		marcaciones,horario,countMarcaciones,countTurnos) values(@p1,@p2,@p3,@p4,@p5,@p6,@p7,@p8,@p9,@p10,@p11)`
	// marcaciones,horario) values(@p1,@p2,@p3,@p4,@p5,@p6,@p7,@p8)`

	_, err = m.Conn.ExecContext(ctx, query, d.AsistenciaDate, d.CardHolderGuid, d.Retraso,d.Retraso2, d.HrsTotales, d.HrsTrabajadas,
		d.HrsTrabajadasEnHorario, d.Marcaciones, d.Horario, d.CountMarcaciones, d.CountTurnos)
	return
}

func (m *asistenciaRepository) UpdateAsistencia(ctx context.Context, d _r.Asistencia) (err error) {
	query := `update TAsistencia set retraso = @p1, hrsTotales = @p2, hrsTrabajadas = @p3, hrsTrabajadasEnHorario = @p4,
	marcaciones = @p5,countMarcaciones = @p8,countTurnos = @p9, retraso2 = @p10, hrsExcedentes = @p11
	where asistenciaDate = @p6 and cardholderGuid = @p7`
	_, err = m.Conn.ExecContext(ctx, query, d.Retraso, d.HrsTotales, d.HrsTrabajadas,
		d.HrsTrabajadasEnHorario, d.Marcaciones, d.AsistenciaDate, d.CardHolderGuid, d.CountMarcaciones,
		d.CountTurnos,d.Retraso2,d.HrsExcedentes)
	return
}

func (m *asistenciaRepository) ExistAsistencia(ctx context.Context, chguid string, fecha string) (res bool, err error) {
	query := `SELECT
    CASE WHEN EXISTS 
    (
        SELECT * FROM TAsistencia WHERE asistenciaDate = @p1 and cardholderGuid = @p2
    )
    THEN 'TRUE'
    ELSE 'FALSE'
    END`
	err = m.Conn.QueryRowContext(ctx, query, fecha, chguid).Scan(&res)
	return
}

func (p *asistenciaRepository) GetEmployeData(ctx context.Context, chGuid string, fecha string) (res _r.Data, horario []_r.Horario, err error) {
	var query string

	// query = `SELECT dd::date as date,
	// (SELECT ARRAY_AGG(date) FROM TMarcacionAsistencia where date::date=dd::date) as times,
	// (SELECT ARRAY_AGG(typeMarcacion) FROM TMarcacionAsistencia where date::date=dd::date) as types
	// FROM generate_series
	// 		( '2024-02-01'::timestamp
	// 		, '2024-01-06'::timestamp
	// 		, '1 day'::interval) dd;`
	log.Println("GETTING DATA")
	query = `select 
	CAST(CONVERT(VARCHAR,fecha,110) as date),
	 STRING_AGG(convert(varchar(25), fecha, 120), ','),
	 STRING_AGG(CAST(typeMarcacion AS VARCHAR), ','),
	 (select TOP 1 convert(varchar(25), fecha, 120) from TMarcacionAsistencia where
	    CAST(fecha as date) = cast(DATEADD(DAY, -1, fecha) as date) and cardholderGuid = @p2
		order by fecha desc
	    ) AS firstM,
		(select TOP 1 convert(varchar(25), fecha, 120) from TMarcacionAsistencia where
	    CAST(fecha as date) = cast(DATEADD(DAY,  1, fecha) as date) and cardholderGuid = @p2
		order by fecha 
	    ) AS lastM,
     (select TOP 1 CAST(typeMarcacion AS VARCHAR) from TMarcacionAsistencia where
	    CAST(fecha as date) = cast(DATEADD(DAY, -1, fecha) as date) and cardholderGuid = @p2
		order by fecha  desc
	    ) AS firstT,
	(select TOP 1 CAST(typeMarcacion AS VARCHAR) from TMarcacionAsistencia where
	    CAST(fecha as date) = cast(DATEADD(DAY, 1, fecha) as date) and cardholderGuid = @p2
		order by fecha 
	    ) AS lastT
	from TMarcacionAsistencia where CAST(fecha as date)= CAST(@p1 as date) and cardholderGuid = @p2
	group by CAST(CONVERT(VARCHAR,fecha,110) as date);`
	err = p.Conn.QueryRowContext(ctx, query, fecha, chGuid).Scan(&res.Date, &res.TimesString, &res.TypesString, &res.FirstM,
		&res.LastM, &res.FirstT, &res.LastT)
	if err != nil {
		log.Println(err)
	}
	log.Println(res)
	query = `SELECT horaEntrada,horaSalida,diaNumber from THorarioPerfil where idPerfil = 8;`
	horario, err = p.fetchHorario(ctx, query)
	if err != nil {
		log.Println("FAIL FECTH HORARIO")
	}
	log.Println("HORARIOS", horario)
	return
}

func (m *asistenciaRepository) InsertMarcacion(ctx context.Context, d _r.TMarcacionAsistencia) (err error) {
	query := `insert into TMarcacionAsistencia(accessPointGuid,cardholderGuid,credentialGuid,doorGuid,eventType,fecha,
		idZona,typeMarcacion) values (@p1,@p2,@p3,@p4,@p5,@p6,@p7,@p8)`
	_, err = m.Conn.ExecContext(ctx, query, d.AccessPointGuid, d.CardHolderGuid, d.CredentialGuid, d.DoorGuid, d.EventType,
		d.Fecha, d.IdZona, d.TypeMarcacion)
	return
}

func (p *asistenciaRepository) fetchAsistenciasUser(ctx context.Context, query string, args ...interface{}) (res []_r.Asistencia, err error) {
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
		)
		res = append(res, t)
	}
	return res, nil
}

func (p *asistenciaRepository) fetchData(ctx context.Context, query string, args ...interface{}) (res []_r.Data, err error) {
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

func (p *asistenciaRepository) fetchHorario(ctx context.Context, query string, args ...interface{}) (res []_r.Horario, err error) {
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
