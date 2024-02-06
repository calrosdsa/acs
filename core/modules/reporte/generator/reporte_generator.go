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
}

func New(reporteUtil _r.ReporteUtil, logger _r.Logger) _r.ReporteGenerator {
	return &reporteGenerator{
		reporteUtil: reporteUtil,
		logger:      logger,
	}
}

func (r *reporteGenerator) GenerateReporteEmploye(asistencias []_r.Asistencia, buffer *bytes.Buffer,lang string) (err error) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	d := _r.ReportInfo{
		EmployeName:  "Jorge Daniel Miranda Lopez",
		GerenciaName: "Recursos Humanos",
		SitioName:    "Equipetrol",
		From:         "01-01-2024",
		To:           "01-02-2024",
	}
	if err = r.CreateSheetEmploye(asistencias, "sheet1", f, d,lang); err != nil {
		return
	}

	err = f.Write(buffer)
	if err != nil {
		log.Println(err)
	}
	return
}

func (r *reporteGenerator) CreateSheetEmploye(items []_r.Asistencia, sheet string, f *excelize.File, d _r.ReportInfo,lang string) (err error) {
	var (
		maxMarks   int
		maxTurnos  int
		cellStyle  int
		titleStyle int
		cell       string
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
	if cellStyle, err = r.reporteUtil.GetCommonCellStyle(f); err != nil {
		return
	}

	//INFO
	if err = r.reporteUtil.SetUpHeader(sheet, f, d, titleStyle, cellStyle,lang); err != nil {
		return
	}

	//Table
	f.SetColWidth(sheet, "A", "A", 5)
	f.SetColWidth(sheet, "B", "B", 13)
	f.SetColWidth(sheet, "C", "C", float64((7 * maxMarks)))
	f.SetColWidth(sheet, "D", "D", float64((11 * maxTurnos)))
	f.SetColWidth(sheet, "E", "H", 16)

	headers := []string{"Fecha"}
	headers = append(headers, "Marcaciones")
	headers = append(headers, "Horario")
	headers = append(headers, "Hrs. Trabajadas")
	headers = append(headers, "Hrs. Total")
	headers = append(headers, "Hrs. Trabajadas 2")
	headers = append(headers, "Hrs. Retraso")

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
		slice := []interface{}{c.AsistenciaDate[0:10]}
		slice = append(slice, c.Marcaciones)
		slice = append(slice, c.Horario)
		slice = append(slice, (time.Second * time.Duration(c.HrsTrabajadas)))
		slice = append(slice, (time.Second * time.Duration(c.HrsTotales)))
		slice = append(slice, (time.Second * time.Duration(c.HrsTrabajadasEnHorario)))
		slice = append(slice, (time.Second * time.Duration(c.Retraso)))
		if err := f.SetCellStyle(sheet, fmt.Sprintf("B%d", (idx+startRow)), fmt.Sprintf("H%d", (idx+startRow)), cellStyle); err != nil {
			log.Println(err)
		}
		cell, err := excelize.CoordinatesToCellName(2, idx+startRow)
		if err != nil {
			log.Println(err)
		}
		f.SetSheetRow(sheet, cell, &slice)
	}
	return
}
