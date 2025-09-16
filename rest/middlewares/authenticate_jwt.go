package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/gowaliullah/ecommerce/config"
)

func AuthenticateJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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

		next.ServeHTTP(w, r)
	})
}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
