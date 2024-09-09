package server

import "net/http"

func RegisterPageRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", HandleRoot)
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.Write([]byte("404 Not Found"))
		return

	}
	w.Write([]byte("hello"))
}
