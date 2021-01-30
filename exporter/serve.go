package exporter

import (
	"github.com/dennisstritzke/trex_exporter/trex"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"net/http"
)

var TrexApiAddress string
var Worker string
var WebListenAddress string

func Serve() {
	var err error

	collector := trex.NewCollector(TrexApiAddress, Worker)
	prometheus.MustRegister(collector)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`<html>
             <head><title>T-Rex NVIDIA GPU miner Exporter</title></head>
             <body>
             <h1>T-Rex NVIDIA GPU miner Exporter</h1>
             <p><a href='/metrics'>Metrics</a></p>
             </body>
             </html>`))
	})
	http.Handle("/metrics", promhttp.Handler())

	log.Infoln("Listening on", WebListenAddress)
	err = http.ListenAndServe(WebListenAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
}
