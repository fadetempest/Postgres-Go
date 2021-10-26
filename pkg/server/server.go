package server

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"restaurant/tools"
)

type Server struct {
	ctx context.Context
	Address string
}

type Methods struct {

}

func NewServer(ctx context.Context, address string) *Server{
	return &Server{
		ctx: ctx,
		Address: address,
	}
}

func (s *Server) Run() error{
	r:= mux.NewRouter()
	r.HandleFunc("/menu", tools.Menu)
	r.HandleFunc("/add", tools.AddMeal)
	r.HandleFunc("/delete/{id}", tools.DelMeal)
	r.HandleFunc("/update", tools.UpdateMeal)
	srv:=&http.Server{
		Addr: s.Address,
		Handler: r,
	}
	log.Printf("Server is running on %s", s.Address)
	return srv.ListenAndServe()
}

