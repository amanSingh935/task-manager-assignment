package model

type ListResponse struct {
	Page           int    `json:"page"`
	Limit          int    `json:"limit"`
	Total          int    `json:"total"`
	PaginatedTasks []Task `json:"data"`
}
