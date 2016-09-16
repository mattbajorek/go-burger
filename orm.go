package main

import (
	"log"

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
		log.Fatal(err)
	}
}

func selectAll() []row {

	// Gets a pointer to database connection
	db := dbConnect()
	defer db.Close()

	// Query to select all
	rows, err := db.Query("SELECT * FROM burgers")
	checkErr(err)
	defer rows.Close()

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

func insertOne(burger_name string) {

	// Gets a pointer to database connection
	db := dbConnect()
	defer db.Close()

	// Prepare insert statement
	stmt, err := db.Prepare("INSERT burgers SET burger_name=?")
	checkErr(err)

	// Execute statement
	_, err = stmt.Exec(burger_name)
	checkErr(err)

}
