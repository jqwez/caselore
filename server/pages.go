package server

import (
	"context"
	"net/http"

	"github.com/jqwez/caselore/view/pages"
)

func RegisterPageRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", HandleRoot)
	mux.HandleFunc("/dashboard", HandleDashboard)
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.Write([]byte("404 Not Found"))
		return

	}

	pages.Index().Render(context.Background(), w)
}

func HandleDashboard(w http.ResponseWriter, r *http.Request) {

	pages.Dashboard().Render(context.Background(), w)
}
