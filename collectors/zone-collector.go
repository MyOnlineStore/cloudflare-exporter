package collectors

import (
	"github.com/prometheus/client_golang/prometheus"
	"log"

	"github.com/cloudflare/cloudflare-go"
)

// Define metrics
var (
	upMetrics = prometheus.NewDesc(
		prometheus.BuildFQName(Namespace, "zones", "up"),
		"Was the Cloudflare Zones scrape successful.",
		nil, nil,
	)

	zoneMetrics = prometheus.NewDesc(
		prometheus.BuildFQName(Namespace, "zones", "current"),
		"Cloudflare zones per plan.",
		[]string{"plan"}, nil,
	)
)

type ZoneCollector struct {
	cfClient *cloudflare.API
}

func NewZoneCollector(cfClient *cloudflare.API) *ZoneCollector {
	return &ZoneCollector {
		cfClient: cfClient,
	}
}

// Collect Cloudflare zone metrics
func (collector *ZoneCollector) Collect(ch chan<- prometheus.Metric) {
	zonesPerPlan := map[string]float64{}

	zones, err := collector.cfClient.ListZones()
	if err != nil {
		ch <- prometheus.MustNewConstMetric(
			upMetrics, prometheus.GaugeValue, 0,
		)
		log.Println(err)
		return
	}

	ch <- prometheus.MustNewConstMetric(
		upMetrics, prometheus.GaugeValue, 1,
	)

	for _, zone := range zones {
		planName := zone.Plan.Name
		planCounter := zonesPerPlan[planName]
		planCounter++
		zonesPerPlan[planName] = planCounter
	}

	for plan, zoneCount := range zonesPerPlan {
		ch <- prometheus.MustNewConstMetric(
			zoneMetrics, prometheus.GaugeValue, zoneCount, plan,
		)
	}

}

// Describe Register metrics
func (collector *ZoneCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- upMetrics
	ch <- zoneMetrics
}