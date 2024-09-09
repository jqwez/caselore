package server

import "net/http"

func RegisterAPI(mux *http.ServeMux) {
	apiMux := http.NewServeMux()
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", apiMux))

	apiMux.HandleFunc("/", HandleRootApi)
}

func HandleRootApi(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.Write([]byte("{}"))
		return
	}
	w.Write([]byte("testerooski"))
}
