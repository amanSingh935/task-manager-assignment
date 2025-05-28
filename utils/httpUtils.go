package utils

import (
	"net/http"
	"strconv"
	"task-manager/model"
)

func getNumericValueFromUrlParam(paramName string, r *http.Request) int {
	if p := r.URL.Query().Get(paramName); p != "" {
		if val, err := strconv.Atoi(p); err == nil && val > 0 {
			return val
		}
	}
	return -1
}

func PrepareListResponse(r *http.Request) *model.ListResponse {
	page := getNumericValueFromUrlParam("page", r)
	if page == -1 {
		page = 1
	}
	limit := getNumericValueFromUrlParam("limit", r)
	if limit == -1 {
		limit = 1
	}
	return &model.ListResponse{
		Page:  page,
		Limit: limit,
	}
}
