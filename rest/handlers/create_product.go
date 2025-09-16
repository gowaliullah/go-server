package handlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gowaliullah/ecommerce/config"
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

	if len(headerArr) != 2 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	accessToken := headerArr[1]
	tokenParts := strings.Split(accessToken, ".")

	if len(tokenParts) != 3 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	jwtHeader := tokenParts[0]
	jwtPayload := tokenParts[1]
	signature := tokenParts[2]

	message := jwtHeader + "." + jwtPayload
	cnf := config.GetConfig()

	byteArrSecret := []byte(cnf.JwtSecretKey)
	byteArrMsg := []byte(message)

	h := hmac.New(sha256.New, byteArrSecret)
	h.Write(byteArrMsg)

	hash := h.Sum(nil)
	newSignature := base64UrlEncode(hash)

	if newSignature != signature {
		http.Error(w, "Unauthorized..! tui hacker", http.StatusUnauthorized)
		return
	}

	fmt.Println("_______ token :")

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

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
