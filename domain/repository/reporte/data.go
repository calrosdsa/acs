package reporte

import (
	"bytes"
	"context"
	"time"
)

type Data struct {
	Date  string `json:"date"`
	// Json string `json:"dataJson"`
	// MarcacionData MarcacionData `json:"marcacion_data"`
	Times []string `json:"times"`
	Types []string `json:"types"`

	TimesString string 
	TypesString string `json:"typesString"`

	HorasTrabajadas []time.Duration `json:"horas_trabajadas"`
	Horario []Horario `json:"horario"`
	Total  time.Duration 
	TotalHrsWorked time.Duration
	Retraso time.Duration

	FirstM *string
	LastM *string
	FirstT *string
	LastT *string
	// Hora []Hora
}

type Hora struct {
	Hora int `json:""`
	Marcaciones []Marcacion `json:"marcaciones"`
}

type Horario struct {
	StartTime time.Time `json:"start"`
	EndTime time.Time `json:"end"`
	Day int `json:"day"`
}

type Marcacion struct {
	Date time.Time `json:"date"`
	Type int `json:"type"`
}

type Interval struct {
	Start string `json:"start"`
	End string `json:"end"`
}

type MarcacionData struct {
	Times []string `json:"time"`
	Types []int `json:"type"`
}

type ReporteRepo interface {
	GetReporteEmpleado(ctx context.Context)(res []Data,horario []Horario,err error)
}

type ReporteUseCase interface {
	GetReporteEmpleado(ctx context.Context,buffer *bytes.Buffer)(err error)
}