package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Create row struct
type row struct {
	Id         int
	BurgerName string
	Devoured   bool
	Date       string
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func selectAll(db *sql.DB) []row {
	// Query to select all
	rows, err := db.Query("SELECT * FROM burgers")
	checkErr(err)

	// Create slice of rows
	burgers := []row{}

	for rows.Next() {
		var r row
		err = rows.Scan(&r.Id, &r.BurgerName, &r.Devoured, &r.Date)
		checkErr(err)
		burgers = append(burgers, r)
	}

	return burgers
}
