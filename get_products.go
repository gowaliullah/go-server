package main

import "net/http"

func getProducts(w http.ResponseWriter, r *http.Request) {
	sendData(w, productList, 200)
}
