package buisness

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"restaurant/meals"
	"restaurant/repository"
)

type Processing struct {
	repo *repository.DishRepo
}

func NewProcess(db *sql.DB) *Processing{
	return &Processing{repo: repository.NewDishRepo(db)}
}

func (p *Processing) Add(r *http.Request) ([]byte,error){
	data, readErr:=ioutil.ReadAll(r.Body)
	if readErr != nil{
		return nil, readErr
	}
	var dish meals.Dish
	err:=json.Unmarshal(data,&dish)
	if err!=nil{
		return nil,err
	}
	coded, jerr:= json.Marshal(p.repo.AddNewValue(&dish))
	if jerr!=nil{
		return nil,jerr
	}
	return coded,nil
}

func (p *Processing) Delete(r *http.Request) ([]byte, error){
	coded, jerr:= json.Marshal(p.repo.DeleteValue(r.URL.Path[8:]))
	if jerr!=nil{
		return nil, jerr
	}
	return coded, nil
}

func (p *Processing) Update(r *http.Request) ([]byte, error){
	data, readErr:=ioutil.ReadAll(r.Body)
	if readErr != nil{
		return nil,readErr
	}
	var dish meals.Dish
	err:=json.Unmarshal(data,&dish)
	if err!=nil{
		return nil,err
	}
	coded, jerr:= json.Marshal(p.repo.UpdateValue(&dish))
	if jerr!=nil{
		return nil, jerr
	}
	return coded,nil
}

func (p *Processing) ReadAll(r *http.Request) ([]byte, error){
	allMenu, err:=p.repo.GetMenu()
	if err != nil{
		return nil, err
	}
	coded, jerr:= json.Marshal(allMenu)
	if jerr!=nil{
		return nil, jerr
	}
	return coded, nil
}