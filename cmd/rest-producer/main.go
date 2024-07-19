package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"sync"
	"time"

	internal "measurement/internal/rest"
)

const (
	MeasurementCount = 1_000_000
	WorkerCount      = 1000
)

func worker(jobs <-chan int, errs chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	for range jobs {
		measurement := internal.Measurement{
			Latitude:    rand.NormFloat64() + 40.0,
			Longitude:   rand.NormFloat64() - 120.0,
			Altitude:    rand.NormFloat64() + 10.0,
			Temperature: rand.NormFloat64() + 20.0,
			Pressure:    rand.NormFloat64() + 1000.0,
			Humidity:    rand.NormFloat64() + 65.0,
			Timestamp:   time.Now(),
			DeviceID:    "SENSOR001",
		}

		jsonData, err := json.Marshal(measurement)
		if err != nil {
			errs <- err
			continue
		}

		resp, err := http.Post("http://localhost:8080", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			errs <- err
			continue
		}

		var response internal.MeasurementResponse
		err = json.NewDecoder(resp.Body).Decode(&response)
		resp.Body.Close()
		if err != nil {
			errs <- err
			continue
		}

		if !response.Success {
			errs <- fmt.Errorf("server failed to process measurement: %v", response.Message)
			continue
		}
	}
}

func main() {
	start := time.Now()

	jobs := make(chan int, MeasurementCount)
	errs := make(chan error, MeasurementCount)
	var wg sync.WaitGroup

	for w := 1; w <= WorkerCount; w++ {
		wg.Add(1)
		go worker(jobs, errs, &wg)
	}

	for i := 1; i <= MeasurementCount; i++ {
		jobs <- i
	}
	close(jobs)

	wg.Wait()

	log.Printf("Elapsed time: %s", time.Since(start))
	log.Printf("Error count: %d", len(errs))
}
