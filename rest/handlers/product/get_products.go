package product

import (
	"net/http"

	"github.com/gowaliullah/basic-ecommerce/util"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	productList, err := h.svc.List()
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	util.SendData(w, productList, http.StatusOK)
}
