package repository

import (
	"bytes"
	"github.com/xuri/excelize/v2"
)


type ReporteGenerator interface {
	GenerateReporteEmploye(asistencias []Asistencia, buffer *bytes.Buffer,lang string)(err error)
}

type ReporteUtil interface {
	SetUpReporteLayout(sheet string,f *excelize.File)(err error)
	GetBlankStyle(f *excelize.File)(styleId int,err error)
	GetTitleStyle(f *excelize.File)(styleId int,err error)
	GetCommonCellStyle(f *excelize.File)(styleId int,err error)
	SetUpHeader(sheet string,f *excelize.File,d ReportInfo,titleStyle,cellStyle int,lang string)(err error)
}