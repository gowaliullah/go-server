package product

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gowaliullah/basic-ecommerce/util"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	reqQuery := r.URL.Query()
	pageAsString := reqQuery.Get("page")
	limitAsString := reqQuery.Get("limit")

	page, _ := strconv.ParseInt(pageAsString, 10, 32)
	limit, _ := strconv.ParseInt(limitAsString, 10, 32)

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	productList, err := h.svc.List(page, limit)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	cnt, err := h.svc.Count()
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		cnt1, _ := h.svc.Count()
		fmt.Print(cnt1)
		wg.Add(-1)
	}()

	wg.Add(1)
	go func() {
		cnt1, _ := h.svc.Count()
		fmt.Print(cnt1)
		wg.Add(-1)
	}()

	wg.Add(1)
	go func() {
		cnt1, _ := h.svc.Count()
		fmt.Print(cnt1)
		wg.Add(-1)
	}()

	time.Sleep(2 * time.Second)

	util.SendPage(w, productList, page, limit, cnt)
}
