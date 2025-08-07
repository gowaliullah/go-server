package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gowaliullah/ecommerce/middleware"
	"github.com/gowaliullah/ecommerce/models"
)

var productList []models.Product

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct models.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me valid JSON data", 400)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

	middleware.SendData(w, newProduct, 201)

}
