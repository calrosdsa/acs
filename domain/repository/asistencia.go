package repository

import "context"

type Asistencia struct {
	Id                     int     `json:"id"`
	AsistenciaDate         string  `json:"asistenciaDate"`
	CardHolderGuid         string  `json:"cardholderGuid"`
	Retraso                float64 `json:"retraso"`
	HrsTotales             float64 `json:"hrsTotales"`
	HrsTrabajadas          float64 `json:"hrsTrabajadas"`
	HrsTrabajadasEnHorario float64 `json:"hrsTrabajadasEnHorario"`
	Marcaciones            string  `json:"marcaciones"`
	Horario                string  `json:"horario"`
	CountTurnos            int     `json:"countTurnos"`
	CountMarcaciones       int     `json:"countMarcaciones"`
}

type AsistenciaRepository interface {
	GetAsistencia(ctx context.Context, chGuid string, fecha string) (res Asistencia, err error)
	GetAsistenciasUser(ctx context.Context, chGuid string, page, size int) (res []Asistencia, count int, err error)

	CreateAsistencia(ctx context.Context, d Asistencia) (err error)
	UpdateAsistencia(ctx context.Context, d Asistencia) (err error)
	ExistAsistencia(ctx context.Context, chguid string, fecha string) (res bool, err error)

	//Get user information markings
	GetEmployeData(ctx context.Context, chGuid string, fecha string) (res Data, horario []Horario, err error)

	//ONLY FOR DEVELOPMENT
	InsertMarcacion(ctx context.Context, d TMarcacionAsistencia) (err error)
}

type AsistenciaUseCase interface {
	GetAsistenciasUser(ctx context.Context, chGuid string, page, size int) (res []Asistencia, nextPage int, count int, err error)
	GetAsistencia(ctx context.Context, chguid string, fecha string) (res Asistencia, err error)

	CreateOrUpdateAsistencia(ctx context.Context, d Asistencia) (err error)
	UpdateAsistenciaFromIncomingData(ctx context.Context, d TMarcacionAsistencia) (err error)

	//ONLY FOR DEVELOPMENT
	InsertMarcacion(ctx context.Context, d TMarcacionAsistencia) (err error)
}
