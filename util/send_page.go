package util

import (
	"net/http"
)

type Pagination struct {
	Page       int64 `json:"page"`
	Limit      int64 `json:"limit"`
	TotalItems int64 `json:"totalItems"`
	TotalPages int64 `json:"totalPages"`
}

type PaginatedData struct {
	Data       any        `json:"data"`
	Pagination Pagination `json:"pagination"`
}

func SendPage(w http.ResponseWriter, data any, page, limit, cnt int64) {

	pagData := PaginatedData{
		Data: data,
		Pagination: Pagination{
			Page:       page,
			Limit:      limit,
			TotalItems: cnt,
			TotalPages: cnt / limit,
		},
	}

	SendData(w, pagData, http.StatusOK)
}
