// Go connection Sample Code:
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/microsoft/go-mssqldb"

	"github.com/tim-krebs/golang-webapp/platform/authenticator"
	"github.com/tim-krebs/golang-webapp/platform/router"
)

var server = os.Getenv("SERVER")
var port = os.Getenv("PORT")
var user = os.Getenv("USER")
var password = os.Getenv("PASSWORD")
var database = os.Getenv("DATABASE")

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}

	var err error

	// Create connection pool

	fmt.Printf("Connected!\n")

	// Initialize the authenticator
	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	// Initialize the router
	rtr := router.New(auth)

	log.Print("Server listening on http://localhost:3000/")
	if err := http.ListenAndServe("0.0.0.0:3000", rtr); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
