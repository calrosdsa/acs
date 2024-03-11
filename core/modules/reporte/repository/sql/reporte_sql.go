package pg

import (
	_r "acs/domain/repository"
	"fmt"

	// _q "acs/domain/util"
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

func (m *repoReporte) GetAsistenciaEmployeeArea(ctx context.Context, d _r.ReporteRequest) (res []_r.Asistencia, err error) {
	var query string
	var (
		filterString string
	)
	if !d.ALLSitios {
		filterString = fmt.Sprintf("and a.idSitio = %d",d.IdSitio)
	}
	// idsSitioStr := _q.ArrayToString(d.IdsSitio, ",")
	// log.Println(d.IdArea, "-------", d.IdSitio,idsSitioStr)
	query = fmt.Sprintf(`select a.id,a.asistenciaDate,a.marcaciones,a.horario,a.hrsTotales,a.hrsTrabajadas,a.hrsTrabajadasEnHorario,a.hrsExcedentes,
	a.retraso, a.retraso2 ,a.countMarcaciones,coalesce(c.firtsName,''),coalesce(c.lastName,''),(''),coalesce(sitio.nombre,'')
	from TAsistencia as a
	left join TCardHolder as c on c.guid = a.cardholderGuid
	left join TSitio as sitio on sitio.id = a.idSitio
	where a.idArea = @p1 %s and a.asistenciaDate >= @p2 AND
	a.asistenciaDate <= @p3`,filterString)
	res, err = m.fetchAsistenciasUser(ctx, query, d.IdArea, d.StartDate, d.EndDate)
	if err != nil {
		return
	}
	log.Println("RESULT", res)
	return
}

func (m *repoReporte) GetAsistenciaEmployeeSitio(ctx context.Context, d _r.ReporteRequest) (res []_r.Asistencia, err error) {
	var query string
	query = `select a.id,a.asistenciaDate,a.marcaciones,a.horario,a.hrsTotales,a.hrsTrabajadas,a.hrsTrabajadasEnHorario,a.hrsExcedentes,
	a.retraso, a.retraso2  ,a.countMarcaciones,coalesce(c.firtsName,''),coalesce(c.lastName,''),coalesce(area.nombre,''),('')
	from TAsistencia as a
	left join TCardHolder as c on c.guid = cardholderGuid
	left join TArea as area on area.id = a.idArea
	where a.idSitio = @p1 and asistenciaDate >= @p2 AND
	asistenciaDate <= @p3`
	res, err = m.fetchAsistenciasUser(ctx, query, d.IdSitio, d.StartDate, d.EndDate)
	if err != nil {
		return
	}
	return
}
func (m *repoReporte) GetAsistenciaEmployee(ctx context.Context, d _r.ReporteRequest) (res []_r.Asistencia, err error) {
	var query string
	query = `select a.id,asistenciaDate,marcaciones,horario,hrsTotales,hrsTrabajadas,hrsTrabajadasEnHorario,hrsExcedentes,
	retraso, retraso2  ,countMarcaciones,(''),(''),(''),('')
	from TAsistencia as a
	where cardholderGuid = @p1 and asistenciaDate >= @p2 AND
	asistenciaDate <= @p3`
	res, err = m.fetchAsistenciasUser(ctx, query, d.CHGuid, d.StartDate, d.EndDate)
	if err != nil {
		return
	}
	return
}

func (m *repoReporte) GetEmpleadoCardHolder(ctx context.Context, guid string) (res _r.Employee, err error) {
	query := `select coalesce(c.firtsName,''), coalesce(c.lastName,''),a.nombre,s.nombre from
	 TCardHolder as c
	 left join TArea as a on a.id = c.idArea
	 left join TSitio as s on s.id = c.idSitio
	 where c.guid = @p1
	 `
	err = m.Conn.QueryRowContext(ctx, query, guid).Scan(&res.FirstName, &res.LastName, &res.Area, &res.Sitio)
	return
}

func (m *repoReporte) GetMarcacionesForReport(ctx context.Context, d _r.ReporteRequest) (res []_r.MarcacionItem, err error) {
	query := `select id,fecha,typeMarcacion,CAST(CONVERT(VARCHAR,fecha,110) AS DATE) AS DATE
	 from TMarcacionAsistencia where cardholderGuid = @p1 and fecha >= @p2 AND
	fecha <= @p3 order by fecha`
	res, err = m.fetchMarcaciones(ctx, query, d.CHGuid, d.StartDate, d.EndDate)
	return
}

func (m *repoReporte) GetEmployeesSitio(ctx context.Context, idSitio int) (res []_r.Employee, err error) {
	query := `select c.guid,coalesce(c.firtsName,''), coalesce(c.lastName,''),
	coalesce(a.nombre,''),('') from TCardHolder as c
	left join TArea as a on a.id = c.idArea 
	where idSitio = @p1`
	res, err = m.fetchEmployees(ctx, query, idSitio)
	return
}

func (m *repoReporte) GetEmployeesArea(ctx context.Context, idArea int,all bool, idSitio int) (res []_r.Employee, err error) {
	// idsStr := _q.ArrayToString(idsSitio, ",")
	var (
		filterString string
	)
	if !all {
		filterString = fmt.Sprintf("and c.idSitio = %d",idSitio)
	}
	query :=fmt.Sprintf(`select c.guid,coalesce(c.firtsName,''), coalesce(c.lastName,''),
	(''),coalesce(s.nombre,'')
	from TCardHolder as c
	left join TSitio as s on s.id = c.idSitio
	where c.idArea = @p1 %s`,filterString)
	res, err = m.fetchEmployees(ctx, query, idArea)
	return
}
func(m *repoReporte) GetSitio(ctx context.Context,idSitio int)(res _r.Sitio,err error){
	query := "select id,nombre from TSitio where id = @p1"
	err = m.Conn.QueryRowContext(ctx,query,idSitio).Scan(&res.Id,&res.Name)
	return
}

func(m *repoReporte) GetArea(ctx context.Context,idArea int)(res _r.Area,err error){
	query := "select id,nombre from TArea where id = @p1"
	err = m.Conn.QueryRowContext(ctx,query,idArea).Scan(&res.Id,&res.Name)
	return
}

// func (p *repoReporte) GetReporteEmpleado(ctx context.Context) (res []_r.Data, horario []_r.Horario, err error) {
// 	var query string
// 	log.Println("GETTING DATA")
// 	query = `select
// 	 cast(DATEADD(DAY, value, '2024-02-01') as date) as date,
// 	(select STRING_AGG(convert(varchar(25), fecha, 120), ',') from TMarcacionAsistencia where
// 	CAST(fecha as date) = cast(DATEADD(DAY, value, '2024-02-01') as date)) AS times,
// 	(select STRING_AGG(CAST(typeMarcacion AS VARCHAR), ',') from TMarcacionAsistencia where
// 	CAST(fecha as date) = cast(DATEADD(DAY, value, '2024-02-01') as date)) AS types,
// 	(select TOP 1 convert(varchar(25), fecha, 120) from TMarcacionAsistencia where
// 	    CAST(fecha as date) = cast(DATEADD(DAY, value -1, '2024-02-01') as date)
// 		order by fecha desc
// 	    ) AS firstM,
//     (select TOP 1 convert(varchar(25), fecha, 120) from TMarcacionAsistencia where
// 	    CAST(fecha as date) = cast(DATEADD(DAY, value + 1, '2024-02-01') as date)
// 		order by fecha
// 	    ) AS lastM,
//      (select TOP 1 CAST(typeMarcacion AS VARCHAR) from TMarcacionAsistencia where
// 	    CAST(fecha as date) = cast(DATEADD(DAY, value - 1, '2024-02-01') as date)
// 		order by fecha  desc
// 	    ) AS firstT,
// 	(select TOP 1 CAST(typeMarcacion AS VARCHAR) from TMarcacionAsistencia where
// 	    CAST(fecha as date) = cast(DATEADD(DAY, value + 1, '2024-02-01') as date)
// 		order by fecha
// 	    ) AS lastT
// 	from GENERATE_SERIES(0, 5);`
// 	res, err = p.fetchData(ctx, query)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	log.Println(res)
// 	query = `SELECT horaEntrada,horaSalida,diaNumber from THorarioPerfil where idPerfil = 1;`
// 	horario, err = p.fetchHorario(ctx, query)
// 	log.Println("HORARIOS", horario)
// 	return
// }

func (p *repoReporte) fetchMarcaciones(ctx context.Context, query string, args ...interface{}) (res []_r.MarcacionItem, err error) {
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
	res = make([]_r.MarcacionItem, 0)
	for rows.Next() {
		t := _r.MarcacionItem{}
		err = rows.Scan(
			&t.Id,
			&t.Date,
			&t.TypeMarcacion,
			&t.DateString,
		)
		res = append(res, t)
	}
	return res, nil
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
			&t.HrsExcedentes,
			&t.Retraso,
			&t.Retraso2,
			&t.CountMarcaciones,
			&t.EmployeeFirstName,
			&t.EmployeeLastName,
			&t.AreaName,
			&t.SitioName,
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

func (p *repoReporte) fetchEmployees(ctx context.Context, query string, args ...interface{}) (res []_r.Employee, err error) {
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
	res = make([]_r.Employee, 0)
	for rows.Next() {
		t := _r.Employee{}
		err = rows.Scan(
			&t.CardHolderGuid,
			&t.FirstName,
			&t.LastName,
			&t.Area,
			&t.Sitio,
		)
		res = append(res, t)
	}
	return res, nil
}
