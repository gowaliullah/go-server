package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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

func helloHandlar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func aboutHandlar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I'm Habi.., a student")
}

type Product struct {
	ID          int     `json:"_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {

	handleCors(w)
	handlePreFlight(w, r)

	if r.Method != http.MethodGet {
		http.Error(w, "Plz give me valid request", 400)
		return
	}

	sendData(w, productList, 200)
}

func createProduct(w http.ResponseWriter, r *http.Request) {

	handleCors(w)
	handlePreFlight(w, r)

	if r.Method != http.MethodPost {
		http.Error(w, "Please give me post request", 400)
		return
	}

	var newProduct Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me valid JSON data", 400)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

	sendData(w, newProduct, 201)

}

func handleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Habib") // preflight request wiht OPTIONS
}

func handlePreFlight(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
	}
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/products", getProducts)
	mux.HandleFunc("/add-product", createProduct)

	mux.HandleFunc("/hello", helloHandlar)

	mux.HandleFunc("/about", aboutHandlar)

	fmt.Println("server running on: 8080")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println("Error: ", err)
	}

}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "It's very delicious and full of vitamin C.",
		Price:       108.00,
		ImgURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTnzv3PeLyF9-dxIj0MGIabXMKYA6CFTB-0OA&s",
	}

	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Crisp and sweet red apples.",
		Price:       120.50,
		ImgURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTnzv3PeLyF9-dxIj0MGIabXMKYA6CFTB-0OA&s",
	}

	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Rich in potassium and easy to digest.",
		Price:       60.00,
		ImgURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTnzv3PeLyF9-dxIj0MGIabXMKYA6CFTB-0OA&s",
	}

	prd4 := Product{
		ID:          4,
		Title:       "Mango",
		Description: "The king of fruits, sweet and juicy.",
		Price:       150.75,
		ImgURL:      "https://example.com/images/mango.jpg",
	}

	prd5 := Product{
		ID:          5,
		Title:       "Watermelon",
		Description: "Refreshing and hydrating summer fruit.",
		Price:       200.00,
		ImgURL:      "https://example.com/images/watermelon.jpg",
	}

	productList = append(productList, prd1, prd2, prd3, prd4, prd5)
}
