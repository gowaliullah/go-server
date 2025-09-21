package product

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gowaliullah/basic-ecommerce/domain"
	"github.com/gowaliullah/basic-ecommerce/util"
)

type ReqCreatedProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"image_url"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req ReqCreatedProduct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me valid JSON data", 400)
		return
	}

	createdProduct, err := h.svc.Create(domain.Product{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImageUrl:    req.ImgURL,
	})

	if err != nil {
		http.Error(w, "Internal server err", http.StatusInternalServerError)
	}

	util.SendData(w, createdProduct, http.StatusCreated)

}
