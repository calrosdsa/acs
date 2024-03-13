package repository

import (
	"bytes"
	"context"
	"time"
)

type Data struct {
	Date time.Time `json:"date"`
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
	GetAsistenciaEmployeeSitio(ctx context.Context, d ReporteRequest) (res []Asistencia, err error)
	GetAsistenciaEmployeeArea(ctx context.Context, d ReporteRequest) (res []Asistencia, err error)
	// GetAsistenciaEmployeSitio(ctx context.Context, d ReporteRequest) (res []Asistencia, err error)
	GetAsistenciaEmployee(ctx context.Context, d ReporteRequest) (res []Asistencia, err error)
	GetMarcacionesForReport(ctx context.Context, d ReporteRequest) (res []MarcacionItem, err error)
	GetEmpleadoCardHolder(ctx context.Context, guid string) (res Employee, err error)

	GetEmployeesSitio(ctx context.Context, idSitio int) (res []Employee, err error)
	GetEmployeesArea(ctx context.Context, idArea int,all bool, idSitio int) (res []Employee, err error)

	GetArea(ctx context.Context, idArea int) (res Area, err error)
	GetSitio(ctx context.Context, idSitio int) (res Sitio, err error)

	// GetReporteEmpleado(ctx context.Context) (res []Data, horario []Horario, err error)
}

type ReporteUseCase interface {
	GetReportEmploye(ctx context.Context, d ReporteRequest, buffer *bytes.Buffer) (err error)

	GetReporteSitio(ctx context.Context, d ReporteRequest, buffer *bytes.Buffer) (err error)

	GetReporteArea(ctx context.Context, d ReporteRequest, buffer *bytes.Buffer) (err error)

	GetReporte(ctx context.Context, d ReporteRequest, buffer *bytes.Buffer) (err error)

	// GetMarcacionesForReport(ctx context.Context,d ReporteRequest, buffer *bytes.Buffer)(err error)
	// GetReporteEmpleado(ctx context.Context, buffer *bytes.Buffer, d TMarcacionAsistencia) (err error)
}

type ReportInfo struct {
	EmployeName   string        `json:"employeName"`
	AreaName      string        `json:"areaName"`
	SitioName     string        `json:"sitioName"`
	From          string        `json:"from"`
	To            string        `json:"to"`
	ReporteType   ReporteType   `json:"reporteType"`
	ReporteFormat ReporteFormat `json:"reporteFormat"`
}

type ReporteRequest struct {
	CHGuid        string        `json:"cardHolderGuid"`
	Lang          string        `json:"lang"`
	StartDate     string        `json:"startDate"`
	EndDate       string        `json:"endDate"`
	ReporteType   ReporteType   `json:"reporteType"`
	ReporteFormat ReporteFormat `json:"reporteFormat"`
	// IdsSitio []int `json:"ids"`
	IdSitio   int  `json:"idSitio"`
	IdArea    int  `json:"idArea"`
	// AllAreas  bool `json:"todasLasAreas"`
	ALLSitios bool `json:"todosLosSitios"`
}

type ReporteType byte

const (
	REPORTE_EMPLOYEE = 1
	REPORTE_AREA     = 2
	REPORTE_SITIO    = 3
)

type ReporteFormat byte

const (
	EMPLOYEE_FORMAT = 1
	GENERAL_FORMAT  = 2
)
