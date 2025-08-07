package handlers

import (
	"net/http"

	"github.com/gowaliullah/ecommerce/middleware"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	middleware.SendData(w, productList, 200)
}
