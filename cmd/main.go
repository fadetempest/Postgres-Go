package main

import (
	"context"
	_ "github.com/lib/pq"
	"log"
	"restaurant/server"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "1488qwerdf"
	dbname = "restaurant"
)

func main(){
	srv:= server.NewServer(context.Background(), ":8080")

	//sqcon:=fmt.Sprintf("host= %s port= %d user= %s password= %s dbname= %s sslmode=disable", host,port,user,password,dbname)
	//db, err:= base.NewDishRepo(sql.Open("postgres",sqcon))
	//if err != nil{
	//	log.Fatal("Error while open DB")
	//}
	//defer db.Db.Close()

	er:=srv.Run()
	if er != nil{
		log.Println("Error while running the server", er)
	}

}
