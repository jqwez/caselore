package dashboard

import (
	"context"
	"net/http"

	"github.com/jqwez/caselore/database"
	"github.com/jqwez/caselore/model"
	"github.com/jqwez/caselore/view/components"
	"github.com/jqwez/caselore/view/pages"
)

func RegisterCases(mux *http.ServeMux) {
	caseMux := http.NewServeMux()
	mux.Handle("/htmx/dashboard/cases/", http.StripPrefix("/htmx/dashboard/cases", caseMux))
	caseMux.HandleFunc("/", HandleCases)
	caseMux.HandleFunc("/id/{id}", HandleCaseHome)
	caseMux.HandleFunc("/editrowstart/{id}", HandleCaseRowEditStart)
	caseMux.HandleFunc("/editrowfinish", HandleCaseRowEditFinish)
}

func HandleCases(w http.ResponseWriter, r *http.Request) {
	casesProps := pages.CasesProps{
		Cases: []model.Case{
			{Idea: "nothing", Title: "Sample", Description: "hello"},
			{Idea: "Something else", Title: "Sammmple 2", Description: "Really long description about some stuff over here."},
		},
	}
	pages.Cases(casesProps).Render(context.Background(), w)

	crumbs := []components.Breadcrumb{{Text: "Cases", Route: "/cases"}}
	crumbProps := components.BreadcrumbProps{Breadcrumbs: crumbs}
	components.HeaderBreadcrumb(crumbProps).Render(context.Background(), w)
}

func HandleCaseHome(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	_ = database.NewRepository()

	pages.CaseHome(id).Render(context.Background(), w)
	crumbs := []components.Breadcrumb{
		{Text: "Cases", Route: "/cases"},
		{Text: "Case Title", Route: "/cases/id/none"}}
	crumbProps := components.BreadcrumbProps{Breadcrumbs: crumbs}
	components.HeaderBreadcrumb(crumbProps).Render(context.Background(), w)
}

func HandleCaseRowEditFinish(w http.ResponseWriter, r *http.Request) {

	c := model.Case{
		Title:       "Some TExt",
		Description: "Some other Text",
	}
	pages.CaseRow(1, c).Render(context.Background(), w)
}

func HandleCaseRowEditStart(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		id = "no-id-here"
	}
	c := model.Case{
		Idea:        id,
		Title:       id,
		Description: "Some aaaother Text",
	}
	pages.CaseRowEdit(1, c).Render(context.Background(), w)
}
