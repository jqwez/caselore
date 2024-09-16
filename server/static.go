package server

import (
	"net/http"
)

func RegisterStatic(mux *http.ServeMux) {
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))
}
