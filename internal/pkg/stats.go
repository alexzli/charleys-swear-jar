package pkg

import (
	"encoding/csv"
	"log"
	"os"
	"time"
)

type Stats struct {
	// all date stats are continuous (e.g. if current time was 8/21/2024 11:45am, the past week is any date >= 8/14/2024 11:45am)
	PastWeek int
	Yesterday int
}

const IncrementAction = "increment"

// Gathers all statistics from past transactions
func GetIncrementStats() (*Stats, error) {
	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		return nil, err
	}
	currentTime := time.Now().In(location)

	pastWeekStart := currentTime.AddDate(0, 0, -7)
	yesterdayStart := currentTime.AddDate(0, 0 , -1)
	pastWeekCounter := 0
	yesterdayCounter := 0

	// read transactions
	path := os.Getenv("TRANSACTIONS_CSV_FILEPATH")
	
	lock := GetLock(path)
	lock.Lock()
	defer lock.Unlock()

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
    if err != nil {
        return nil, err
    }
    defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading CSV:", err)
	}

	for _, record := range records {
		dateString := record[0]
		actionString := record[1]

		if actionString != IncrementAction {
			continue
		}

		date, err := time.Parse(time.RFC3339, dateString)
		if err != nil {
			return nil, err
		}

		if date.After(pastWeekStart) {
			pastWeekCounter++
		}

		if date.After(yesterdayStart) {
			yesterdayCounter++
		}
	}

	out := Stats{
		PastWeek: pastWeekCounter,
		Yesterday: yesterdayCounter,
	}
	return &out, nil
}