package base

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"restaurant/meals"
)

type DishRepo struct {
	Db *sql.DB
}

func NewDishRepo(db *sql.DB) *DishRepo{
	return &DishRepo{Db: db}
}

func AddNewValue(dish *meals.Dish, db *sql.DB) string{
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

func DeleteValue(id string, db *sql.DB) string{
	delValue:=`DELETE FROM meals WHERE id=$1`
	_, er:=db.Exec(delValue, id)
	if er!= nil{
		return "Error while deleting dish"
	}
	return "Successfully deleted"
}

func UpdateValue(dish *meals.Dish, db *sql.DB) string{
	updValue:= `UPDATE meals SET description=$1, composition=$2, price=$3 WHERE id=$4`
	_, er:=db.Exec(updValue, dish.Description, dish.Composition, dish.Price, dish.ID)
	if er != nil{
		return "Error while updating the dish"
	}
	return fmt.Sprintf("Successfully updated dish #%d", dish.ID)
}

func GetMenu(db *sql.DB) ([]meals.Dish, error){
	rows, er := db.Query("SELECT * FROM meals ORDER BY id")
	if er != nil {
		log.Fatal("DB operation error")
	}
	defer rows.Close()

	var dishes []meals.Dish

	for rows.Next() {
		var dish meals.Dish
		if scanEr := rows.Scan(&dish.ID, &dish.Description, &dish.Composition, &dish.Price); scanEr != nil {
			return dishes, scanEr
		}
		dishes = append(dishes, dish)
	}
	if rowErr := rows.Err(); rowErr != nil {
		return dishes, rowErr
	}
	return dishes, nil
}
