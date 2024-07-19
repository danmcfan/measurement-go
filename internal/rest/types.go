package rest

import "time"

type Measurement struct {
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Altitude    float64   `json:"altitude"`
	Temperature float64   `json:"temperature"`
	Pressure    float64   `json:"pressure"`
	Humidity    float64   `json:"humidity"`
	Timestamp   time.Time `json:"timestamp"`
	DeviceID    string    `json:"device_id"`
}

type MeasurementResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
