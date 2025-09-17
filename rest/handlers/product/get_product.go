package product

import (
	"net/http"
	"strconv"

	"github.com/gowaliullah/ecommerce/util"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Please give the valid product id", http.StatusBadRequest)
		return
	}

	product, err := h.productRepo.Get(id)
	if product == nil {
		util.SendError(w, http.StatusNotFound, "Product not found")
		return
	}
	util.SendData(w, product, http.StatusOK)
}
