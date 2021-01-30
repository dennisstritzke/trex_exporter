package trex

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func queryStatus(address string) (*Summary, error) {
	response, err := http.Get(fmt.Sprintf("%s/summary", address))
	if err != nil {
		return nil, err
	}

	var summary Summary

	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&summary)
	if err != nil {
		return nil, err
	}

	return &summary, nil
}

type GpuSummary struct {
	DeviceId    int `json:"device_id"`
	Hashrate    int
	Name        string
	Vendor      string
	Power       int
	Temperature int
	FanSpeed    int `json:"fan_speed"`
	Efficiency  string
}

type Summary struct {
	Hashrate      int
	AcceptedCount int `json:"accepted_count"`
	RejectedCount int `json:"rejected_count"`
	SolvedCount   int `json:"solved_count"`
	GpuTotal      int `json:"gpu_total"`
	Uptime        int
	Gpus          []*GpuSummary
}
