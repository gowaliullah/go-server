package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		fmt.Println("ami middleware =age= print hobo")

		next.ServeHTTP(w, r)
		fmt.Println("ami middleware =sese= print hobo")
		fmt.Println(r.Method, r.URL.Path, time.Since(start))

	})
}
