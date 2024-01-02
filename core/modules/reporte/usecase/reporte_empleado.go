package usecase

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	_reporte "acs/domain/repository/reporte"
	"bytes"

	"github.com/xuri/excelize/v2"
)

func ReporteEmpleado(reservas []_reporte.Data,maxTurnos int,maxMarks int, buffer *bytes.Buffer) (err error) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()


	CreateSheet(reservas,maxTurnos,maxMarks,"sheet1" ,f)

	err = f.Write(buffer)
	if err != nil {
		log.Println(err)
	}
	return
	// Save spreadsheet by the given path.
	// if err := f.SaveAs("Book1.xlsx"); err != nil {
	// 	fmt.Println(err)
	// }
}

func CreateSheet(item []_reporte.Data,maxTurnos int,maxMarks int, sheet string, f *excelize.File) {
	log.Println("Generate reporte")
	f.NewSheet(sheet)
	f.SetColWidth(sheet, "A", "A", 5)
	f.SetColWidth(sheet, "B", "B", 15)
	f.SetColWidth(sheet, "C", "C", float64((8 * maxMarks)))
	f.SetColWidth(sheet,"D","K",18)
	

	headers := []string{"", "Fecha", "Marcaciones"}
	for i := 1;i < (maxTurnos+1);i++{
	    //  f.SetColWidth(sheet, "C", "C", float64((10 * maxMarks)))
		headers = append(headers, fmt.Sprintf("Turno %s",strconv.Itoa(i)))
	}
	for i := 0;i < (maxMarks/2);i++{
		headers = append(headers, fmt.Sprintf("Horas %s",strconv.Itoa(i+1)))
	}
	headers = append(headers, "Hrs. Trabajadas")
	headers = append(headers, "Hrs. Total")
	headers = append(headers, "Hrs. Trabajadas 2")

	titleStyle, err := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Color: "1f7f3b", Bold: true, Family: "Arial"},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"E6F4EA"}, Pattern: 1},
		Alignment: &excelize.Alignment{Vertical: "center", Horizontal: "center"},
		Border: []excelize.Border{{Type: "top", Style: 2, Color: "1f7f3b"}, {Type: "left", Style: 2, Color: "1f7f3b"},
			{Type: "bottom", Style: 2, Color: "1f7f3b"}, {Type: "right", Style: 2, Color: "1f7f3b"}},
		// Border:    []excelize.Border{{Type: "Bottom", Style: 2, Color: "1f7f3b"}},
	})
	if err != nil {
		log.Println(err)
	}
	// set style for the 'SUNDAY' to 'SATURDAY'
	if err := f.SetCellStyle(sheet, "A2", "I2", titleStyle); err != nil {
		log.Println(err)
		return
	}


	cell, err := excelize.CoordinatesToCellName(1, 2)
	if err != nil {
		log.Println(err)
	}


	f.SetSheetRow(sheet, cell, &headers)

	if err != nil {
		fmt.Println(err)
	}
	for idx, c := range item {
	// 	if c.ProfileApellido != nil {
	// 		apellido = *c.ProfileApellido
	// 	} else {
	// 		apellido = ""
	// 	}
	// 	name = *c.ProfileNombre + " " + apellido
	// 	fechaInicio = c.StartDate
	// 	// if c.FechaFin == nil { fechaFin = "" }else{ fechaFin = *c.FechaFin }
		slice := []interface{}{idx + 1, c.Date[0:10]}
		var (
			marcaciones string
			hrsTrabajadas time.Duration
		)
		for j,m := range c.Times {
			marcaciones +=  m[11:16] +" " + getType(c.Types[j])
			marcaciones += " - "
		}
		marcaciones = strings.TrimSuffix(marcaciones," - ")
			
		slice = append(slice, marcaciones)

		for j := 0;j <maxTurnos;j++ {
			if len(c.Horario) >= (1+j){
				slice = append(slice, c.Horario[j].StartTime.Format("15:04") + " - " + c.Horario[j].EndTime.Format("15:04"))
			}else {
				slice = append(slice, " - ")

			}
		}

		for j := 0;j <(maxMarks /2);j++ {
			if len(c.HorasTrabajadas) >= (1+j){
			slice = append(slice, c.HorasTrabajadas[j])
			hrsTrabajadas += c.HorasTrabajadas[j]
			}else{
				slice = append(slice, " - ")

			}
		}
		slice = append(slice, hrsTrabajadas)
		slice = append(slice, c.Total)
		slice = append(slice, c.TotalHrsWorked)
			// fechaInicio, strconv.FormatFloat(float64(c.Hours), 'g', 5, 64) + "hrs",
			// ReservaExpirada(int(c.Estado))}
		cell, err := excelize.CoordinatesToCellName(1, idx+3)
		// strconv.FormatFloat(float64(c.Paid), 'g', 5, 64) + "BOB",
		// f.SetColWidth("Sheet1","B", 35)
		if err != nil {
			log.Panicln(err)
		}
		f.SetSheetRow(sheet, cell, &slice)
	}
}

func getType(marcacionType string) (res string) {
	switch marcacionType{
	case "1":
		return "E"
	case "2":
		return "S"
	default:
		return ""
	}
}
