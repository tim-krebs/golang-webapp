// Go connection Sample Code:
package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	_ "github.com/microsoft/go-mssqldb"

	"github.com/tim-krebs/golang-webapp/platform/authenticator"
	"github.com/tim-krebs/golang-webapp/platform/router"
)

var db *sql.DB
var server = "auth-0-db.database.windows.net"
var port = 1433
var user = "timkrebs"
var password = "Y+faZ5KZ3wVRcpdb"
var database = "auth-0-db"

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}
	//// Build connection string
	//connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
	//	server, user, password, port, database)
	//
	//var err error

	//// Create connection pool
	//db, err = sql.Open("sqlserver", connString)
	//if err != nil {
	//	log.Fatal("Error creating connection pool: ", err.Error())
	//}
	//ctx := context.Background()
	//err = db.PingContext(ctx)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//fmt.Printf("Connected!\n")

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
