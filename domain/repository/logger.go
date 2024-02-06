package repository


type Logger interface {
	LogError(method string,file string,err error)
}