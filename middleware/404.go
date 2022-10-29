package middleware

import (
	"fmt"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Page does not exist - 404")
	return
}
