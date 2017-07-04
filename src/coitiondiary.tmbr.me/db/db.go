package db

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

func Initialize(host, user, password, dbname string, port int) {

	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)

	var err error
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("connected to db")

}

func Close() {
	DB.Close()
}
