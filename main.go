// Run with go build && go-burger.exe
package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Gets a pointer to database connection
	db := dbConnect()
	defer db.Close()

	// Selects all burgers from db
	burgers := selectAll(db)

	// Prints out burgers
	fmt.Println(burgers)
}
