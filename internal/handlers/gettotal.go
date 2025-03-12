package handlers

import (
	"charleyswearjar/internal/pkg"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetTotalHandler(w http.ResponseWriter, r *http.Request) {
	globals, err := pkg.GetGlobals()
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	data := map[string]string{"total": fmt.Sprintf("%d", globals.Total), "today": fmt.Sprintf("%d", globals.Today)}
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}