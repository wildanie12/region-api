package http

import "github.com/wildanie12/region-api/internal/http/handler"

// Provide is providing assigned routes and handlers for http server to serve.
func (ah *AppHTTP) Provide(
	mh *handler.MetricHandler,
	rh *handler.RegionHandler,
) {
	// metrics
	mg := ah.e.Group("/metrics")
	mg.GET("/healthcheck", mh.HealthCheck)

	// region
	ah.e.GET("/provinces", rh.Provinces)
	ah.e.GET("/regencies", rh.Regencies)
	ah.e.GET("/districts", rh.Districts)
}
