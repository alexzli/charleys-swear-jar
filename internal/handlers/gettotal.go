package handlers

import (
	"christineswearjar/internal/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetGetTotalHandler(client *services.SpreadsheetClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
        currentCellValue := client.GetTotal()

		w.Header().Set("Content-Type", "application/json")
		data := map[string]string{"value": fmt.Sprintf("%d", currentCellValue)}
		jsonData, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
    }
}