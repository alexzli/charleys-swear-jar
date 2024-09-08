package handlers

import (
	"christineswearjar/internal/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetIncrementHandler(client *services.SpreadsheetClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
        newValue := client.IncrementTotal()

		w.Header().Set("Content-Type", "application/json")
		data := map[string]string{"value": fmt.Sprintf("%d", newValue)}
		jsonData, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
    }
}