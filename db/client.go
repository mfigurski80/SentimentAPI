package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Start() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/sentiment")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("connected to mysql database")

	results, err := db.Query("SELECT time FROM TimeSeries")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {

	}
}
