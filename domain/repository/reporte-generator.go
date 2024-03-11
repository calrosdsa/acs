package repository

import (
	"bytes"
	// "time"

	"github.com/xuri/excelize/v2"
)


type EmployeeAsistencia struct {
	Employee Employee	
	Asistencias []Asistencia
}

type ReporteGenerator interface {
	GenerateReporteEmploye(info ReportInfo ,asistencias []Asistencia, dataMarcaciones []MarcacionGroup,
		buffer *bytes.Buffer, lang string) (err error)
	//REPORTE QUE MUESTRA TODOS LAS ASISTENCIAS EN UNA SOLA HOJA DE EXCEL	
	GenerateReporteSitioGeneral(info ReportInfo,asistencia []Asistencia, buffer *bytes.Buffer, lang string)(err error)
	GenerateReporteAreaGeneral(info ReportInfo,asistencia []Asistencia, buffer *bytes.Buffer, lang string)(err error)

	//REPORTE QUE MUESTRA LAS ASISTENCIAS DE LOS USUARIOS EN DIFERENTES HOJAS 
	GenerateReporteSitioEmployees(info ReportInfo,asistencia []EmployeeAsistencia, buffer *bytes.Buffer, lang string)(err error)
	GenerateReporteAreaEmployees(info ReportInfo,asistencia []EmployeeAsistencia, buffer *bytes.Buffer, lang string)(err error)

}

type ReporteUtil interface {
	SetUpReporteLayout(sheet string, f *excelize.File) (err error)
	GetBlankStyle(f *excelize.File) (styleId int, err error)
	GetTitleStyle(f *excelize.File) (styleId int, err error)
	GetTitleStyle2(f *excelize.File) (styleId int, err error)
	GetCommonCellStyle(f *excelize.File) (styleId int, err error)
	GetCellCenterStyle(f *excelize.File) (styleId int, err error)
	SetUpHeader(sheet string, f *excelize.File, d ReportInfo, titleStyle, titleStyle2, cellCenterStyle, cellStyle int, lang string) (err error)

	SetUpTotal(sheet string, f *excelize.File, startCol, startRow, titleStyle,
		cellStyle int, col1, col2, col3, lang string, args ...interface{}) (err error)
}
