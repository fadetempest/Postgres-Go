package server

import (
	"context"
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"restaurant/buisness"
)

type Server struct {
	ctx context.Context
	Address string
	DatabaseUrl string
}

type Handler struct {
	r *mux.Router
	process *buisness.Processing
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

	rp:= &Handler{
		r:    mux.NewRouter(),
		process: buisness.NewProcess(db),
	}

	//store:= NewMethods(rp.Base)

	rp.r.HandleFunc("/add", rp.AddDish)
	rp.r.HandleFunc("/menu", rp.ReadDishes)
	rp.r.HandleFunc("/delete/{id}", rp.DeleteDish)
	rp.r.HandleFunc("/update", rp.UpdateDish)

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

func (h *Handler) AddDish(w http.ResponseWriter, r *http.Request){
	resp, err:= h.process.Add(r)
	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func (h *Handler) DeleteDish(w http.ResponseWriter, r *http.Request){
	resp, err:=h.process.Delete(r)
	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func (h *Handler) UpdateDish(w http.ResponseWriter, r *http.Request){
	resp, err:=h.process.Update(r)
	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func (h *Handler) ReadDishes(w http.ResponseWriter, r *http.Request){
	resp, err:=h.process.ReadAll(r)
	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
