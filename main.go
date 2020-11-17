package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"cloudflare-exporter/collectors"

	"github.com/cloudflare/cloudflare-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	addr        = flag.String("web.listen-address", ":9178", "The address to listen on for HTTP requests.")
	metricsPath = flag.String("web.metrics-path", "/metrics", "Path under which to expose metrics.")
	healthPath  = flag.String("web.health-path", "/health", "Path under which to expose exporter health.")
)

var (
	zoneCollector collectors.ZoneCollector
)

func initCollectors() {
	cfClient, err := cloudflare.NewWithAPIToken(os.Getenv("CF_API_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	zoneCollector = *collectors.NewZoneCollector(cfClient)
}

func main() {
	flag.Parse()

	initCollectors()

	// Register Prometheus Collectors
	prometheus.MustRegister(prometheus.NewBuildInfoCollector())
	prometheus.MustRegister(&zoneCollector)

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
