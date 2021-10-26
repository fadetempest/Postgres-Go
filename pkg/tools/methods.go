package tools

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"restaurant/base"
	"restaurant/meals"
)

type Methods struct {
	Store *base.DishRepo
}

func NewMethods(store *base.DishRepo) *Methods{
	return &Methods{Store: store}
}

func (m *Methods) AddMeal(w http.ResponseWriter, r *http.Request){
	data, readErr:=ioutil.ReadAll(r.Body)
	if readErr != nil{
		w.WriteHeader(http.StatusBadRequest)
	}
	var dish meals.Dish
	err:=json.Unmarshal(data,&dish)
	if err!=nil{
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Error while reading request"))
	}
	coded, jerr:= json.Marshal(m.Store.AddNewValue(&dish))
	if jerr!=nil{
		w.WriteHeader(http.StatusConflict)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(coded)
}

func (m *Methods) DelMeal(w http.ResponseWriter, r *http.Request){
	coded, jerr:= json.Marshal(m.Store.DeleteValue(r.URL.Path[8:]))
	if jerr!=nil{
		w.WriteHeader(http.StatusConflict)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(coded)
}

func (m *Methods) UpdateMeal(w http.ResponseWriter, r *http.Request){
	data, readErr:=ioutil.ReadAll(r.Body)
	if readErr != nil{
		w.WriteHeader(http.StatusBadRequest)
	}
	var dish meals.Dish
	err:=json.Unmarshal(data,&dish)
	if err!=nil{
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Error while reading request"))
	}
	w.Header().Set("Content-Type", "application/json")
	coded, jerr:= json.Marshal(m.Store.UpdateValue(&dish))
	if jerr!=nil{
		w.WriteHeader(http.StatusConflict)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(coded)
}

func (m *Methods) Menu(w http.ResponseWriter, r *http.Request){
	allMenu, err:=m.Store.GetMenu()
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
	}
	coded, jerr:= json.Marshal(allMenu)
	if jerr!=nil{
		w.WriteHeader(http.StatusConflict)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(coded)
}
