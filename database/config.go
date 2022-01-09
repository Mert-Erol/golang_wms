package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectDB() *sql.DB {

	db, err := sql.Open("mysql", "root:@/wms")

	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected!")

	return db
}
