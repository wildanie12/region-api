package http

import "github.com/wildanie12/region-api/internal/http/handler"

// Provide is providing assigned routes and handlers for http server to serve.
func (ah *AppHTTP) Provide(
	mh *handler.MetricHandler,
) {
	// metrics
	mg := ah.e.Group("/metric")
	mg.GET("/healthcheck", mh.HealthCheck)
}
