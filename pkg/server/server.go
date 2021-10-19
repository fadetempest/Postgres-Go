package server

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	ctx context.Context
	Address string
}

func NewServer(ctx context.Context, address string) *Server{
	return &Server{
		ctx: ctx,
		Address: address,
	}
}

func (s *Server) Run() error{
	r:= mux.NewRouter()
	r.HandleFunc("/", menu)
	srv:=&http.Server{
		Addr: s.Address,
		Handler: r,
	}
	log.Printf("Server is running on %s", s.Address)
	return srv.ListenAndServe()
}

func menu(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello world"))
}
