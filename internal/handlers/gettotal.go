package handlers

import (
	"christineswearjar/internal/pkg"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetTotal() int {
	globals, err := pkg.GetGlobals()
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	return globals.Total
}

func GetTotalHandler(w http.ResponseWriter, r *http.Request) {
	currentCellValue := GetTotal()

	w.Header().Set("Content-Type", "application/json")
	data := map[string]string{"value": fmt.Sprintf("%d", currentCellValue)}
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}