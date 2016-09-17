// Run with go build && go-burger.exe
package main

import (
	"log"
	"net/http"
	"os"

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

	// Get port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Fatal(http.ListenAndServe(":"+port, router))

}
