package repository

import (
	"bytes"
	"context"
	"time"
)

type Data struct {
	Date string `json:"date"`
	// Json string `json:"dataJson"`
	// MarcacionData MarcacionData `json:"marcacion_data"`
	Times []string `json:"times"`
	Types []string `json:"types"`

	TimesString string
	TypesString string `json:"typesString"`

	HorasTrabajadas          []time.Duration `json:"horas_trabajadas"`
	Horario                  []Horario       `json:"horario"`
	HorasRestantes           time.Duration   `json:"horas_restantes"`
	HorasAsignadas           time.Duration
	TotalHrsWorkedInSchedule time.Duration
	Retraso                  time.Duration
	Retraso2                 time.Duration

	FirstM *string
	LastM  *string
	FirstT *string
	LastT  *string
	// Hora []Hora
}

type Hora struct {
	Hora        int         `json:""`
	Marcaciones []Marcacion `json:"marcaciones"`
}

type Horario struct {
	StartTime time.Time `json:"start"`
	EndTime   time.Time `json:"end"`
	Day       int       `json:"day"`
}

type Marcacion struct {
	Date time.Time `json:"date"`
	Type int       `json:"type"`
}

type Interval struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type MarcacionData struct {
	Times []string `json:"time"`
	Types []int    `json:"type"`
}

type ReporteRepo interface {
	// GetReporteEmpleado(ctx context.Context) (res []Data, horario []Horario, err error)
	GetReportEmploye(ctx context.Context, d ReporteRequest) (res []Asistencia, err error)
	GetMarcacionesForReport(ctx context.Context,d ReporteRequest)(res []MarcacionItem,err error)

	// GetReporteEmpleado(ctx context.Context) (res []Data, horario []Horario, err error)
}

type ReporteUseCase interface {
	GetReportEmploye(ctx context.Context, d ReporteRequest, buffer *bytes.Buffer) (err error)
	// GetMarcacionesForReport(ctx context.Context,d ReporteRequest, buffer *bytes.Buffer)(err error)
	// GetReporteEmpleado(ctx context.Context, buffer *bytes.Buffer, d TMarcacionAsistencia) (err error)
}

type ReportInfo struct {
	EmployeName  string `json:"employeName"`
	GerenciaName string `json:"gerenciaName"`
	SitioName    string `json:"sitioName"`
	From         string `json:"from"`
	To           string `json:"to"`
}

type ReporteRequest struct {
	CHGuid string `json:"cardHolderGuid"`
	Lang   string `json:"lang"`
	StartDate string `json:"start_date"`
	EndDate string `json:"end_date"`
	
	
}
