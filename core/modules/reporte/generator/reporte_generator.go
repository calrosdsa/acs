package generator

import (
	_r "acs/domain/repository"
	"fmt"
	"log"
	// "sort"
	"time"

	"github.com/xuri/excelize/v2"

	"bytes"
)

type reporteGenerator struct {
	reporteUtil _r.ReporteUtil
	logger      _r.Logger
	locale      _r.Locale
}

func New(reporteUtil _r.ReporteUtil, logger _r.Logger, locale _r.Locale) _r.ReporteGenerator {
	return &reporteGenerator{
		reporteUtil: reporteUtil,
		logger:      logger,
		locale:      locale,
	}
}

func (r *reporteGenerator) GenerateReporteEmploye(asistencias []_r.Asistencia, dataMarcaciones []_r.MarcacionGroup,
	buffer *bytes.Buffer, lang string) (err error) {
	f := excelize.NewFile()
	// f, err := excelize.OpenFile("./media/template.xlsx")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// if err = f.AddPicture("Sheet1", "A1", "./petrobras-logo.jpg", nil); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	d := _r.ReportInfo{
		EmployeName:  "Jorge Daniel Miranda Lopez",
		GerenciaName: "Recursos Humanos",
		SitioName:    "Equipetrol",
		From:         "01-01-2024",
		To:           "01-02-2024",
	}
	if err = r.CreateSheetEmploye(asistencias, d.EmployeName, f, d, lang); err != nil {
		return
	}
	if err = r.CreateSheetMarcaciones("Marcaciones", f, dataMarcaciones, d, lang); err != nil {
		return
	}
	f.DeleteSheet("sheet1")
	err = f.Write(buffer)
	if err != nil {
		log.Println(err)
	}
	return
}

func (r *reporteGenerator) CreateSheetMarcaciones(sheet string, f *excelize.File, dataMarcaciones []_r.MarcacionGroup,
	d _r.ReportInfo, lang string) (err error) {
	var (
		cellStyle       int
		cellCenterStyle int
		titleStyle      int
		titleStyle2     int
		cell            string
	)
	f.NewSheet(sheet)
	r.reporteUtil.SetUpReporteLayout(sheet, f)
	if titleStyle, err = r.reporteUtil.GetTitleStyle(f); err != nil {
		return
	}
	if titleStyle2, err = r.reporteUtil.GetTitleStyle2(f); err != nil {
		return
	}
	if cellStyle, err = r.reporteUtil.GetCommonCellStyle(f); err != nil {
		return
	}
	if cellCenterStyle, err = r.reporteUtil.GetCellCenterStyle(f); err != nil {
		return
	}

	//INFO
	if err = r.reporteUtil.SetUpHeader(sheet, f, d, titleStyle, titleStyle2, cellCenterStyle, cellStyle, lang); err != nil {
		return
	}

	f.SetColWidth(sheet, "A", "A", 5)
	f.SetColWidth(sheet, "B", "B", 20)
	f.SetColWidth(sheet, "C", "D", 22)
	f.SetColWidth(sheet, "E", "F", 18)

	headers := []string{
		r.locale.MustLocalize("Date", lang),
		r.locale.MustLocalize("Entrance", lang),
		r.locale.MustLocalize("Exit", lang),
	}

	if err = f.SetCellStyle(sheet, "B6", "D6", titleStyle); err != nil {
		log.Println(err)
		return
	}

	cell, err = excelize.CoordinatesToCellName(2, 6)
	if err != nil {
		return
	}
	f.SetSheetRow(sheet, cell, &headers)

	startRow := 7
	log.Println("DATA MARCACIONES", dataMarcaciones)
	// var keys Time.
	// for key := range dataMarcaciones {
	// 	keys = append(keys, key)
	// }
	// sort.Ints()
	for i := 0; i < len(dataMarcaciones); i++ {
		slice := []interface{}{dataMarcaciones[i].Date.Format(time.DateOnly)}
		log.Println("KEY", dataMarcaciones[i].Date.Format(time.DateOnly))
		if err := f.SetCellStyle(sheet, fmt.Sprintf("B%d", startRow), fmt.Sprintf("B%d", (startRow-1)+len(dataMarcaciones[i].Marcaciones)), titleStyle); err != nil {
			log.Println(err)
		}
		f.MergeCell(sheet, fmt.Sprintf("B%d", startRow), fmt.Sprintf("B%d", (startRow-1)+len(dataMarcaciones[i].Marcaciones)))
		cell, err := excelize.CoordinatesToCellName(2, startRow)
		if err != nil {
			log.Println(err)
		}
		f.SetSheetRow(sheet, cell, &slice)
		for j := 0; j < len(dataMarcaciones[i].Marcaciones); j++ {
			slice2 := []interface{}{}
			if dataMarcaciones[i].Marcaciones[j].MarcacionE != nil {
				slice2 = append(slice2, dataMarcaciones[i].Marcaciones[j].MarcacionE.Date.Format(time.DateTime))
			} else {
				slice2 = append(slice2, "N/A")
			}
			if dataMarcaciones[i].Marcaciones[j].MarcacionS != nil {
				slice2 = append(slice2, dataMarcaciones[i].Marcaciones[j].MarcacionS.Date.Format(time.DateTime))
			} else {
				slice2 = append(slice2, "N/A")
			}
			if err := f.SetCellStyle(sheet, fmt.Sprintf("C%d", startRow+j), fmt.Sprintf("D%d", startRow+j), cellStyle); err != nil {
				log.Println(err)
			}
			cell, err := excelize.CoordinatesToCellName(3, startRow+j)
			if err != nil {
				log.Println(err)
			}
			f.SetSheetRow(sheet, cell, &slice2)

		}
		startRow += len(dataMarcaciones[i].Marcaciones)
	}

	return
}
func mergeCells() {

}

func (r *reporteGenerator) CreateSheetEmploye(items []_r.Asistencia, sheet string, f *excelize.File, d _r.ReportInfo, lang string) (err error) {
	var (
		maxMarks        int
		maxTurnos       int
		cellStyle       int
		cellCenterStyle int
		titleStyle      int
		titleStyle2     int
		cell            string

		totalHrsWorked           time.Duration
		totalHrsWorkedInSchedule time.Duration
		totalHrsExcedentes       time.Duration
		totalHrsDelay            time.Duration
		totalHrsDelay2           time.Duration
	)

	for i := 0; i < len(items); i++ {
		if items[i].CountMarcaciones > maxMarks {
			maxMarks = items[i].CountMarcaciones
		}
		if items[i].CountTurnos > maxTurnos {
			maxTurnos = items[i].CountTurnos
		}
	}

	f.NewSheet(sheet)
	r.reporteUtil.SetUpReporteLayout(sheet, f)
	if titleStyle, err = r.reporteUtil.GetTitleStyle(f); err != nil {
		return
	}
	if titleStyle2, err = r.reporteUtil.GetTitleStyle2(f); err != nil {
		return
	}
	if cellStyle, err = r.reporteUtil.GetCommonCellStyle(f); err != nil {
		return
	}
	if cellCenterStyle, err = r.reporteUtil.GetCellCenterStyle(f); err != nil {
		return
	}

	//INFO
	if err = r.reporteUtil.SetUpHeader(sheet, f, d, titleStyle, titleStyle2, cellCenterStyle, cellStyle, lang); err != nil {
		return
	}
	if maxMarks == 0 || maxMarks == 1 {
		maxMarks = 2
	}

// 	display, tooltip := "https://github.com/xuri/excelize", "Excelize on GitHub"
// if err := f.SetCellHyperLink("Sheet1", "A3",
//     "https://github.com/xuri/excelize", "External", excelize.HyperlinkOpts{
//         Display: &display,
//         Tooltip: &tooltip,
//     }); err != nil {
//     fmt.Println(err)
// }

	f.SetColWidth(sheet, "A", "A", 5)
	f.SetColWidth(sheet, "B", "B", 13)

	f.SetColWidth(sheet, "C", "C", float64(8*maxMarks))
	f.SetColWidth(sheet, "D", "D", 22)
	f.SetColWidth(sheet, "E", "J", 18)

	headers := []string{r.locale.MustLocalize("Date", lang)}
	headers = append(headers, r.locale.MustLocalize("Markings", lang))
	headers = append(headers, r.locale.MustLocalize("Schedule", lang))
	headers = append(headers, r.locale.MustLocalize("TotalHours", lang))
	headers = append(headers, r.locale.MustLocalize("TotalHoursWorkedInSchedule", lang))
	headers = append(headers, r.locale.MustLocalize("HoursWorked", lang))
	headers = append(headers, r.locale.MustLocalize("ExcessHours", lang))
	headers = append(headers, r.locale.MustLocalize("Delay", lang))
	headers = append(headers, r.locale.MustLocalize("Delay2", lang))

	if err = f.SetCellStyle(sheet, "B6", "J6", titleStyle); err != nil {
		log.Println(err)
		return
	}

	cell, err = excelize.CoordinatesToCellName(2, 6)
	if err != nil {
		return
	}
	f.SetSheetRow(sheet, cell, &headers)

	if err != nil {
		return
	}

	startRow := 7

	for idx, c := range items {

		hrsTrabajadas := time.Second * time.Duration(c.HrsTrabajadas)
		hrsTrabajadasEnHorario := time.Second * time.Duration(c.HrsTrabajadasEnHorario)
		hrsExcedentes := time.Second * time.Duration(c.HrsExcedentes)
		hrsDelay := time.Second * time.Duration(c.Retraso)
		hrsDelay2 := time.Second * time.Duration(c.Retraso2)

		totalHrsWorked += hrsTrabajadas
		totalHrsWorkedInSchedule += hrsTrabajadasEnHorario
		totalHrsExcedentes += hrsExcedentes
		totalHrsDelay += hrsDelay
		totalHrsDelay2 += hrsDelay2

		slice := []interface{}{c.AsistenciaDate[0:10]}
		slice = append(slice, c.Marcaciones)
		slice = append(slice, c.Horario)
		slice = append(slice, _r.Timespan(time.Second*time.Duration(c.HrsTotales)).Format())
		slice = append(slice, _r.Timespan(hrsTrabajadasEnHorario).Format())
		slice = append(slice, _r.Timespan(hrsTrabajadas).Format())
		slice = append(slice, _r.Timespan(hrsExcedentes).Format())
		slice = append(slice, _r.Timespan(hrsDelay).Format())
		slice = append(slice, _r.Timespan(hrsDelay2).Format())

		if err := f.SetCellStyle(sheet, fmt.Sprintf("B%d", (idx+startRow)), fmt.Sprintf("J%d", (idx+startRow)), cellStyle); err != nil {
			log.Println(err)
		}
		cell, err := excelize.CoordinatesToCellName(2, idx+startRow)
		if err != nil {
			log.Println(err)
		}
		f.SetSheetRow(sheet, cell, &slice)
	}

	if err = r.reporteUtil.SetUpTotal(
		sheet,
		f,
		5, (startRow + len(items) + 1), titleStyle2, cellStyle,
		"E", "F", "J",
		lang,
		totalHrsWorkedInSchedule.String(),
		totalHrsWorked.String(),
		totalHrsExcedentes.String(),
		totalHrsDelay.String(),
		totalHrsDelay2.String(),
	); err != nil {
		return
	}

	return
}
