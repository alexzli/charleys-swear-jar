package handlers

import (
	"charleyswearjar/internal/pkg"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func MakeGetStatsHandler() (func(http.ResponseWriter, *http.Request), error) {
	stats, err := pkg.GetIncrementStats()
	if err != nil {
		return nil, err
	}

	// create daily update timer
	go func() {
		for {
			location, err := time.LoadLocation("America/New_York")
			if err != nil {
				fmt.Printf("couldn't get location in daily update\n")
				continue
			}
			currentTime := time.Now().In(location)
			tomorrowStart := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day()+1, 0, 1, 0, 0, location)

			<-time.After(tomorrowStart.Sub(currentTime))

			stats, err = pkg.GetIncrementStats()
			if err != nil {
				fmt.Printf("couldn't get increment stats in daily update\n")
				continue
			}

			err = pkg.ResetToday()
			if err != nil {
				fmt.Printf("couldn't reset the 'today' stat in daily update\n")
				continue
			}
		}
	}()

	handler := func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data := map[string]string{"pastweek": fmt.Sprintf("%d", stats.PastWeek), "yesterday": fmt.Sprintf("%d", stats.Yesterday)}
		jsonData, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
	}

	return handler, nil
}