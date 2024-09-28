package handlers

import (
	"christineswearjar/internal/pkg"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func IncrementTotal() int {
	globals, err := pkg.GetGlobals()
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	globals.Total += 1
	out := globals.Total
	err = pkg.WriteGlobals(globals)
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	return out
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
	data := map[string]string{"value": fmt.Sprintf("%d", newValue)}
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}