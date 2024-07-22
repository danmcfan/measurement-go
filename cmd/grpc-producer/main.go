package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	internal "measurement/internal/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	MeasurementCount = 1_000_000
	WorkerCount      = 1000
)

func worker(client internal.MeasurementServiceClient, jobs <-chan int, errs chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	for range jobs {
		measurement := &internal.Measurement{
			Latitude:    rand.NormFloat64() + 40.0,
			Longitude:   rand.NormFloat64() - 120.0,
			Altitude:    rand.NormFloat64() + 10.0,
			Temperature: rand.NormFloat64() + 20.0,
			Pressure:    rand.NormFloat64() + 1000.0,
			Humidity:    rand.NormFloat64() + 65.0,
			Timestamp:   time.Now().Unix(),
			DeviceId:    "SENSOR001",
		}

		resp, err := client.SendMeasurement(context.Background(), measurement)
		if err != nil {
			errs <- err
			continue
		}

		if !resp.Success {
			errs <- fmt.Errorf("server failed to process measurement: %v", resp.Message)
			continue
		}
	}
}

func main() {
	start := time.Now()

	conn, err := grpc.NewClient(
		"dns:///consumer-headless.default.svc.cluster.local:50051",
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect: %s", err)
	}
	defer conn.Close()

	client := internal.NewMeasurementServiceClient(conn)

	jobs := make(chan int, MeasurementCount)
	errs := make(chan error, MeasurementCount)
	var wg sync.WaitGroup

	for w := 1; w <= WorkerCount; w++ {
		wg.Add(1)
		go worker(client, jobs, errs, &wg)
	}

	for i := 1; i <= MeasurementCount; i++ {
		jobs <- 1
	}
	close(jobs)

	wg.Wait()

	log.Printf("Elapsed time: %s", time.Since(start))
	log.Printf("Error count: %d", len(errs))
}
