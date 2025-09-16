package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gowaliullah/ecommerce/config"
	"github.com/gowaliullah/ecommerce/database"
	"github.com/gowaliullah/ecommerce/util"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid JSON", http.StatusBadRequest)
		return
	}

	usr := database.Find(reqLogin.Email, reqLogin.Password)
	// if err == nil {
	// 	http.Error(w, "Invalid credentials", http.StatusBadRequest)
	// }

	cnf := config.GetConfig()

	accessToken, err := util.CreateJWT(cnf.JwtSecretKey, util.Payload{
		Sub:       usr.ID,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
	})

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	fmt.Println(accessToken)
	util.SendData(w, accessToken, http.StatusOK)
}
