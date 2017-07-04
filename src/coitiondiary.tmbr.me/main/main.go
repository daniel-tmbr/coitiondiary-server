package main

import (
	"coitiondiary.tmbr.me/db"
)

const (
	host     = "localhost"
	port     = 1234
	user     = "danieltmbr"
	password = ""
	dbname   = "coitiondiary"
)

var A App

func main() {
	A := App{}
	A.Initialize(host, user, password, dbname, port)
	A.Run(":8080")
	defer db.Close()
}
