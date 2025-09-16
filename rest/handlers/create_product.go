package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gowaliullah/ecommerce/database"
	"github.com/gowaliullah/ecommerce/util"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	header := r.Header.Get("Authorization")

	if header == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	headerArr := strings.Split(header, " ")

	if len(header) != 2 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	fmt.Println("_______ token :", headerArr)

	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me valid JSON data", 400)
		return
	}

	createdProduct := database.Store(newProduct)

	util.SendData(w, createdProduct, http.StatusCreated)

}
