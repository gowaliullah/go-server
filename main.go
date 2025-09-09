package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {

	// cmd.Serve()
	practice()

}

func practice() {

	/*
		var s string
		s = "0h"
		byteArr := []byte(s)

		fmt.Println(s)
		fmt.Println(byteArr)

		enc := base64.URLEncoding.WithPadding(base64.NoPadding)
		b64Str := enc.EncodeToString(byteArr)

		fmt.Println(b64Str)
	*/

	data := []byte("Hello")
	hash := sha256.Sum256(data)
	fmt.Println(hash)

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
