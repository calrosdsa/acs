package repository

import "time"

type ReporteMarcacionesData struct {
	MarcacionGroup []MarcacionGroup
}

type MarcacionGroup struct {
	Date        time.Time
	Marcaciones []MarcacionInterval
}

type MarcacionInterval struct {
	MarcacionE *MarcacionItem
	MarcacionS *MarcacionItem
}

type MarcacionItem struct {
	Id            int
	Date          time.Time
	TypeMarcacion int
	DateString    time.Time
}
