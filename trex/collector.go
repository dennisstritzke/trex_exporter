package trex

import (
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
)

var (
	metricHashrate      = prometheus.NewDesc("trex_hashrate_total", "hashrate of the worker", []string{"worker"}, nil)
	metricAcceptedCount = prometheus.NewDesc("trex_accepted_count", "accepted share count", []string{"worker"}, nil)
	metricRejectedCount = prometheus.NewDesc("trex_rejected_count", "rejected share count", []string{"worker"}, nil)
	metricSolvedCount   = prometheus.NewDesc("trex_solved_count", "solved block count", []string{"worker"}, nil)
	metricGpuTotal      = prometheus.NewDesc("trex_gpu_total", "number of gpus", []string{"worker"}, nil)
	metricUptime        = prometheus.NewDesc("trex_uptime", "uptime in seconds", []string{"worker"}, nil)

	metricGpuHashrate    = prometheus.NewDesc("trex_gpu_hashrate", "hashrate of the gpu", []string{"worker", "gpu"}, nil)
	metricGpuPower       = prometheus.NewDesc("trex_gpu_power", "gpu power draw", []string{"worker", "gpu"}, nil)
	metricGpuTemperature = prometheus.NewDesc("trex_gpu_temperature", "gpu temperature", []string{"worker", "gpu"}, nil)
	metricGpuFanSpeed    = prometheus.NewDesc("trex_gpu_fan_speed", "gpu fan speed", []string{"worker", "gpu"}, nil)
	metricGpuEfficiency  = prometheus.NewDesc("trex_gpu_efficiency", "gpu efficiency", []string{"worker", "gpu"}, nil)
)

func NewCollector(trexApiAddress string, worker string) *Collector {
	return &Collector{
		trexApiAddress: trexApiAddress,
		worker:         worker,
	}
}

type Collector struct {
	trexApiAddress string
	worker         string
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- metricHashrate
	ch <- metricAcceptedCount
	ch <- metricRejectedCount
	ch <- metricSolvedCount
	ch <- metricGpuTotal
	ch <- metricUptime

	ch <- metricGpuHashrate
	ch <- metricGpuPower
	ch <- metricGpuTemperature
	ch <- metricGpuFanSpeed
	ch <- metricGpuEfficiency
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	summary, err := queryStatus(c.trexApiAddress)
	if err != nil {
		return
	}

	ch <- prometheus.MustNewConstMetric(metricHashrate, prometheus.GaugeValue, float64(summary.Hashrate), c.worker)
	ch <- prometheus.MustNewConstMetric(metricAcceptedCount, prometheus.CounterValue, float64(summary.AcceptedCount), c.worker)
	ch <- prometheus.MustNewConstMetric(metricRejectedCount, prometheus.CounterValue, float64(summary.RejectedCount), c.worker)
	ch <- prometheus.MustNewConstMetric(metricSolvedCount, prometheus.CounterValue, float64(summary.SolvedCount), c.worker)
	ch <- prometheus.MustNewConstMetric(metricGpuTotal, prometheus.GaugeValue, float64(summary.GpuTotal), c.worker)
	ch <- prometheus.MustNewConstMetric(metricUptime, prometheus.GaugeValue, float64(summary.Uptime), c.worker)

	for _, gpuSummary := range summary.Gpus {
		ch <- prometheus.MustNewConstMetric(metricGpuHashrate, prometheus.GaugeValue, float64(gpuSummary.Hashrate), c.worker, strconv.Itoa(gpuSummary.DeviceId))
		ch <- prometheus.MustNewConstMetric(metricGpuPower, prometheus.GaugeValue, float64(gpuSummary.Power), c.worker, strconv.Itoa(gpuSummary.DeviceId))
		ch <- prometheus.MustNewConstMetric(metricGpuTemperature, prometheus.GaugeValue, float64(gpuSummary.Temperature), c.worker, strconv.Itoa(gpuSummary.DeviceId))
		ch <- prometheus.MustNewConstMetric(metricGpuFanSpeed, prometheus.GaugeValue, float64(gpuSummary.FanSpeed), c.worker, strconv.Itoa(gpuSummary.DeviceId))
	}
}
