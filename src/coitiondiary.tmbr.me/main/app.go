package main

import (
	"log"
	"net/http"

	"coitiondiary.tmbr.me/db"
	"coitiondiary.tmbr.me/routes"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize(host, user, password, dbname string, port int) {
	db.Initialize(host, user, password, dbname, port)
	a.Router = routes.NewRouter()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8080", a.Router))
}
