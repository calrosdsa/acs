package repository

type Locale interface {
	MustLocalize(id string,lang string) (res string)
}
