package usecase

import (
	_r "acs/domain/repository"
	"context"
	"log"
	"strings"

	// "sync"
	"time"
)

type asistenciaUseCase struct {
	timeout time.Duration
	repo    _r.AsistenciaRepository
	logger  _r.Logger
	util    _r.Util
}

func NewUseCase(timeout time.Duration, repo _r.AsistenciaRepository, logger _r.Logger, util _r.Util) _r.AsistenciaUseCase {
	return &asistenciaUseCase{
		repo:    repo,
		timeout: timeout,
		logger:  logger,
		util:    util,
	}
}

func (u *asistenciaUseCase) GetAsistenciasUser(ctx context.Context, chGuid string, page, size int) (res []_r.Asistencia,
	nextPage int, count int, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	page = u.util.PaginationValues(page)
	res, count, err = u.repo.GetAsistenciasUser(ctx, chGuid, page, size)
	if err != nil {
		u.logger.LogError("GetAsistenciasUser", "asistencia_usecase", err)
		return
	}
	nextPage = u.util.GetNextPage(len(res), size, page+1)
	return
}

func (u *asistenciaUseCase) GetAsistencia(ctx context.Context, chGuid string, fecha string) (res _r.Asistencia, err error) {
	return
}

func (u *asistenciaUseCase) InsertMarcacion(ctx context.Context, d _r.TMarcacionAsistencia) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	err = u.repo.InsertMarcacion(ctx, d)
	if err != nil {
		u.logger.LogError("InsertMarcacion", "asistencia_usecase", err)
		return
	}
	return
}

func (u *asistenciaUseCase) CreateOrUpdateAsistencia(ctx context.Context, d _r.Asistencia) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	exist, err := u.repo.ExistAsistencia(ctx, d.CardHolderGuid, d.AsistenciaDate)
	if err != nil {
		u.logger.LogError("CreateOrUpdateAsistencia_ExistAsistencia", "asistencia_usecase", err)
		return
	}
	if exist {
		log.Println("UPDATE ASISTENCIA")
		err = u.repo.UpdateAsistencia(ctx, d)
		if err != nil {
			u.logger.LogError("CreateOrUpdateAsistencia_UpdateAsistencia", "asistencia_usecase", err)
			return
		}
	} else {
		log.Println("CREATE ASISTENCIA")
		err = u.repo.CreateAsistencia(ctx, d)
		if err != nil {
			u.logger.LogError("CreateOrUpdateAsistencia_CreateAsistencia", "asistencia_usecase", err)
			return
		}
	}
	return
}

func (u *asistenciaUseCase) UpdateAsistenciaFromIncomingData(ctx context.Context, d _r.TMarcacionAsistencia) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	res, horario, err := u.repo.GetEmployeData(ctx, d.CardHolderGuid, d.Fecha)
	if err != nil {
		u.logger.LogError("UpdateAsistenciaFromIncomingData_GetEmployeData", "asistencia_usecase.go", err)
	}
	horarioByDay := make(map[int][]_r.Horario)
	var (
		maxMarks  int
		maxTurnos int
	)
	for i := 0; i < len(horario); i++ {
		current, exist := horarioByDay[horario[i].Day]
		if exist {
			horarioByDay[horario[i].Day] = append(current, horario[i])
		} else {
			horarioByDay[horario[i].Day] = []_r.Horario{horario[i]}
		}
	}
	if res.FirstM != nil && res.FirstT != nil {
		if *res.FirstT == MarcacionEntrada {
			res.Times = append(res.Times, *res.FirstM)
			res.Types = append(res.Types, *res.FirstT)
		}
	}
	res.Times = append(res.Times, strings.Split(res.TimesString, ",")...)
	res.Types = append(res.Types, strings.Split(res.TypesString, ",")...)
	if res.LastM != nil && res.LastT != nil {
		if *res.LastT == MarcacionSalida {
			res.Times = append(res.Times, *res.LastM)
			res.Types = append(res.Types, *res.LastT)
		}
	}
	currentT, err := getDateTime(res.Date, "2006-01-02T15:04:05Z")
	if err != nil {
		log.Println("fail to parse", err)
	}
	res.Horario = horarioByDay[currentT.Day()]
	times := res.Times
	types := res.Types
	for j := 0; j < len(times); j++ {
		if types[j] == MarcacionEntrada {
			if len(times) >= (j + 2) {
				if types[j+1] == MarcacionSalida {
					start, err := getDateTime(times[j], "2006-01-02 15:04:05")
					if err != nil {
						log.Println("Fail parse", err)
					}
					end, err := getDateTime(times[j+1], "2006-01-02 15:04:05")
					if err != nil {
						log.Println("Fail parse", err)
					}
					diff := end.Sub(start)
					res.HorasTrabajadas = append(res.HorasTrabajadas, diff)
					if len(res.HorasTrabajadas) > maxTurnos {
						maxTurnos = len(res.Horario)
					}
					th, thw, retraso := checkWorkedHours(res.Horario, start, end, currentT)
					res.Retraso += retraso
					res.Total = th
					res.TotalHrsWorked += thw
					j++
				}
			} else {
				maxMarks = j + 1
			}
		} else {
			log.Println("Is out")
		}
		if j > maxMarks {
			maxMarks = j + 1
		}
	}
	log.Println("DATA RESULT", res)

	var (
		marcaciones   string
		hrsTrabajadas time.Duration
		horarioString string
	)
	for j, m := range res.Times {
		log.Println(m)
		if m != "" {
			marcaciones += m[11:16] + " " + getType(res.Types[j])
			marcaciones += " - "
		}
	}
	marcaciones = strings.TrimSuffix(marcaciones, " - ")

	for j := 0; j < len(res.Horario); j++ {
		if len(res.Horario) >= (1 + j) {
			horarioString += res.Horario[j].StartTime.Format("15:04") + " - " + res.Horario[j].EndTime.Format("15:04") + "   "
		}
	}

	for j := 0; j < (maxMarks / 2); j++ {
		if len(res.HorasTrabajadas) >= (1 + j) {
			hrsTrabajadas += res.HorasTrabajadas[j]
		}
	}

	asistencia := _r.Asistencia{
		CardHolderGuid:         d.CardHolderGuid,
		AsistenciaDate:         res.Date,
		Retraso:                res.Retraso.Seconds(),
		HrsTrabajadasEnHorario: res.TotalHrsWorked.Seconds(),
		HrsTotales:             res.Total.Seconds(),
		HrsTrabajadas:          hrsTrabajadas.Seconds(),
		Horario:                horarioString,
		Marcaciones:            marcaciones,
		CountTurnos:            maxTurnos,
		CountMarcaciones:       maxMarks,
	}
	u.CreateOrUpdateAsistencia(context.Background(), asistencia)
	return
}

func checkWorkedHours(horario []_r.Horario, start time.Time, end time.Time, currentT time.Time) (
	total time.Duration, totalWorked time.Duration, retraso time.Duration) {
	for i := 0; i < len(horario); i++ {
		var count int
		diff := horario[i].EndTime.Sub(horario[i].StartTime)
		total += diff
		// end = time.Date(0000, 01, 01, end.Hour(), end.Minute(), 00, 100, time.UTC)
		// start = time.Date(0000, 01, 01, start.Hour(), start.Minute(), 00, 100, time.UTC)
		StartTime := time.Date(currentT.Year(), currentT.Month(), currentT.Day(), horario[i].StartTime.Hour(), horario[i].StartTime.Minute(), 00, 100, time.UTC)
		EndTime := time.Date(currentT.Year(), currentT.Month(), currentT.Day(), horario[i].EndTime.Hour(), horario[i].EndTime.Minute(), 00, 100, time.UTC)
		// log.Println("Compare", StartTime, EndTime, start, end)
		// log.Println("VALIDATE", EndTime.After(start), end.After(StartTime))
		var countT time.Time
		var countTm time.Time
		if EndTime.After(start) && end.After(StartTime) {
			if start.Before(StartTime) {
				countTm = StartTime.Add(time.Minute * 1)
			} else {
				countTm = start.Add(time.Minute * 1)
			}
			for j := 0; j < int(diff.Abs().Minutes()); j++ {
				countT = StartTime.Add(time.Minute * time.Duration(j))
				// log.Println(countT, countTm)
				if countT.After(start) && countTm.After(StartTime) && countT.Before(end) {
					count++
					totalWorked = time.Minute * time.Duration(count)
				} else {
					if count == 0 {
						retraso = time.Minute * time.Duration(j+1)

					}
				}
			}
		}

	}
	return
}

func getDateTime(str string, layout string) (t time.Time, err error) {
	// log.Println("Time to parse", str)
	t, err = time.Parse(layout, str)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

const (
	MarcacionEntrada = "1"
	MarcacionSalida  = "2"
)

func getType(marcacionType string) (res string) {
	switch marcacionType {
	case "1":
		return "E"
	case "2":
		return "S"
	default:
		return ""
	}
}
