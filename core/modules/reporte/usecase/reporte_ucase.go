package usecase

import (
	_reporte "acs/domain/repository/reporte"
	"strings"

	"bytes"
	"context"
	"log"
	"time"
)

type reporteUseCase struct {
	reporteRepo _reporte.ReporteRepo
	timeout     time.Duration
}

func NewUseCase(timeout time.Duration, reporteRepo _reporte.ReporteRepo) _reporte.ReporteUseCase {
	return &reporteUseCase{
		timeout:     timeout,
		reporteRepo: reporteRepo,
	}
}

func (u *reporteUseCase) GetReporteEmpleado(ctx context.Context, buffer *bytes.Buffer) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	res, horario, err := u.reporteRepo.GetReporteEmpleado(ctx)

	horarioByDay := make(map[int][]_reporte.Horario)
	var (
		maxMarks  int
		maxTurnos int
	)
	for i := 0; i < len(horario); i++ {
		
		log.Println(horario[i].StartTime.Hour(), horario[i].StartTime.Minute())
		current, exist := horarioByDay[horario[i].Day]

		if exist {
			horarioByDay[horario[i].Day] = append(current, horario[i])
		} else {
			horarioByDay[horario[i].Day] = []_reporte.Horario{horario[i]}
		}
	}
	for i := 0; i < len(res); i++ {
		res[i].Times = strings.Split(res[i].TimesString, ",")
		res[i].Types = strings.Split(res[i].TypesString, ",")
		log.Println(res[i].Times)
		log.Println(res[i].Types)
		currentT, err := getDateTime(res[i].Date, "2006-01-02T15:04:05Z")
		if err != nil {
			log.Println("fail to parse", err)
		}
		// horarios, _ := horarioByDay[currentT.Day()]
		res[i].Horario = horarioByDay[currentT.Day()]
		times := res[i].Times
		types := res[i].Types
		log.Println(types)
		for j := 0; j < len(times); j++ {
			if types[j] == "1" {

				if len(times) >= (j + 2) {
					if types[j+1] == "2" {
						
						j++
						start, err := getDateTime(times[j-1], "2006-01-02 15:04:05")
						if err != nil {
							log.Println("Fail parse", err)
						}
						log.Println("OBTEING TIME END")
						end, err := getDateTime(times[j], "2006-01-02 15:04:05")
						if err != nil {
							log.Println("Fail parse", err)
						}
						diff := end.Sub(start)
						res[i].HorasTrabajadas = append(res[i].HorasTrabajadas, diff)
						if len(res[i].HorasTrabajadas) > maxTurnos {
							maxTurnos = len(res[i].HorasTrabajadas)
						}
						th,thw,retraso :=checkWorkedHours(res[i].Horario,start,end)
						res[i].Retraso += retraso
						res[i].Total = th
						res[i].TotalHrsWorked +=thw 
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
	}

	err = ReporteEmpleado(res, maxTurnos, maxMarks, buffer)
	if err != nil {
		log.Println(err)
	}

	return
}


func checkWorkedHours(horario []_reporte.Horario,start time.Time,end time.Time)(
	total time.Duration,totalWorked time.Duration,retraso time.Duration){
	for i := 0;i <len(horario);i++{
		var count int
		diff := horario[i].EndTime.Sub(horario[i].StartTime)
		total += diff
		end = time.Date(0000,01,01,end.Hour(),end.Minute(),00,100,time.UTC)
		start = time.Date(0000,01,01,start.Hour(),start.Minute(),00,100,time.UTC)
		horario[i].StartTime =  time.Date(0000,01,01,horario[i].StartTime.Hour(),horario[i].StartTime.Minute(),00,100,time.UTC)
		horario[i].EndTime =  time.Date(0000,01,01,horario[i].EndTime.Hour(),horario[i].EndTime.Minute(),00,100,time.UTC)
		// log.Println("Compare",horario[i].StartTime,horario[i].EndTime,start,end)
		
		var countT time.Time
		var countTm time.Time
		if horario[i].EndTime.After(start) && end.After(horario[i].StartTime){
			if start.Before(horario[i].StartTime){
				countTm = horario[i].StartTime.Add(time.Minute+1)
			}
				for j := 0;j<int(diff.Abs().Minutes());j++{
				countT = horario[i].StartTime.Add(time.Minute * time.Duration(j+1))
				if countT.After(start) && countTm.After(horario[i].StartTime) && countT.Before(end) {
					count++
					totalWorked = time.Minute * time.Duration(count)
					}else{
						if count == 0{
							retraso = time.Minute * time.Duration(j +1)
						}
				}
			}
		}
			
	}
	return
}


func getDateTime(str string, layout string) (t time.Time, err error) {
	log.Println("Time to parse", str)
	t, err = time.Parse(layout, str)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
