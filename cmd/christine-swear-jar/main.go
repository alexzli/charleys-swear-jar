package main

import (
	"christineswearjar/internal/handlers"
	"christineswearjar/internal/services"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load("config/.env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v\n", err)
    }

    spreadsheetId := os.Getenv("SPREADSHEET_ID")
    spreadsheetClient := services.InitializeSpreadsheetClient(spreadsheetId)

    http.HandleFunc("/", handlers.ServeIndexHTML)
    http.HandleFunc("/api/increment", handlers.GetIncrementHandler(spreadsheetClient))
    http.HandleFunc("/api/gettotal", handlers.GetGetTotalHandler(spreadsheetClient))

    addr := "localhost:8080"
    err = http.ListenAndServe(addr, nil)
    if err != nil {
        fmt.Printf("oops server broke\n")
        fmt.Printf("error: %v\n", err)
    }
}