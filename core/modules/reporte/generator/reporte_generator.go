package generator

import (
	_r "acs/domain/repository"
	"fmt"
	"log"
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

func (r *reporteGenerator) GenerateReporteEmploye(asistencias []_r.Asistencia, buffer *bytes.Buffer, lang string) (err error) {
	f, err := excelize.OpenFile("./media/template.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
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
	if err = r.CreateSheetEmploye(asistencias, "sheet1", f, d, lang); err != nil {
		return
	}

	err = f.Write(buffer)
	if err != nil {
		log.Println(err)
	}
	return
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
		totalHrsDelay            time.Duration
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

	// f := excelize.NewFile()
	//     defer func() {
	//         if err := f.Close(); err != nil {
	//             fmt.Println(err)
	//         }
	//     }()
	// Insert a picture.

	// Insert a picture scaling in the cell with location hyperlink.

	// f.AddPicture()

	//Table

	f.SetColWidth(sheet, "A", "A", 5)
	f.SetColWidth(sheet, "B", "B", 13)
	f.SetColWidth(sheet, "C", "C", 25)
	f.SetColWidth(sheet, "D", "D", float64((11 * maxTurnos)))
	f.SetColWidth(sheet, "E", "H", 18)

	headers := []string{r.locale.MustLocalize("Date", lang)}
	headers = append(headers, r.locale.MustLocalize("Markings", lang))
	headers = append(headers, r.locale.MustLocalize("Schedule", lang))
	headers = append(headers, r.locale.MustLocalize("TotalHours", lang))
	headers = append(headers, r.locale.MustLocalize("TotalHoursWorkedInSchedule", lang))
	headers = append(headers, r.locale.MustLocalize("HoursWorked", lang))
	headers = append(headers, r.locale.MustLocalize("Delay", lang))


	

	// set style for the 'SUNDAY' to 'SATURDAY'
	if err = f.SetCellStyle(sheet, "B6", "H6", titleStyle); err != nil {
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
		hrsDelay := time.Second * time.Duration(c.Retraso)

		totalHrsWorked += hrsTrabajadas
		totalHrsWorkedInSchedule += hrsTrabajadasEnHorario
		totalHrsDelay += hrsDelay

		slice := []interface{}{c.AsistenciaDate[0:10]}
		slice = append(slice, c.Marcaciones)
		slice = append(slice, c.Horario)
		slice = append(slice, _r.Timespan(time.Second*time.Duration(c.HrsTotales)).Format())
		slice = append(slice, _r.Timespan(hrsTrabajadasEnHorario).Format())
		slice = append(slice, _r.Timespan(hrsTrabajadas).Format())
		slice = append(slice, _r.Timespan(hrsDelay).Format())
		if err := f.SetCellStyle(sheet, fmt.Sprintf("B%d", (idx+startRow)), fmt.Sprintf("H%d", (idx+startRow)), cellStyle); err != nil {
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
		totalHrsWorked,
		totalHrsWorkedInSchedule,
		totalHrsDelay, 5, (startRow + len(items) + 1), titleStyle2, cellStyle,
		"E", "F", "G", "H",
		lang,
	); err != nil {
		return
	}

	return
}
