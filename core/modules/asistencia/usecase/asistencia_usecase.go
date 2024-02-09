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
	// Get data from db
	res, horario, err := u.repo.GetEmployeData(ctx, d.CardHolderGuid, d.Fecha)
	if err != nil {
		u.logger.LogError("UpdateAsistenciaFromIncomingData_GetEmployeData", "asistencia_usecase.go", err)
	}
	// Declare variables for horario
	horarioByDay := make(map[int][]_r.Horario)
	var (
		maxMarks  int
		maxTurnos int

		lastTime     time.Time
		hrsExcedente time.Duration
	)
	lastRetrasoR := make(map[int]time.Duration)

	log.Println("DEFAULT VALUE TIME", lastTime)
	// Append data to map horarioByDay
	for i := 0; i < len(horario); i++ {
		current, exist := horarioByDay[horario[i].Day]
		if exist {
			horarioByDay[horario[i].Day] = append(current, horario[i])
		} else {
			horarioByDay[horario[i].Day] = []_r.Horario{horario[i]}
		}
	}

	// Validar que la marcacion del dia de ayer no sea nula
	if res.FirstM != nil && res.FirstT != nil {
		// Validar que la marcacion del dia anterior se de tipo entrada
		if *res.FirstT == _r.MarcacionEntrada {
			// Append marcacion a la lista de marcaciones hora  y marcaciones tipo
			res.Times = append(res.Times, *res.FirstM)
			res.Types = append(res.Types, *res.FirstT)
		}
	}
	// Convertir string to slice y adjuntar a  lista de marcaciones hora  y marcaciones tipo
	res.Times = append(res.Times, strings.Split(res.TimesString, ",")...)
	res.Types = append(res.Types, strings.Split(res.TypesString, ",")...)

	// Validar que la marcacion del dia siguiente no sea nula
	if res.LastM != nil && res.LastT != nil {
		// Validar que la marcacion del dia siguiente sea de tipo salida
		if *res.LastT == _r.MarcacionSalida {
			// Append marcacion a la lista de marcaciones hora  y marcaciones tipo
			res.Times = append(res.Times, *res.LastM)
			res.Types = append(res.Types, *res.LastT)
		}
	}
	// Convert string date to time
	log.Println("DATE",res.Date)
	currentT, err := getDateTime(res.Date, "2006-01-02T15:04:05Z")
	if err != nil {
		log.Println("fail to parse", err)
	}
	// Get Horario dando el dia de la fecha
	res.Horario = horarioByDay[int(currentT.Weekday())]

	log.Println("CURRENT DAY", currentT, currentT.Day(),int(currentT.Weekday()))
	times := res.Times
	types := res.Types

	// Iterar sobre las fechas de las marcaciones
	for j := 0; j < len(times); j++ {
		if types[j] == _r.MarcacionEntrada {

			if len(times) >= (j+2) && types[j+1] == _r.MarcacionSalida {
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

				th, thw, retraso, lastRetrasoR := checkWorkedHours(res.Horario, start, end, currentT, lastTime, lastRetrasoR)
				

				log.Println(lastRetrasoR)
				lastTime = end
				res.Retraso += retraso
				res.HorasAsignadas = th
				res.TotalHrsWorkedInSchedule += thw
				// res.Retraso2 += retraso2
				log.Println("TotalWorkedHrs", res.TotalHrsWorkedInSchedule)
				j++

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
	res.HorasRestantes = res.HorasAsignadas - res.TotalHrsWorkedInSchedule
	for _, v := range lastRetrasoR {
		res.Retraso2 += v
	}

	if hrsTrabajadas > res.HorasAsignadas {
		hrsExcedente = hrsTrabajadas - res.HorasAsignadas
	}

	asistencia := _r.Asistencia{
		CardHolderGuid:         d.CardHolderGuid,
		AsistenciaDate:         res.Date,
		Retraso:                res.Retraso.Seconds(),
		Retraso2:               res.Retraso2.Seconds(),
		HrsTrabajadasEnHorario: res.TotalHrsWorkedInSchedule.Seconds(),
		HrsTotales:             res.HorasAsignadas.Seconds(),
		HrsTrabajadas:          hrsTrabajadas.Seconds(),
		HrsExcedentes:          hrsExcedente.Seconds(),
		Horario:                horarioString,
		Marcaciones:            marcaciones,
		CountTurnos:            maxTurnos,
		CountMarcaciones:       maxMarks,
	}
	u.CreateOrUpdateAsistencia(context.Background(), asistencia)
	return
}

func checkWorkedHours(horario []_r.Horario, mStart time.Time, mEnd time.Time, currentT time.Time,
	lastPrevTime time.Time, lastRetraso map[int]time.Duration) (
	total time.Duration, totalWorked time.Duration, retraso time.Duration, lastRetrasoR map[int]time.Duration) {
	for i := 0; i < len(horario); i++ {

		var (
			count int
			// Declarar variable para contar los minutos empezando desde el inicio del horario
			countT time.Time
			//Declarar variable inicial
			countTm time.Time

			retraso2 time.Duration
		)
		// var countRetraso2 int

		//Get intervalo de tiempoe del turno
		diff := horario[i].EndTime.Sub(horario[i].StartTime)

		//Update horas asignadas
		total += diff

		//StartTime del turno
		StartTime := time.Date(currentT.Year(), currentT.Month(), currentT.Day(), horario[i].StartTime.Hour(), horario[i].StartTime.Minute(), 00, 100, time.UTC)

		log.Println("Continue", i)
		//EndTime del turno
		EndTime := time.Date(currentT.Year(), currentT.Month(), currentT.Day(), horario[i].EndTime.Hour(), horario[i].EndTime.Minute(), 00, 100, time.UTC)

		//Example
		// StartTime = 08:00  EndTime = 12:00
		// mStart = 08:12  mEnd = 09:40
		// lastPrevTime = 09:40
		// Solo permitir agragar retraso si lastPrevTime es menor a la hora de entrada
		// mStart = 10:12  mEnd = 12:40

		//Retraso 12 mn

		// Validar que la hora de salida sea despues a la marcacion de entrada
		// Validar que la hora de la marcacion de salida sea despues del horario de entrada
		if EndTime.After(mStart) && mEnd.After(StartTime) {
			// Si la marcacion de entrada es antes del inicio del horario
			// Tomar como inicio la hora inicial del horario
			if mStart.Before(StartTime) {
				countTm = StartTime.Add(time.Minute * 1)
			} else {
				// Si la marcacion de entrada es despues del inicio de horario
				// Tomar como inicio la hora de la marcacion
				countTm = mStart.Add(time.Minute * 1)
			}
			for j := 0; j <= int(diff.Abs().Minutes()); j++ {
				countT = StartTime.Add(time.Minute * time.Duration(j))
				// log.Println(countT, mEnd)
				// Validar que el countT sea despues a la hora de la marcacion de entrada
				// Validar que countTm sea despues a la hora de entrada
				// Validar que el countT sea menor a la hora de la marcacion de salida
				if countT.After(mStart) && countTm.After(StartTime) && countT.Before(mEnd) {

					count++
					totalWorked = time.Minute * time.Duration(count)

				} else {
					if count == 0 && (lastPrevTime.Before(StartTime) || lastPrevTime.Equal(StartTime)) {
						retraso = time.Minute * time.Duration(j)
					} else if countT.After(mEnd) || countT.Equal(mEnd) {
						// log.Println("ADDED RETRASO 2")
						retraso2 += time.Minute * time.Duration(1)
					}
				}
			}


			// Add retraso2
			v, e := lastRetraso[i]
			if e {
				if (retraso2 < v && retraso2 > 0) || mEnd.Before(EndTime) {
					retraso2 -=  (time.Minute * time.Duration(1))
					lastRetraso[i] = retraso2 + (time.Second * time.Duration(60-mEnd.Second()))
				}
			} else {
				if retraso2 > 0 || mEnd.Before(EndTime) {
					retraso2 -= (time.Minute * time.Duration(1))
					lastRetraso[i] = retraso2 + (time.Second * time.Duration(60-mEnd.Second()))
				}
			}
			lastRetrasoR = lastRetraso

			// Solo sumar segundos de la marcacion de entrada al retraso si el retraso es mayor a 0
			if retraso > 0 || mStart.After(StartTime) {
				log.Println("Adding sconds",count)
				retraso += time.Second * time.Duration(mStart.Second())
			}

			if mStart.After(StartTime){
				totalWorked -= time.Second * time.Duration(mStart.Second())
			}

			// Solo sumar segundos de la marcacion final si esta es menor a la hora de la salida
			if mEnd.Before(EndTime) {
				totalWorked += time.Second * time.Duration(mEnd.Second())
			}

			// En caso que el total de horas trabajadas en horario se mayor al intervalo entre entrada y salida restart 1 minuto
			if totalWorked > diff {
				totalWorked = totalWorked - (time.Minute * time.Duration(1))
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
