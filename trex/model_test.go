package trex

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"os"
	"testing"
)

func TestParsesTrexSummary_Api3_3(t *testing.T) {
	apiResponse, err := os.Open("../trex-api-responses/summary-api_3-3.json")
	if err != nil {
		t.Error(err)
		return
	}
	defer apiResponse.Close()

	var parsedSummary summary

	err = json.NewDecoder(apiResponse).Decode(&parsedSummary)
	if err != nil {
		t.Error(err)
		return
	}

	expectedSummary := summary{
		Hashrate:      91189394,
		AcceptedCount: 2,
		RejectedCount: 3,
		Uptime:        76,
		GpuTotal:      1,
		SolvedCount:   4,
		Gpus: []gpuSummary{
			{Hashrate: 91189394, FanSpeed: 54, Temperature: 65, Power: 237, DeviceId: 0},
		},
	}

	if !cmp.Equal(parsedSummary, expectedSummary) {
		t.Error("Summary api not correctly parsed")
		t.Logf("Parsed   %+v", parsedSummary)
		t.Logf("Expected %+v", expectedSummary)
	}
}
