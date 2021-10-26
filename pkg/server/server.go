package server

import (
	"context"
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"restaurant/base"
	"restaurant/tools"
)

type Server struct {
	ctx context.Context
	Address string
	DatabaseUrl string
}

type Handl struct {
	r *mux.Router
	Base *base.DishRepo
}

func NewServer(ctx context.Context, address string, database string) *Server{
	return &Server{
		ctx: ctx,
		Address: address,
		DatabaseUrl: database,
	}
}

func (s *Server) Run() error{
	db, err:= openDb(s.DatabaseUrl)
	if err != nil{
		log.Println("Error while opening DataBase")
	}
	defer db.Close()

	rp:= &Handl{
		r: mux.NewRouter(),
		Base: base.NewDishRepo(db),
	}

	store:=tools.NewMethods(rp.Base)

	rp.r.HandleFunc("/add", store.AddMeal)
	rp.r.HandleFunc("/menu", store.Menu)
	rp.r.HandleFunc("/delete/{id}", store.DelMeal)
	rp.r.HandleFunc("/update", store.UpdateMeal)

	srv:=&http.Server{
		Addr: s.Address,
		Handler: rp.r,
	}

	log.Printf("Server is running on %s", s.Address)
	return srv.ListenAndServe()
}

func openDb(baseUrl string) (*sql.DB, error){
	db, err:=sql.Open("postgres", baseUrl)
	if err!=nil{
		return nil, err
	}
	if err:=db.Ping();err!=nil{
		return nil,err
	}
	return db,nil
}

