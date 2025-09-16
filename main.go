package main

import (
	"fmt"

	"github.com/gowaliullah/ecommerce/cmd"
	"github.com/gowaliullah/ecommerce/util"
)

func main() {

	cmd.Serve()
	// practice()

}

func practice() {

	jwt, err := util.CreateJWT("my-secret", util.Payload{
		Sub:         45,
		FirstName:   "Habiba",
		LastName:    "akar",
		Email:       "habiba@gmail.com",
		IsShopOwner: false,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(jwt)

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
