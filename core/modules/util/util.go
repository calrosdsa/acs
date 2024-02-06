package util

import (
	_r "acs/domain/repository"
)

type util struct{}

func New() _r.Util {
	return &util{}
}

func (u *util) PaginationValues(p int) (page int) {
	if p == 1 || p == 0 {
		page = 0
	} else {
		page = p - 1
	}
	return
}

func (h *util) GetNextPage(results int, pageSize int, page int) (nextPage int) {
	if results == 0 {
		nextPage = 0
	} else if results != pageSize {
		nextPage = 0
	} else {
		nextPage = page + 1
	}
	return
}
