// Run with go build && go-burger.exe
package main

import (
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func handle(path string, handleFunc func(res http.ResponseWriter, req *http.Request) error) {
	http.HandleFunc(path, func(res http.ResponseWriter, req *http.Request) {
		err := handleFunc(res, req)
		if err != nil {
			http.Error(res, err.Error(), 500)
		}
	})
}

func main() {

	// Static files
	http.Handle("/assets/",
		http.StripPrefix("/assets",
			http.FileServer(http.Dir("./assets"))))

	// Template
	handle("/", func(res http.ResponseWriter, req *http.Request) error {

		// Selects all burgers from db
		burgers := selectAll()

		// Prints out burgers
		fmt.Println(burgers)

		// Render to template
		tpl, err := template.ParseFiles("assets/templates/index.gohtml")
		if err != nil {
			return err
		}
		err = tpl.Execute(res, burgers)
		if err != nil {
			return err
		}

		return nil
	})

	// Post
	handle("/burgers/add", func(res http.ResponseWriter, req *http.Request) error {

		// Parse url parameter for burger name
		burger_name := req.FormValue("burger_name")

		// Insert burger name into db
		insertOne(burger_name)

		// Redirect to main page
		http.Redirect(res, req, "/", 301)

		return nil
	})

	http.ListenAndServe(":8000", nil)

}
