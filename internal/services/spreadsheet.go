package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type SpreadsheetClient struct {
    srv *sheets.Service
    spreadsheetId string
    mu sync.Mutex
}

func InitializeSpreadsheetClient(spreadsheetId string) *SpreadsheetClient {
    ctx := context.Background()

    // Read the service account key JSON file
    data, err := os.ReadFile("config/credentials.json")
    if err != nil {
        log.Fatalf("Unable to read client secret file: %v", err)
    }

	// Parse the service account JSON credentials
    config, err := google.JWTConfigFromJSON(data, sheets.SpreadsheetsScope)
    if err != nil {
        log.Fatalf("Unable to parse client secret file to config: %v", err)
    }

	// Create a new Sheets service client
    srv, err := sheets.NewService(ctx, option.WithHTTPClient(config.Client(ctx)))
    if err != nil {
        log.Fatalf("Unable to retrieve Sheets client: %v", err)
    }

    return &SpreadsheetClient{
        srv: srv,
        spreadsheetId: spreadsheetId,
    }
}

func (client *SpreadsheetClient) GetGlobals() map[string]string {
    client.mu.Lock()
    defer client.mu.Unlock()

    out := make(map[string]string)
    
    // column A stores all the keys, column B stores all the values
    readRange := "Globals!A:B"

	// Retrieve the values from the specified range
    resp, err := client.srv.Spreadsheets.Values.Get(client.spreadsheetId, readRange).Do()
    if err != nil {
        log.Fatalf("Unable to retrieve data from sheet: %v", err)
    }

    if len(resp.Values) == 0 {
        fmt.Println("No data found.")
    } else {
        for _, row := range resp.Values {
            key := row[0].(string)
            value := row[1].(string)
            out[key] = value
        }
    }

    return out
}

// Increments the total swear jar value
func (client *SpreadsheetClient) IncrementTotal() int64 {
    client.mu.Lock()
    defer client.mu.Unlock()

	rangeToUpdate := os.Getenv("SPREADSHEET_GLOBALS_NAME") + "!" + os.Getenv("GLOBALS_TOTAL_CELL_ID")

    // get old value
    resp, err := client.srv.Spreadsheets.Values.Get(client.spreadsheetId, rangeToUpdate).Do()
    if err != nil {
        log.Fatalf("Unable to get spreadsheet's 'total' cell value. error: %v\n", err)
    }
    currentCellValue := resp.Values[0][0].(string)
    currentCellValueInt, err := strconv.ParseInt(currentCellValue, 10, 64)
    if err != nil {
        log.Fatalf("Unable to convert cell's value to int64. error: %v\n", err)
    }

    // set up new value to update with
    newCellValueInt := currentCellValueInt + 1
	vr := &sheets.ValueRange{
        MajorDimension: "ROWS",
        Values: [][]interface{}{
            {newCellValueInt},
        },
    }

	// write data to the specified cell
    _, err = client.srv.Spreadsheets.Values.Update(client.spreadsheetId, rangeToUpdate, vr).ValueInputOption("RAW").Do()
    if err != nil {
        log.Fatalf("Unable to update cell: %v\n", err)
    }

    return newCellValueInt
}

// Gets the current swear jar value
func (client *SpreadsheetClient) GetTotal() int64 {
    client.mu.Lock()
    defer client.mu.Unlock()
    rangeToUpdate := os.Getenv("SPREADSHEET_GLOBALS_NAME") + "!" + os.Getenv("GLOBALS_TOTAL_CELL_ID")

    resp, err := client.srv.Spreadsheets.Values.Get(client.spreadsheetId, rangeToUpdate).Do()
    if err != nil {
        log.Fatalf("Unable to get spreadsheet's 'total' cell value. error: %v\n", err)
    }
    currentCellValue := resp.Values[0][0].(string)
    currentCellValueInt, err := strconv.ParseInt(currentCellValue, 10, 64)
    if err != nil {
        log.Fatalf("Unable to convert cell's value to int64. error: %v\n", err)
    }

    return currentCellValueInt
}