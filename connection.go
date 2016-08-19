package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func dbConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:@/burger_db")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	return db
}
