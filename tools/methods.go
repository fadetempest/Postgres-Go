package tools

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"restaurant/base"
	"restaurant/meals"
)

func AddMeal(w http.ResponseWriter, r *http.Request){
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
	coded, jerr:= json.Marshal(base.AddNewValue(&dish))
	if jerr!=nil{
		w.WriteHeader(http.StatusConflict)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(coded)
}

func DelMeal(w http.ResponseWriter, r *http.Request){
	coded, jerr:= json.Marshal(base.DeleteValue(r.URL.Path[8:]))
	if jerr!=nil{
		w.WriteHeader(http.StatusConflict)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(coded)
}

func UpdateMeal(w http.ResponseWriter, r *http.Request){
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
	coded, jerr:= json.Marshal(base.UpdateValue(&dish))
	if jerr!=nil{
		w.WriteHeader(http.StatusConflict)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(coded)
}


func Menu(w http.ResponseWriter, r *http.Request){
	allMenu, err:=base.GetMenu()
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