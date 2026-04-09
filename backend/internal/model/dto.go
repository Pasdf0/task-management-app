package model

type PaginatedTasks struct {
	Data []Task     `json:"data"`
	Meta Pagination `json:"meta"`
}

type Pagination struct {
	TotalItems  int64 `json:"totalItems"`
	TotalPages  int   `json:"totalPages"`
	CurrentPage int   `json:"currentPage"`
	PageSize    int   `json:"pageSize"`
}
