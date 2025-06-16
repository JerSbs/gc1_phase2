package main

import (
	"log"
	"net/http"
	"os"

	"gc1_phase2/config"
	"gc1_phase2/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	// ✅ Connect to MySQL database
	config.ConnectDB()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT not found in environment variables! Heroku requires a dynamic PORT.")
	}

	router := routes.SetupRouter()
	log.Println("Server is running on port:", port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
