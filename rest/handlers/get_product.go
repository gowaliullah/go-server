package handlers

import (
	"net/http"
	"strconv"

	"github.com/gowaliullah/ecommerce/database"
	"github.com/gowaliullah/ecommerce/util"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Please give the valid product id", 400)
		return
	}

	product := database.Get(id)
	if product == nil {
		util.SendError(w, http.StatusNotFound, "Product not found")
		return
	}
	util.SendData(w, product, http.StatusOK)
}
