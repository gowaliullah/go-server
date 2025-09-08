package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gowaliullah/ecommerce/database"
	"github.com/gowaliullah/ecommerce/util"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser database.user
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid JSON", http.StatusBadRequest)
		return
	}

	createdUser := database.Store(newUser)

	util.SendData(w, createdUser, http.StatusCreated)

}
