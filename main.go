// Run with go build && go-burger.exe
package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	// Set up router
	router := httprouter.New()
	router.GET("/", Index)
	router.POST("/burgers/add", Add)
	router.POST("/burgers/devour/:id", Devour)

	// Serve static files from the ./public directory
	router.NotFound = http.FileServer(http.Dir("public"))

	log.Fatal(http.ListenAndServe(":8000", router))

}
