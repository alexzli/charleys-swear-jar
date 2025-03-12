package handlers

import (
	"charleyswearjar/internal/pkg"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func IncrementTotal() *pkg.Globals {
	globals, err := pkg.GetGlobals()
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	globals.Total += 1
	globals.Today += 1
	err = pkg.WriteGlobals(globals)
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	return globals
}

func AppendIncrementAction() {
	err := pkg.AppendIncrementAction()
	if err != nil {
		log.Fatalf("err: %v", err)
	}
}

func IncrementHandler(w http.ResponseWriter, r *http.Request) {
	newValue := IncrementTotal()
	AppendIncrementAction()

	w.Header().Set("Content-Type", "application/json")
	data := map[string]string{"total": fmt.Sprintf("%d", newValue.Total), "today": fmt.Sprintf("%d", newValue.Today)}
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}