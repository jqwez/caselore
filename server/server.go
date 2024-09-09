package server

import (
	"log"
	"net/http"
)

type Server struct {
	Mux        *http.ServeMux
	Repository interface{}
}

func NewServer(repository interface{}) *Server {
	return &Server{
		Repository: repository,
	}
}

func (s *Server) Run() {
	err := http.ListenAndServe(":8080", s.Mux)
	if err != nil {
		log.Fatal(err)
	}
}
