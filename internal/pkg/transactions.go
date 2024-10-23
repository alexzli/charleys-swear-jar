package pkg

import (
	"encoding/csv"
	"os"
	"time"
)

const Action = "increment"

func AppendIncrementAction() error {
	path := os.Getenv("TRANSACTIONS_CSV_FILEPATH")
	
	lock := GetLock(path)
	lock.Lock()
	defer lock.Unlock()

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