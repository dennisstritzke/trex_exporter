# T-Rex Exporter ![Test](https://github.com/dennisstritzke/trex_exporter/workflows/Test/badge.svg)

Prometheus exporter for T-Rex NVIDIA GPU miner metrics.


## Using Docker with environment variables

Using Docker you can pass your settings via environment variables
- Address on which to expose metrics: `TREX_EXPORTER_BIND_ADDRESS` (default: `0.0.0.0`) and `TREX_EXPORTER_PORT` (default: `9788`)
- Address of the t-rex API: `TREX_MINER_URL` (default: `http://localhost:4057`)
- Name to identify the T-Rex Miner. The name will be included in every metric as the label 'worker': `TREX_WORKER_NAME` (default: `trex`)
