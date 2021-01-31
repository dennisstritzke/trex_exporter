package trex

type gpuSummary struct {
	DeviceId    int `json:"device_id"`
	Hashrate    int
	Power       int
	Temperature int
	FanSpeed    int `json:"fan_speed"`
}

type summary struct {
	Hashrate      int
	AcceptedCount int `json:"accepted_count"`
	RejectedCount int `json:"rejected_count"`
	SolvedCount   int `json:"solved_count"`
	GpuTotal      int `json:"gpu_total"`
	Uptime        int
	Gpus          []gpuSummary
}
