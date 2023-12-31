package main

import (
	akh "apikeys/internal/app/apikey/handler"
	rth "apikeys/internal/app/relationtype/handler"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	// Initialize Logging
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// Initialize Database
	var err error
	connStr := "host=db port=5432 user=username password=password dbname=app_db sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Database connection is not alive: %s", err)
	}

	// Initialize other resources like cache, message queues, etc. here
	// ...
}

func main() {
	// HTTP routes
	setUpHandlers()

	// Start HTTP server
	log.Println("Starting API keys service on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}

func setUpHandlers() {
	handleApiKeyMethods()
	handleRelationTypeMethods()
}

func handleApiKeyMethods() {
	handler := &akh.APIKeyHandler{DB: db}
	http.HandleFunc("/api/keys/validate", handler.ValidateAPIKeyHandler)
	http.HandleFunc("/api/keys/disable", handler.DisableAPIKeyHandler)
	http.HandleFunc("/api/keys/generate", handler.GenerateAPIKeyHandler)
}

func handleRelationTypeMethods() {
	handler := &rth.RelationTypeHandler{DB: db}
	http.HandleFunc("/api/relation-type/search", handler.GetRelationTypeByName)
}
