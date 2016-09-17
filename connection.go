package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func dbConnect() *sql.DB {
	db, err := sql.Open("mysql", "ilmd1mjdiyyftpuy:z5wjohny3dimm257@tcp(sp6xl8zoyvbumaa2.cbetxkdyhwsb.us-east-1.rds.amazonaws.com:3306)/tq3ji0sayjqfjwb7")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	return db
}
