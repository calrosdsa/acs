package sql

import (
	_r "acs/domain/repository"
	"context"
	"database/sql"
	"log"
	"time"
)

type asistenciaRepository struct {
	Conn *sql.DB
}

func NewRepository(conn *sql.DB) _r.AsistenciaRepository {
	return &asistenciaRepository{
		Conn: conn,
	}
}

func (m *asistenciaRepository) GetAllCardHolders(ctx context.Context) (res []_r.CardHolderUser, err error) {
	query := `select c.guid,c.idArea,c.idSitio from  TCardHolder as c`
	res,err = m.fetchCardHolders(ctx,query)
	return 
}



func (m *asistenciaRepository) CreateAsistencia(ctx context.Context, d _r.Asistencia) (err error) {
	log.Println(d)
	query := `insert into TAsistencia(asistenciaDate,cardholderGuid,retraso,retraso2,hrsTotales,hrsTrabajadas,hrsTrabajadasEnHorario,
		marcaciones,horario,countMarcaciones,idSitio,idArea,doorGuid,hrsExcedentes)
		values(@p1,@p2,@p3,@p4,@p5,@p6,@p7,@p8,@p9,@p10,@p11,@p12,@p13,@p14)`
	// marcaciones,horario) values(@p1,@p2,@p3,@p4,@p5,@p6,@p7,@p8)`
	log.Println("COUNT MARCACIONES -----------",d.CountMarcaciones)
	_, err = m.Conn.ExecContext(ctx, query, d.AsistenciaDate, d.CardHolderGuid, d.Retraso, d.Retraso2, d.HrsTotales, d.HrsTrabajadas,
		d.HrsTrabajadasEnHorario, d.Marcaciones, d.Horario, d.CountMarcaciones, d.IdSitio, d.IdArea, d.DoorGuid,d.HrsExcedentes)
	return
}

func (m *asistenciaRepository) UpdateAsistencia(ctx context.Context, d _r.Asistencia) (err error) {
	query := `update TAsistencia set retraso = @p1, hrsTotales = @p2, hrsTrabajadas = @p3, hrsTrabajadasEnHorario = @p4,
	marcaciones = @p5,countMarcaciones = @p8, retraso2 = @p9, hrsExcedentes = @p10
	where asistenciaDate = @p6 and cardholderGuid = @p7`
	_, err = m.Conn.ExecContext(ctx, query, d.Retraso, d.HrsTotales, d.HrsTrabajadas,
		d.HrsTrabajadasEnHorario, d.Marcaciones, d.AsistenciaDate, d.CardHolderGuid, d.CountMarcaciones, d.Retraso2, d.HrsExcedentes)
	return
}

func (m *asistenciaRepository) ExistAsistencia(ctx context.Context, chguid string, fecha time.Time) (res bool, err error) {
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
func (p *asistenciaRepository) GetEmployeDataHorarioNocturno(ctx context.Context, chGuid string, fecha string,idPerfil int) (res _r.Data, horario []_r.Horario, err error) {
	var query string
	log.Println("GETTING DATA")
	query = `select 
	CAST(CONVERT(VARCHAR,fecha,110) as date),
	(''),(''),
	 (select TOP 1 convert(varchar(25), fecha, 120) from TMarcacionAsistencia where
	    CAST(fecha as date) = cast(DATEADD(DAY, -1, @p1) as date) 
		and  cardholderGuid = @p2
		order by fecha desc
	    ) AS lastME,
		(select TOP 1 convert(varchar(25), fecha, 120) from TMarcacionAsistencia where
	    CAST(fecha as date) = @p1 
		and cardholderGuid = @p2
		order by fecha desc
	    ) AS lastMS,
     (select TOP 1 CAST(typeMarcacion AS VARCHAR) from TMarcacionAsistencia where
	    CAST(fecha as date) = cast(DATEADD(DAY, -1, @p1) as date) 
	    and cardholderGuid = @p2
		order by fecha  desc
	    ) AS T1,
		(select TOP 1 CAST(typeMarcacion AS VARCHAR) from TMarcacionAsistencia where
	    CAST(fecha as date) = @p1 
		and cardholderGuid = @p2
		order by fecha desc
	    ) AS T2
	from TMarcacionAsistencia where CAST(fecha as date)= CAST(@p1 as date) 
	and cardholderGuid = @p2
	group by CAST(CONVERT(VARCHAR,fecha,110) as date);`
	err = p.Conn.QueryRowContext(ctx, query, fecha, chGuid).Scan(&res.Date, &res.TimesString, &res.TypesString, &res.FirstM,
		&res.LastM, &res.FirstT, &res.LastT)
	if err != nil {
		log.Println(err)
	}
	log.Println("EMPLOYEE DATA HORARIO NOCTUNO",res)
	query = `SELECT horaEntrada,horaSalida,diaNumber from THorarioPerfil where idPerfil = @p1;`
	horario, err = p.fetchHorario(ctx, query,idPerfil)
	if err != nil {
		log.Println("FAIL FECTH HORARIO")
	}
	log.Println("HORARIOS", horario)
	return
}

func (p *asistenciaRepository) GetEmployeData(ctx context.Context, chGuid string, fecha string,idPerfil int) (res _r.Data, horario []_r.Horario, err error) {
	var query string
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
	log.Println("EMPLOYEE DATA",res)
	query = `SELECT horaEntrada,horaSalida,diaNumber from THorarioPerfil where idPerfil = @p1;`
	horario, err = p.fetchHorario(ctx, query,idPerfil)
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

func (p *asistenciaRepository) fetchCardHolders(ctx context.Context, query string, args ...interface{}) (res []_r.CardHolderUser, err error) {
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
	res = make([]_r.CardHolderUser, 0)
	for rows.Next() {
		t := _r.CardHolderUser{}
		err = rows.Scan(
			&t.Guid,
			&t.IdArea,
			&t.IdSitio,
		)
		res = append(res, t)
	}
	return res, nil
}
