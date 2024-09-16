package dashboard

import (
	"context"
	"net/http"

	"github.com/jqwez/caselore/model"
	"github.com/jqwez/caselore/view/components"
	"github.com/jqwez/caselore/view/pages"
)

func RegisterTasks(mux *http.ServeMux) {
	tasksMux := http.NewServeMux()
	mux.Handle("/htmx/dashboard/tasks/", http.StripPrefix("/htmx/dashboard/tasks", tasksMux))
	tasksMux.HandleFunc("/", HandleTasks)
}

func HandleTasks(w http.ResponseWriter, r *http.Request) {
	casesProps := pages.CasesProps{
		Cases: []model.Case{
			{Idea: "nothing", Title: "Sample", Description: "hello"},
		},
	}
	pages.Cases(casesProps).Render(context.Background(), w)

	crumbs := []components.Breadcrumb{{Text: "Tasks", Route: "/tasks"}}
	crumbProps := components.BreadcrumbProps{Breadcrumbs: crumbs}
	components.HeaderBreadcrumb(crumbProps).Render(context.Background(), w)
}
