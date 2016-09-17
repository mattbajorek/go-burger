package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Template
func Index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	// Selects all burgers from db
	burgers := selectAll()

	// Prints out burgers
	fmt.Println(burgers)

	// Render to template
	tpl, err := template.ParseFiles("public/assets/templates/index.gohtml")
	if err != nil {
		fmt.Println(err)
	}
	err = tpl.Execute(res, burgers)
	if err != nil {
		fmt.Println(err)
	}

}

// Post
func Add(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	// Parse form data for burger name
	burger_name := req.FormValue("burger_name")

	// Insert burger name into db
	insertOne(burger_name)

	// Redirect to main page
	http.Redirect(res, req, "/", 301)

}

// Put
func Devour(res http.ResponseWriter, req *http.Request, params httprouter.Params) {

	// Parse form data for value
	devoured := req.FormValue("devoured")

	fmt.Println(devoured)

	// Grab id from params
	id := params.ByName("id")

	fmt.Println(id)

	// Update burger with id to devoured
	updateOne(devoured, id)

	// Redirect to main page
	http.Redirect(res, req, "/", 301)

}
