package handlers

import (
	"fmt"
	"net/http"
)

func HelloHandlar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
