package product

import (
	"net/http"
	"strconv"

	"github.com/gowaliullah/basic-ecommerce/domain"
	"github.com/gowaliullah/basic-ecommerce/util"
)

type Pagination struct {
	Data       []*domain.Product `json:"data"`
	Page       int64             `json:"page"`
	Limit      int64             `json:"limit"`
	TotalItems int64             `json:"totalItems"`
	TotalPages int64             `json:"totalPages"`
}

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	reqQuery := r.URL.Query()
	pageAsString := reqQuery.Get("page")
	limitAsString := reqQuery.Get("limit")

	page, _ := strconv.ParseInt(pageAsString, 10, 32)
	limit, _ := strconv.ParseInt(limitAsString, 10, 32)

	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	productList, err := h.svc.List(page, limit)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	util.SendData(w, productList, http.StatusOK)
}
