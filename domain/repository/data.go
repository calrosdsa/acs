package repository

import "time"

type ResponseMessage struct {
	Message string `json:"message"`
}

type Timespan time.Duration

func (t Timespan) Format() string {
	return time.Unix(0, 0).UTC().Add(time.Duration(t)).Format("15:04:05")
}
