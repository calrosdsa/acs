package repository

import "context"

type Asistencia struct {
	Id                     int     `json:"id"`
	AsistenciaDate         string  `json:"asistenciaDate"`
	CardHolderGuid         string  `json:"cardholderGuid"`
	Retraso                float64 `json:"retraso"`
	Retraso2               float64 `json:"retraso2"`
	HrsTotales             float64 `json:"hrsTotales"`
	HrsTrabajadas          float64 `json:"hrsTrabajadas"`
	HrsExcedentes          float64 `json:"hrsExcedentes"`
	HrsTrabajadasEnHorario float64 `json:"hrsTrabajadasEnHorario"`
	Marcaciones            string  `json:"marcaciones"`
	Horario                string  `json:"horario"`
	CountMarcaciones       int     `json:"countMarcaciones"`
	IdSitio                int     `json:"idSitio"`
	IdArea                 int     `json:"idArea"`
	DoorGuid               string  `json:"doorGuid"`

	EmployeeFirstName string `json:"employeeFirstName"`
	EmployeeLastName  string `json:"employeeLastName"`
	AreaName          string `json:"areaName"`
	SitioName         string `json:"sitioName"`
}

type CardHolderUser struct {
	Guid string
	IdArea int
	IdSitio int
}

type AsistenciaRepository interface {
	GetAsistencia(ctx context.Context, chGuid string, fecha string) (res Asistencia, err error)
	GetAsistenciasUser(ctx context.Context, chGuid string, page, size int) (res []Asistencia, count int, err error)

	GetAllCardHolders(ctx context.Context) (res []CardHolderUser,err error)

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
	RevocerAsistenciaAllUsers(ctx context.Context, fecha string) (err error)

	//ONLY FOR DEVELOPMENT
	InsertMarcacion(ctx context.Context, d TMarcacionAsistencia) (err error)
}

type TMarcacionAsistencia struct {
	Id              int    `json:"id"`
	AccessPointGuid string `json:"accessPointGuid"`
	CardHolderGuid  string `json:"cardHolderGuid"`
	CredentialGuid  string `json:"credentialGuid"`
	DoorGuid        string `json:"doorGuid"`
	EventType       string `json:"eventType"`
	Fecha           string `json:"fecha"`
	IdZona          int    `json:"idZona"`
	TypeMarcacion   int    `json:"typeMarcacion"`

	// IdPerfil int `json:"idPerfil"`
	IdArea  int `json:"idArea"`
	IdSitio int `json:"IdSitio"`
}

