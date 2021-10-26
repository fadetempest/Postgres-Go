package base

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"restaurant/meals"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "1488qwerdf"
	dbname = "restaurant"
)

type DishRepo struct {
	Db *sql.DB
}

func NewDishRepo(db *sql.DB, err error) (*DishRepo, error){
	if err != nil{
		return nil, err
	}
	return &DishRepo{Db: db}, nil
}

func AddNewValue(dish *meals.Dish) string{
	sqcon:=fmt.Sprintf("host= %s port= %d user= %s password= %s dbname= %s sslmode=disable", host,port,user,password,dbname)
	db, err:= sql.Open("postgres",sqcon)
	if err != nil{
		log.Fatal("Error while open DB")
	}
	defer db.Close()

	searchQuery:= `SELECT id FROM meals WHERE id=$1`

	if db.QueryRow(searchQuery, dish.ID).Scan(&dish.ID) == sql.ErrNoRows {
		insertValues := `INSERT INTO meals (id, description, composition, price) VALUES ($1, $2, $3, $4)`
		_, er := db.Exec(insertValues, dish.ID, dish.Description, dish.Composition, dish.Price)
		if er != nil {
			return "Error while adding dish"
		}
		return "Successfully added to the menu"
	}
	return fmt.Sprintf("Dish with id=%d already exist", dish.ID)
}

func DeleteValue(id string) string{
	sqcon:=fmt.Sprintf("host= %s port= %d user= %s password= %s dbname= %s sslmode=disable", host,port,user,password,dbname)
	db, err:= sql.Open("postgres",sqcon)
	if err != nil{
		log.Fatal("Error while open DB")
	}
	defer db.Close()

	delValue:=`DELETE FROM meals WHERE id=$1`
	_, er:=db.Exec(delValue, id)
	if er!= nil{
		return "Error while deleting dish"
	}
	return "Successfully deleted"
}

func UpdateValue(dish *meals.Dish) string{
	sqcon:=fmt.Sprintf("host= %s port= %d user= %s password= %s dbname= %s sslmode=disable", host,port,user,password,dbname)
	db, err:= sql.Open("postgres",sqcon)
	if err != nil{
		log.Fatal("Error while open DB")
	}
	defer db.Close()

	updValue:= `UPDATE meals SET description=$1, composition=$2, price=$3 WHERE id=$4`
	_, er:=db.Exec(updValue, dish.Description, dish.Composition, dish.Price, dish.ID)
	if er != nil{
		return "Error while updating the dish"
	}
	return fmt.Sprintf("Successfully updated dish #%d", dish.ID)
}

func GetMenu() ([]meals.Dish, error){
	sqcon:=fmt.Sprintf("host= %s port= %d user= %s password= %s dbname= %s sslmode=disable", host,port,user,password,dbname)
	db, err:= sql.Open("postgres",sqcon)
	if err != nil{
		log.Fatal("Error while open DB")
	}
	defer db.Close()

	rows, er:= db.Query("SELECT * FROM meals ORDER BY id")
	if er!=nil{
		log.Fatal("DB operation error")
	}
	defer rows.Close()

	var dishes []meals.Dish

	for rows.Next(){
		var dish meals.Dish
		if scanEr := rows.Scan(&dish.ID, &dish.Description, &dish.Composition, &dish.Price); scanEr != nil{
			return dishes, scanEr
		}
		dishes=append(dishes, dish)
	}
	if rowErr := rows.Err(); rowErr != nil{
		return dishes,rowErr
	}
	return dishes, nil
}