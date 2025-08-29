package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gowaliullah/ecommerce/database"
	"github.com/gowaliullah/ecommerce/util"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Please give the valid product id", http.StatusBadRequest)
		return
	}

	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me valid JSON data", http.StatusBadRequest)
		return
	}

	newProduct.ID = id

	database.Update(newProduct)

	util.SendData(w, "Successfully updated product", http.StatusOK)
}
