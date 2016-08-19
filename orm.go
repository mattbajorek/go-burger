package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Create row struct
type row struct {
	id          int
	burger_name string
	devoured    bool
	date        string
}

func selectAll(db *sql.DB) []row {
	// Query to select all
	rows, err := db.Query("SELECT * FROM burgers")
	checkErr(err)

	// Create slice of rows
	burgers := []row{}

	for rows.Next() {
		var r row
		err = rows.Scan(&r.id, &r.burger_name, &r.devoured, &r.date)
		checkErr(err)
		burgers = append(burgers, r)
	}

	return burgers
}
