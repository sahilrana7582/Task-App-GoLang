package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sahilrana7582/Task-App-GoLang/internal/database"
	router "github.com/sahilrana7582/Task-App-GoLang/internal/routes"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("❌ Error loading .env file", err)
	}

	db := database.InitDB()
	database.MigrateDB(db)

	r := mux.NewRouter()
	router.RegisterRoutes(r)

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
	port := "8080"

	log.Printf("✅ Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
