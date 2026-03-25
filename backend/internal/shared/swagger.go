package shared

import (
	_ "embed"
	"net/http"
)

//go:embed swagger_ui.html
var swaggerUIHTML []byte

func swaggerUIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(swaggerUIHTML)
}

func openapiSpecHandler(spec []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/yaml")
		w.Write(spec)
	}
}
