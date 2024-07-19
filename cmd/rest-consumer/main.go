package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	internal "measurement/internal/rest"
)

func handleMeasurement(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var measurement internal.Measurement
	err := json.NewDecoder(r.Body).Decode(&measurement)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	response := internal.MeasurementResponse{
		Success: true,
		Message: "Measurement received successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", handleMeasurement)
	fmt.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
