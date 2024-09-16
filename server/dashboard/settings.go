package dashboard

import (
	"context"
	"net/http"

	"github.com/jqwez/caselore/view/components"
	"github.com/jqwez/caselore/view/pages"
)

func RegisterSettings(mux *http.ServeMux) {
	settingsMux := http.NewServeMux()
	mux.Handle("/htmx/dashboard/settings/", http.StripPrefix("/htmx/dashboard/settings", settingsMux))
	settingsMux.HandleFunc("/", HandleSettings)
}

func HandleSettings(w http.ResponseWriter, r *http.Request) {
	pages.Settings(pages.SettingsProps{}).Render(context.Background(), w)
	crumbs := []components.Breadcrumb{{Text: "Settings", Route: "/settings"}}
	crumbProps := components.BreadcrumbProps{Breadcrumbs: crumbs}
	components.HeaderBreadcrumb(crumbProps).Render(context.Background(), w)
}
