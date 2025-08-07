package handlers

import (
	"fmt"
	"net/http"
)

func AboutHandlar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I'm Habi.., a student")
}
