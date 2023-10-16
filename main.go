package main

import (
	"encoding/json"
	"fmt"
	"github.com/cruzzan/stc-intel/StcStatistics"
	"net/http"
	"os"
	"time"
)

func main() {
	stats := StcStatistics.NewService(&http.Client{})

	timestamp := time.Now()
	noClassesData, err := stats.NoClasses()
	if err != nil {
		panic(err)
	}
	err = writeToJsonFile(fmt.Sprintf("%d_no_classes.json", timestamp.Unix()), noClassesData)
	if err != nil {
		panic(err)
	}

	classCount, err := stats.ClassCount()
	if err != nil {
		panic(err)
	}
	err = writeToJsonFile(fmt.Sprintf("%d_class_count.json", timestamp.Unix()), classCount)
	if err != nil {
		panic(err)
	}

	classBookingPercentage, err := stats.FullyBookedPercentage()
	if err != nil {
		panic(err)
	}
	err = writeToJsonFile(fmt.Sprintf("%d_class_booking_percentage.json", timestamp.Unix()), classBookingPercentage)
	if err != nil {
		panic(err)
	}

	os.Exit(0)
}

func writeToJsonFile(fileName string, data any) error {
	a, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshall the result to json: %w", err)
	}

	err = os.WriteFile(fileName, a, 0644)
	if err != nil {
		return fmt.Errorf("failed to write result to file %s: %w", fileName, err)
	}

	return nil
}
