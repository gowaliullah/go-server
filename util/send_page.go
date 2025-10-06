package util

import "net/http"

type Pagination struct {
	Page       int64 `json:"page"`
	Limit      int64 `json:"limit"`
	TotalItems int64 `json:"totalItems"`
	TotalPages int64 `json:"totalPages"`
}

type paginatedData struct {
	Data       any        `json:"data"`
	Pagination Pagination `json:"pagination"`
}

func SendPage(w http.ResponseWriter, page paginatedData) {
	SendData(w, page, http.StatusOK)
}
