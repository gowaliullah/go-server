package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func practice_needed_topic() {

	var s string
	s = "0h"
	byteArr := []byte(s)

	fmt.Println(s)
	fmt.Println(byteArr)

	enc := base64.URLEncoding.WithPadding(base64.NoPadding)
	b64Str := enc.EncodeToString(byteArr)

	fmt.Println(b64Str)

	data := []byte("Hello")
	hash := sha256.Sum256(data)
	fmt.Println(hash)

	secret := []byte("my-secret")
	message := []byte("Hello World")

	h := hmac.New(sha256.New, secret)
	h.Write(message)

	text := h.Sum(nil)

	fmt.Println(text)

}
