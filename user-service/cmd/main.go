package main

import (
	"log"
	"net/http"
	"social-app/user-service/db"
	"social-app/user-service/handlers"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	database, err := db.ConnectDB()
	if err != nil {
		log.Println(err)
		log.Fatal("DB connection could not be set")
	}
	storage := db.NewStorage(database)
	handler:=handlers.NewHandler(storage)
	mux := http.NewServeMux()
	mux.HandleFunc("POST /register", handler.RegisterUser)
	mux.HandleFunc("POST /login", handler.LoginUser)

	log.Println("User Service running on port 8080:")
	http.ListenAndServe(":8080", mux)
}
