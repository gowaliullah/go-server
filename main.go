package main

import (
	"fmt"
	"net/http"

	"github.com/gowaliullah/ecommerce/handlers"
	"github.com/gowaliullah/ecommerce/middleware"
	"github.com/gowaliullah/ecommerce/models"
)

var productList []models.Product

func main() {
	mux := http.NewServeMux()

	muxRouter := middleware.GlobalRouter(mux)

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("GET /hello", http.HandlerFunc(handlers.HelloHandlar))
	mux.Handle("GET /about", http.HandlerFunc(handlers.AboutHandlar))
	mux.Handle("POST /add-product", http.HandlerFunc(handlers.CreateProduct))

	fmt.Println("server running on: 8080")

	err := http.ListenAndServe(":8080", muxRouter)

	if err != nil {
		fmt.Println("Error: ", err)
	}

}

func init() {
	prd1 := models.Product{
		ID:          1,
		Title:       "Orange",
		Description: "It's very delicious and full of vitamin C.",
		Price:       108.00,
		ImgURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTnzv3PeLyF9-dxIj0MGIabXMKYA6CFTB-0OA&s",
	}

	prd2 := models.Product{
		ID:          2,
		Title:       "Apple",
		Description: "Crisp and sweet red apples.",
		Price:       120.50,
		ImgURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTnzv3PeLyF9-dxIj0MGIabXMKYA6CFTB-0OA&s",
	}

	prd3 := models.Product{
		ID:          3,
		Title:       "Banana",
		Description: "Rich in potassium and easy to digest.",
		Price:       60.00,
		ImgURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTnzv3PeLyF9-dxIj0MGIabXMKYA6CFTB-0OA&s",
	}

	prd4 := models.Product{
		ID:          4,
		Title:       "Mango",
		Description: "The king of fruits, sweet and juicy.",
		Price:       150.75,
		ImgURL:      "https://example.com/images/mango.jpg",
	}

	prd5 := models.Product{
		ID:          5,
		Title:       "Watermelon",
		Description: "Refreshing and hydrating summer fruit.",
		Price:       200.00,
		ImgURL:      "https://example.com/images/watermelon.jpg",
	}

	productList = append(productList, prd1, prd2, prd3, prd4, prd5)
}

/*

	1. network interface card --> NIC
	2. socket received buffer
	3. write buffer
	3. electronic magnify
	4. file descriptor --> 0 < 1, 2, 3, 4...........

	6. router - wifi adaptar - NIC - write buffer - interapct kurnel - copy write buffer and all thing is stored in socket received buffer
	send buffer kurnel niye read buffer ar kase dai

	NIC  electromagnatic kore router ar kase patai

*/
