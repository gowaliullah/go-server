package product

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gowaliullah/ecommerce/repo"
	"github.com/gowaliullah/ecommerce/util"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Please give the valid product id", http.StatusBadRequest)
		return
	}

	var req ReqCreatedProduct

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me valid JSON data", http.StatusBadRequest)
		return
	}

	_, err = h.productRepo.Update(repo.Product{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImgURL:      req.ImgURL,
	})

	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	util.SendData(w, "Successfully updated product", http.StatusOK)
}
