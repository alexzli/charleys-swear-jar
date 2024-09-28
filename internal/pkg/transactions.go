package pkg

import (
	"encoding/csv"
	"os"
	"sync"
	"time"
)

const Action = "increment"

var transactionsMutex sync.Mutex

func AppendIncrementAction() error {
	transactionsMutex.Lock()
	defer transactionsMutex.Unlock()

	path := os.Getenv("TRANSACTIONS_CSV_FILEPATH")
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

	writer := csv.NewWriter(file)
	currentTime := time.Now()
	formattedTime := currentTime.Format(time.RFC3339)
	record := []string{formattedTime, Action}

	err = writer.Write(record)
	if err != nil {
		return err
	}

	writer.Flush()
	err = writer.Error()
	if err != nil {
		return err
	}

	return nil
}