package main

// SubSpace Browser Automation Backend
// Entry point for the Go-Rod automation server.
// Handles API requests, browser lifecycle, and database initialization.
//
// Usage: go run cmd/server/main.go

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/sudhan/browser-automation/internal/api"
	"github.com/sudhan/browser-automation/internal/automation"
	"github.com/sudhan/browser-automation/internal/models"
	"github.com/sudhan/browser-automation/internal/store"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	cfg := models.LoadConfig()

	if cfg.Email == "" || cfg.Password == "" {
		log.Fatal("Please set LINKEDIN_EMAIL and LINKEDIN_PASSWORD environment variables")
	}

	// Initialize Store
	db := store.NewStore("automation.db")

	// Initialize Browser Manager
	bm := automation.NewBrowserManager(cfg.Headless)

	// Initialize API Handler
	handler := api.NewHandler(bm, db)

	// Setup Routes
	http.HandleFunc("/api/start", handler.StartAutomation)
	http.HandleFunc("/api/stop", handler.StopAutomation)
	http.HandleFunc("/api/data", handler.GetData)
	http.HandleFunc("/api/status", handler.GetStatus)

	// Start Server
	go func() {
		fmt.Println("Server starting on :8080...")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	}()

	// Keep main blocked properly or handle signals?
	// For this prototype, we'll just block on a channel or select{}
	select {}
}
