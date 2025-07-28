package metrics

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
)

func InitPrometheus(port string) {
	exporter, err := prometheus.New()
	if err != nil {
		log.Fatalf("failed to initialize prometheus exporter: %v", err)
	}

	provider := metric.NewMeterProvider(metric.WithReader(exporter))
	otel.SetMeterProvider(provider)

	//use in main.go
	// exporter, err := metrics.InitPrometheusMetrics()
	// if err != nil {
	//   log.Fatalf("metrics init error: %v", err)
	// }
	// metrics.ServeMetrics(exporter, "9090")
	http.Handle("/metrics", promhttp.Handler()) // serves metrics
	go func() {
		log.Println("📊 Prometheus metrics available at http://localhost:9090/metrics")
		if err := http.ListenAndServe(":9090", nil); err != nil {
			log.Fatalf("metrics endpoint error: %v", err)
		}
	}()
	// go func() {
	// 	log.Printf("Prometheus metrics available at http://localhost:%s/metrics", port)
	// 	if err := http.ListenAndServe(":"+port, nil); err != nil {
	// 		log.Fatalf("metrics server failed: %v", err)
	// 	}
	// }()
}
