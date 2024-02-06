package repository

type Util interface {
	PaginationValues(page int) int
	GetNextPage(results int, pageSize int, page int) (nextPage int)
}