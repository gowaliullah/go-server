package middleware

import "net/http"

func CorsWithPreflight(mux *http.ServeMux) http.Handler {
	handleAllRequest := func(w http.ResponseWriter, r *http.Request) {
		// handle cors
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Habib") // preflight request wiht OPTIONS

		// handle preflight request
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}
		mux.ServeHTTP(w, r)
	}
	return http.HandlerFunc(handleAllRequest)
}
