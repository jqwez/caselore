package server

import (
	"log"
	"net/http"

	"github.com/jqwez/caselore/database"
)

type Server struct {
	Mux        *http.ServeMux
	Repository *database.Repository
}

func NewServer(repository *database.Repository) *Server {
	mux := http.NewServeMux()
	RegisterPageRoutes(mux)
	RegisterAPI(mux)
	return &Server{
		Mux:        mux,
		Repository: repository,
	}
}

func (s *Server) Run() {
	err := http.ListenAndServe(":8080", s.Mux)
	if err != nil {
		log.Fatal(err)
	}
}
