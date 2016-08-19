package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := sql.Open("mysql", "root:@/burger_db")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Query to select all
	rows, err := db.Query("SELECT * FROM burgers")
	checkErr(err)

	// Create row struct
	type row struct {
		id          int
		burger_name string
		devoured    bool
		date        string
	}

	// Create slice of rows
	burgers := []row{}

	for rows.Next() {
		var r row
		err = rows.Scan(&r.id, &r.burger_name, &r.devoured, &r.date)
		checkErr(err)
		burgers = append(burgers, r)
	}

	fmt.Println(burgers)
}
