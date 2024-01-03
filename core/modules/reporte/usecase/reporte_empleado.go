package usecase

import (
	"fmt"
	"log"
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
	f.SetColWidth(sheet, "D", "D", float64((9 * maxTurnos)))
	f.SetColWidth(sheet,"E","H",20)
	

	headers := []string{"", "Fecha", "Marcaciones","Horario"}	
	headers = append(headers, "Hrs. Trabajadas")
	headers = append(headers, "Hrs. Total")
	headers = append(headers, "Hrs. Trabajadas 2")
	headers = append(headers, "Hrs. Retraso")

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
	if err := f.SetCellStyle(sheet, "A2", "H2", titleStyle); err != nil {
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
		var horarioStr string
		for j := 0;j <len(c.Horario);j++ {
			if len(c.Horario) >= (1+j){
				horarioStr +=  c.Horario[j].StartTime.Format("15:04") + " - " + c.Horario[j].EndTime.Format("15:04") + "   "
				// horarioStr = strings.TrimSuffix(horarioStr," ")
			}
		}
		slice = append(slice, horarioStr)
		for j := 0;j <(maxMarks /2);j++ {
			if len(c.HorasTrabajadas) >= (1+j){
			// slice = append(slice, c.HorasTrabajadas[j])
			hrsTrabajadas += c.HorasTrabajadas[j]
			// }else{
				// slice = append(slice, " - ")
			}
		}
		// dec := time.Minute * 66
		slice = append(slice, hrsTrabajadas)
		slice = append(slice, c.Total)
		slice = append(slice, c.TotalHrsWorked)
		slice = append(slice, c.Retraso)
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
