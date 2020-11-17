package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	addr        = flag.String("web.listen-address", ":9178", "The address to listen on for HTTP requests.")
	metricsPath = flag.String("web.metrics-path", "/metrics", "Path under which to expose metrics.")
	healthPath  = flag.String("web.health-path", "/health", "Path under which to expose exporter health.")
)

func main() {
	flag.Parse()

	// Register Prometheus Collectors
	prometheus.MustRegister(prometheus.NewBuildInfoCollector())

	// Expose the registered metrics via HTTP.
	http.Handle(*metricsPath, promhttp.Handler())

	http.HandleFunc(*healthPath, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`Healthy`))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>Cloudflare Exporter</title></head>
             <body>
             <h1>CloudFlare Exporter</h1>
             <p><a href='` + *metricsPath + `'>Metrics</a></p>
             <p><a href='` + *healthPath + `'>Health</a></p>
             </body>
             </html>`))
	})

	// Start webserver
	log.Printf("Cloudflare exporter is listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
