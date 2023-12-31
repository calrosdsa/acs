package usecase

import (
	_reporte "acs/domain/repository/reporte"
	"bytes"
	"context"
	"log"
	"time"
)

type reporteUseCase struct {
	reporteRepo _reporte.ReporteRepo
	timeout time.Duration
}

func NewUseCase(timeout time.Duration,reporteRepo _reporte.ReporteRepo) _reporte.ReporteUseCase{
	return &reporteUseCase{
		timeout: timeout,
		reporteRepo: reporteRepo,
	}
}

func (u *reporteUseCase) GetReporteEmpleado(ctx context.Context,buffer *bytes.Buffer)(err error){
	ctx,cancel := context.WithTimeout(ctx,u.timeout)
	defer cancel()
	res,horario,err := u.reporteRepo.GetReporteEmpleado(ctx)

	horarioByDay := make(map[int][]_reporte.Horario)
	var (
		maxMarks int
		maxTurnos int
	)
	for i := 0;i <len(horario);i++{
		log.Println(horario[i].StartTime.Hour(),horario[i].StartTime.Minute())
	    current,exist  :=horarioByDay[horario[i].Day]

		if exist {
			horarioByDay[horario[i].Day] = append(current,horario[i] )
		}else{
			horarioByDay[horario[i].Day] = []_reporte.Horario{horario[i]}
		}
	}
	for i := 0;i <len(res);i++ {
		currentT,err := getDateTime(res[i].Date,"2006-01-02T15:04:05Z") 
		if err != nil {
			log.Println("fail to parse",err)
		}
		horarios,_ := horarioByDay[currentT.Day()]
		res[i].Horario = horarios
		times := res[i].Times
		types := res[i].Types
		log.Println()
		for j := 0;j <len(times);j++ {
			if types[j] == "1" {		

					if len(times) >= (j +2){
						if types[j+1] == "2"{
						    j++
							start,err :=getDateTime(times[j-1],"2006-01-02 15:04:05")
							if err != nil {
								log.Println("Fail parse",err)
							}
							log.Println("OBTEING TIME END")
							end,err := getDateTime(times[j],"2006-01-02 15:04:05")
							if err != nil {
								log.Println("Fail parse",err)
							}
							diff := end.Sub(start)
							res[i].HorasTrabajadas = append(res[i].HorasTrabajadas, diff) 
							if len(res[i].HorasTrabajadas) > maxTurnos{
								maxTurnos = len(res[i].HorasTrabajadas)
								}	
								log.Printf("Minutes: %s\n", diff.String())
						}
					}else{
				        maxMarks = j + 1
					}

			}else{
				log.Println("Is out")
			}
			if j > maxMarks {
				maxMarks = j + 1
			}
		}
	}

	err = ReporteEmpleado(res,maxTurnos,maxMarks,buffer)
	if err != nil {
		log.Println(err)
	}

	return
}


func getDateTime(str string,layout string)(t time.Time,err error){
	log.Println("Time to parse",str)
	t, err = time.Parse(layout, str)
	if err != nil {
		log.Println(err)
		return
	}
	return
}