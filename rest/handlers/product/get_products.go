package product

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

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

	/*
		wg := WaitGroup{
			noCopy: noCopy{},
			state : atomic.Unit64{
				_ : noCopy{},
				_ : align64,
				v : 0
			},
			sema : 0
		}

	*/

	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		defer wg.Done()

		cnt1, _ := h.svc.Count()
		fmt.Print(cnt1)
	}()

	// wg.Add(1)
	go func() {
		defer wg.Done()
		cnt1, _ := h.svc.Count()
		fmt.Print(cnt1)

	}()

	// wg.Add(1)
	go func() {
		defer wg.Done()
		cnt1, _ := h.svc.Count()
		fmt.Print(cnt1)
		// wg.Add(-1)
	}()

	wg.Wait() // if the hign 32 bits = 0, this goroutine will not sleep

	// time.Sleep(2 * time.Second)

	util.SendPage(w, productList, page, limit, cnt)
	go hey(&wg)
}

func hey(wg *sync.WaitGroup) {
	wg.Wait()
}
