package main

import (
	"database/sql"
	"github.com/mert-erol/golang_wms/database"
	"github.com/mert-erol/golang_wms/routes"
)

var db *sql.DB

func main() {

	db = database.ConnectDB()

	routes.Application()
}
