package usecase

import (
	_r "acs/domain/repository"
	"fmt"
	"log"

	// "strings"
	// "sync"

	"bytes"
	"context"

	// "log"
	"time"
)

type reporteUseCase struct {
	reporteRepo       _r.ReporteRepo
	reporteGenerator  _r.ReporteGenerator
	asistenciaUseCase _r.AsistenciaUseCase
	timeout           time.Duration
	logger            _r.Logger
	locale      _r.Locale
}

func NewUseCase(timeout time.Duration, reporteRepo _r.ReporteRepo, reporteGenerator _r.ReporteGenerator,
	asisreporteUseCase _r.AsistenciaUseCase, logger _r.Logger,locale _r.Locale) _r.ReporteUseCase {
	return &reporteUseCase{
		timeout:           timeout,
		reporteRepo:       reporteRepo,
		asistenciaUseCase: asisreporteUseCase,
		logger:            logger,
		reporteGenerator:  reporteGenerator,
		locale: locale,
	}
}

func (u *reporteUseCase) getCardHolder(ctx context.Context, guid string) (res _r.Employee, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	res, err = u.reporteRepo.GetEmpleadoCardHolder(ctx, guid)
	return
}

func (u *reporteUseCase) getMarcacionesForReport(ctx context.Context, d _r.ReporteRequest) (marcacionesGroups []_r.MarcacionGroup, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	res, err := u.reporteRepo.GetMarcacionesForReport(ctx, d)
	if err != nil {
		return
	}
	log.Println(res)
	// data := []_r.MarcacionData{}

	data := make(map[time.Time][]_r.MarcacionInterval)

	for i := 0; i < len(res); i++ {
		var marcacionInterval _r.MarcacionInterval
		// formatTime
		log.Println(i, res[i])
		if res[i].TypeMarcacion == _r.MarcacionEntradaInt {
			if (i+2) <= len(res) && res[i+1].TypeMarcacion == _r.MarcacionSalidaInt {
				marcacionInterval = _r.MarcacionInterval{
					MarcacionE: &res[i],
					MarcacionS: &res[i+1],
				}
				i++
			} else {
				marcacionInterval = _r.MarcacionInterval{MarcacionE: &res[i]}
			}
		} else {
			marcacionInterval = _r.MarcacionInterval{MarcacionS: &res[i]}
		}

		if currentValue, exist := data[res[i].DateString]; exist {
			data[res[i].DateString] = append(currentValue, marcacionInterval)
		} else {
			data[res[i].DateString] = []_r.MarcacionInterval{marcacionInterval}
		}
	}
	for k, v := range data {
		n := _r.MarcacionGroup{
			Date:        k,
			Marcaciones: v,
		}
		marcacionesGroups = append(marcacionesGroups, n)
	}

	log.Println(marcacionesGroups)

	return
}

func (u *reporteUseCase) GetReporte(ctx context.Context, d _r.ReporteRequest, buffer *bytes.Buffer) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	switch d.ReporteType {
	case _r.REPORTE_EMPLOYEE:
		u.GetReportEmploye(ctx, d, buffer)
	case _r.REPORTE_SITIO:
		u.GetReporteSitio(ctx, d, buffer)
	case _r.REPORTE_AREA:
		u.GetReporteArea(ctx, d, buffer)
	}
	return
}

func (u *reporteUseCase) GetReporteSitio(ctx context.Context, d _r.ReporteRequest, buffer *bytes.Buffer) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	sitio,err := u.reporteRepo.GetSitio(ctx,d.IdSitio)
	if err != nil {
		u.logger.LogError("GetReporteSitio.GetSitio", "reporte_ucase", err)
	}
	switch d.ReporteFormat {
	case _r.EMPLOYEE_FORMAT:
		employees, err1 := u.reporteRepo.GetEmployeesSitio(ctx, d.IdSitio)
		log.Println("EMPLOYEES", employees)
		if err1 != nil {
			u.logger.LogError("GetReporteSitio.GetEmployeesSitio", "reporte_ucase", err1)
			return err1
		}
		var employeeAsistencia []_r.EmployeeAsistencia
		for _, employee := range employees {
			d.CHGuid = employee.CardHolderGuid
			asistencias, err := u.reporteRepo.GetAsistenciaEmployee(ctx, d)
			if err != nil {
				u.logger.LogError("GetReporteSitio.GetRepoteEmploye", "reporte_ucase", err)
			}
			t := _r.EmployeeAsistencia{
				Asistencias: asistencias,
				Employee:    employee,
			}
			employeeAsistencia = append(employeeAsistencia, t)
		}
		info := _r.ReportInfo{
			SitioName:     sitio.Name,
			From:          d.StartDate,
			To:            d.EndDate,
			ReporteType:   d.ReporteType,
			ReporteFormat: d.ReporteFormat,
		}
		err = u.reporteGenerator.GenerateReporteSitioEmployees(info, employeeAsistencia, buffer, d.Lang)
		if err != nil {
			u.logger.LogError("GetReporteSitio.GenerateReporteSitioEmployee", "reporte_ucase", err)
			return
		}
	case _r.GENERAL_FORMAT:
		asistencias, err := u.reporteRepo.GetAsistenciaEmployeeSitio(ctx, d)
		if err != nil {
			u.logger.LogError("GetReporteSitio.GetAsistenciaEmployeArea", "reporte_ucase", err)
			return err
		}		
		info := _r.ReportInfo{
			SitioName:     sitio.Name,
			AreaName: u.locale.MustLocalize("AllAreas",d.Lang),
			From:          d.StartDate,
			To:            d.EndDate,
			ReporteType:   d.ReporteType,
			ReporteFormat: d.ReporteFormat,
		}
		err = u.reporteGenerator.GenerateReporteSitioGeneral(info, asistencias, buffer, d.Lang)
		if err != nil {
			u.logger.LogError("GetReporteSitio.GenerateReporteSitioEmployee", "reporte_ucase", err)
			return err
		}

	}
	return
}

func (u *reporteUseCase) GetReporteArea(ctx context.Context, d _r.ReporteRequest, buffer *bytes.Buffer) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	area,err := u.reporteRepo.GetArea(ctx,d.IdArea)
	if err != nil {
		u.logger.LogError("GetReporteSitio.GetArea", "reporte_ucase", err)
	}
	switch d.ReporteFormat {
	case _r.EMPLOYEE_FORMAT:
		employees, err1 := u.reporteRepo.GetEmployeesArea(ctx, d.IdArea,d.ALLSitios, d.IdSitio)
		if err1 != nil {
			u.logger.LogError("GetReporteSitio.GetEmployeesArea", "reporte_ucase", err1)
			return err1
		}
		var employeeAsistencia []_r.EmployeeAsistencia
		for _, employee := range employees {
			d.CHGuid = employee.CardHolderGuid
			asistencias, err := u.reporteRepo.GetAsistenciaEmployee(ctx, d)
			if err != nil {
				u.logger.LogError("GetReporteSitio.GetRepoteEmploye", "reporte_ucase", err)
			}
			t := _r.EmployeeAsistencia{
				Asistencias: asistencias,
				Employee:    employee,
			}
			employeeAsistencia = append(employeeAsistencia, t)
		}
		info := _r.ReportInfo{
			AreaName:      area.Name,
			// SitioName:     sitio.Name,
			From:          d.StartDate,
			To:            d.EndDate,
			ReporteType:   d.ReporteType,
			ReporteFormat: d.ReporteFormat,
		}
		err = u.reporteGenerator.GenerateReporteAreaEmployees(info, employeeAsistencia, buffer, d.Lang)
		if err != nil {
			u.logger.LogError("GetReporteSitio.GenerateReporteSitioEmployee", "reporte_ucase", err)
			return
		}
	case _r.GENERAL_FORMAT:
		asistencias, err := u.reporteRepo.GetAsistenciaEmployeeArea(ctx, d)
		if err != nil {
			u.logger.LogError("GetReporteSitio.GetAsistenciaEmployeArea", "reporte_ucase", err)
			return err
		}
		var sitio _r.Sitio
		if d.ALLSitios {
			sitio = _r.Sitio{Name: u.locale.MustLocalize("AllPlaces", d.Lang)}
		}else{
			sitio,err = u.reporteRepo.GetSitio(ctx,d.IdSitio)
			if err != nil {
				u.logger.LogError("GetReporteSitio.GetSitio", "reporte_ucase", err)
			}
		}
		info := _r.ReportInfo{
			AreaName:      area.Name,
			SitioName:     sitio.Name,
			From:          d.StartDate,
			To:            d.EndDate,
			ReporteType:   d.ReporteType,
			ReporteFormat: d.ReporteFormat,
		}
		err = u.reporteGenerator.GenerateReporteAreaGeneral(info, asistencias, buffer, d.Lang)
		if err != nil {
			u.logger.LogError("GetReporteSitio.GenerateReporteSitioEmployee", "reporte_ucase", err)
			return err
		}

	}
	return
}

func (u *reporteUseCase) GetReportEmploye(ctx context.Context, d _r.ReporteRequest, buffer *bytes.Buffer) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	employee, err := u.getCardHolder(ctx, d.CHGuid)
	if err != nil {
		u.logger.LogError("GetRepoteEmploye.getCardHolder", "reporte_ucase", err)
	}
	dataMarcaciones, err := u.getMarcacionesForReport(ctx, d)
	if err != nil {
		u.logger.LogError("GetRepoteEmploye.getMarcacionesForReport", "reporte_ucase", err)
	}
	res, err := u.reporteRepo.GetAsistenciaEmployee(ctx, d)
	if err != nil {
		u.logger.LogError("GetRepoteEmploye", "reporte_ucase", err)
	}
	info := _r.ReportInfo{
		EmployeName:   fmt.Sprintf("%s %s", employee.FirstName, employee.LastName),
		AreaName:      employee.Area,
		SitioName:     employee.Sitio,
		From:          d.StartDate,
		To:            d.EndDate,
		ReporteType:   d.ReporteType,
		ReporteFormat: d.ReporteFormat,
	}
	err = u.reporteGenerator.GenerateReporteEmploye(info, res, dataMarcaciones, buffer, d.Lang)
	if err != nil {
		u.logger.LogError("GetRepoteEmploye_ReportEmploye", "reporte_ucase", err)
	}
	return
}

// func (u *reporteUseCase) GetReporteEmpleado(ctx context.Context, buffer *bytes.Buffer, d _r.TMarcacionAsistencia) (err error) {
// 	ctx, cancel := context.WithTimeout(ctx, u.timeout)
// 	defer cancel()
// 	res, horario, err := u.reporteRepo.GetReporteEmpleado(ctx)
// 	horarioByDay := make(map[int][]_r.Horario)
// 	var (
// 		maxMarks  int
// 		maxTurnos int
// 	)
// 	for i := 0; i < len(horario); i++ {
// 		current, exist := horarioByDay[horario[i].Day]
// 		if exist {
// 			horarioByDay[horario[i].Day] = append(current, horario[i])
// 		} else {
// 			horarioByDay[horario[i].Day] = []_r.Horario{horario[i]}
// 		}
// 	}
// 	var wg sync.WaitGroup
// 	wg.Add(len(res))
// 	for i := 0; i < len(res); i++ {
// 		go func(i int) {
// 			if res[i].FirstM != nil && res[i].FirstT != nil {
// 				if *res[i].FirstT == MarcacionEntrada {
// 					res[i].Times = append(res[i].Times, *res[i].FirstM)
// 					res[i].Types = append(res[i].Types, *res[i].FirstT)
// 				}
// 			}
// 			res[i].Times = append(res[i].Times, strings.Split(res[i].TimesString, ",")...)
// 			res[i].Types = append(res[i].Types, strings.Split(res[i].TypesString, ",")...)
// 			if res[i].LastM != nil && res[i].LastT != nil {
// 				if *res[i].LastT == MarcacionSalida {
// 					res[i].Times = append(res[i].Times, *res[i].LastM)
// 					res[i].Types = append(res[i].Types, *res[i].LastT)
// 				}
// 			}
// 			currentT, err := getDateTime(res[i].Date, "2006-01-02T15:04:05Z")
// 			if err != nil {
// 				log.Println("fail to parse", err)
// 			}
// 			res[i].Horario = horarioByDay[currentT.Day()]
// 			times := res[i].Times
// 			types := res[i].Types
// 			for j := 0; j < len(times); j++ {
// 				if types[j] == MarcacionEntrada {

// 					if len(times) >= (j + 2) {

// 						if types[j+1] == MarcacionSalida {
// 							start, err := getDateTime(times[j], "2006-01-02 15:04:05")
// 							if err != nil {
// 								log.Println("Fail parse", err)
// 							}
// 							end, err := getDateTime(times[j+1], "2006-01-02 15:04:05")
// 							if err != nil {
// 								log.Println("Fail parse", err)
// 							}
// 							diff := end.Sub(start)
// 							res[i].HorasTrabajadas = append(res[i].HorasTrabajadas, diff)
// 							if len(res[i].HorasTrabajadas) > maxTurnos {
// 								maxTurnos = len(res[i].Horario)
// 							}
// 							th, thw, retraso := checkWorkedHours(res[i].Horario, start, end, currentT)
// 							res[i].Retraso += retraso
// 							res[i].Total = th
// 							res[i].TotalHrsWorked += thw
// 							j++
// 						}
// 					} else {
// 						maxMarks = j + 1
// 					}

// 				} else {
// 					log.Println("Is out")
// 				}
// 				if j > maxMarks {
// 					maxMarks = j + 1
// 				}
// 			}
// 			log.Println("DATA RESULT", res[i])

// 			var (
// 				marcaciones   string
// 				hrsTrabajadas time.Duration
// 				horarioString string
// 			)
// 			for j, m := range res[i].Times {
// 				log.Println(m)
// 				if m != "" {
// 					marcaciones += m[11:16] + " " + getType(res[i].Types[j])
// 					marcaciones += " - "
// 				}
// 			}
// 			marcaciones = strings.TrimSuffix(marcaciones, " - ")

// 			for j := 0; j < len(res[i].Horario); j++ {
// 				if len(res[i].Horario) >= (1 + j) {
// 					horarioString += res[i].Horario[j].StartTime.Format("15:04") + " - " + res[i].Horario[j].EndTime.Format("15:04") + "   "
// 				}
// 			}

// 			for j := 0; j < (maxMarks / 2); j++ {
// 				if len(res[i].HorasTrabajadas) >= (1 + j) {
// 					hrsTrabajadas += res[i].HorasTrabajadas[j]
// 				}
// 			}

// 			asistencia := _r.Asistencia{
// 				CardHolderGuid:         d.CardHolderGuid,
// 				AsistenciaDate:         res[i].Date,
// 				Retraso:                res[i].Retraso.Minutes(),
// 				HrsTrabajadasEnHorario: hrsTrabajadas.Minutes(),
// 				HrsTotales:             res[i].Total.Minutes(),
// 				HrsTrabajadas:          res[i].TotalHrsWorked.Minutes(),
// 				Horario:                horarioString,
// 				Marcaciones:            marcaciones,
// 			}
// 			u.asistenciaUseCase.CreateOrUpdateAsistencia(context.Background(), asistencia)
// 			wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()

// 	return
// }

// func checkWorkedHours(horario []_r.Horario, start time.Time, end time.Time, currentT time.Time) (
// 	total time.Duration, totalWorked time.Duration, retraso time.Duration) {
// 	for i := 0; i < len(horario); i++ {
// 		var count int
// 		diff := horario[i].EndTime.Sub(horario[i].StartTime)
// 		total += diff
// 		// end = time.Date(0000, 01, 01, end.Hour(), end.Minute(), 00, 100, time.UTC)
// 		// start = time.Date(0000, 01, 01, start.Hour(), start.Minute(), 00, 100, time.UTC)
// 		StartTime := time.Date(currentT.Year(), currentT.Month(), currentT.Day(), horario[i].StartTime.Hour(), horario[i].StartTime.Minute(), 00, 100, time.UTC)
// 		EndTime := time.Date(currentT.Year(), currentT.Month(), currentT.Day(), horario[i].EndTime.Hour(), horario[i].EndTime.Minute(), 00, 100, time.UTC)
// 		// log.Println("Compare", StartTime, EndTime, start, end)
// 		// log.Println("VALIDATE", EndTime.After(start), end.After(StartTime))
// 		var countT time.Time
// 		var countTm time.Time
// 		if EndTime.After(start) && end.After(StartTime) {
// 			if start.Before(StartTime) {
// 				countTm = StartTime.Add(time.Minute * 1)
// 			} else {
// 				countTm = start.Add(time.Minute * 1)
// 			}
// 			for j := 0; j < int(diff.Abs().Minutes()); j++ {
// 				countT = StartTime.Add(time.Minute * time.Duration(j))
// 				// log.Println(countT, countTm)
// 				if countT.After(start) && countTm.After(StartTime) && countT.Before(end) {
// 					count++
// 					totalWorked = time.Minute * time.Duration(count)
// 				} else {
// 					if count == 0 {
// 						retraso = time.Minute * time.Duration(j+1)

// 					}
// 				}
// 			}
// 		}

// 	}
// 	return
// }

// func getDateTime(str string, layout string) (t time.Time, err error) {
// 	// log.Println("Time to parse", str)
// 	t, err = time.Parse(layout, str)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	return
// }

// const (
// 	MarcacionEntrada = "1"
// 	MarcacionSalida  = "2"
// )

// func getType(marcacionType string) (res string) {
// 	switch marcacionType {
// 	case "1":
// 		return "E"
// 	case "2":
// 		return "S"
// 	default:
// 		return ""
// 	}
// }
