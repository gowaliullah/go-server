package main

import (
	"github.com/gowaliullah/ecommerce/cmd"
	"github.com/gowaliullah/ecommerce/database"
)

func main() {
	cmd.Serve()

}

func init() {
	prd1 := database.Product{
		ID:          1,
		Title:       "Orange",
		Description: "It's very delicious and full of vitamin C.",
		Price:       108.00,
		ImgURL:      "https://example.com/images/orange.jpg",
	}

	prd2 := database.Product{
		ID:          2,
		Title:       "Apple",
		Description: "Crisp and sweet red apples.",
		Price:       120.50,
		ImgURL:      "https://example.com/images/apple.jpg",
	}

	prd3 := database.Product{
		ID:          3,
		Title:       "Banana",
		Description: "Rich in potassium and easy to digest.",
		Price:       60.00,
		ImgURL:      "https://example.com/images/banana.jpg",
	}

	prd4 := database.Product{
		ID:          4,
		Title:       "Mango",
		Description: "The king of fruits, sweet and juicy.",
		Price:       150.75,
		ImgURL:      "https://example.com/images/mango.jpg",
	}

	prd5 := database.Product{
		ID:          5,
		Title:       "Watermelon",
		Description: "Refreshing and hydrating summer fruit.",
		Price:       200.00,
		ImgURL:      "https://example.com/images/watermelon.jpg",
	}

	database.ProductList = append(database.ProductList, prd1, prd2, prd3, prd4, prd5)
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

	entity gulai muloto resourse hoy

*/
