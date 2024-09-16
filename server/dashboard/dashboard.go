package dashboard

import (
	"context"
	"net/http"

	"github.com/jqwez/caselore/view/components"
	"github.com/jqwez/caselore/view/pages"
)

func RegisterDashboard(mux *http.ServeMux) {
	dashMux := http.NewServeMux()
	mux.Handle("/htmx/dashboard/", http.StripPrefix("/htmx/dashboard", dashMux))
	dashMux.HandleFunc("/", HandleDashboardHome)
	RegisterTasks(mux)
	RegisterCases(mux)
	RegisterSettings(mux)
	RegisterToast(mux)
}

func HandleDashboardHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Dashboard HTMX route is invalid", http.StatusNotFound)
		return
	}
	pages.Index().Render(context.Background(), w)

	crumbs := []components.Breadcrumb{}
	crumbProps := components.BreadcrumbProps{Breadcrumbs: crumbs}
	components.HeaderBreadcrumb(crumbProps).Render(context.Background(), w)
}
