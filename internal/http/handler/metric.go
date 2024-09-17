package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wildanie12/region-api/config"
	"github.com/wildanie12/region-api/internal/entity/web"
)

// MetricHandler holds all dependency for metric app http handler.
type MetricHandler struct{}

// NewMetricHandler construct app http MetricHandler.
func NewMetricHandler() *MetricHandler {
	return &MetricHandler{}
}

// HealthCheck handles `GET` `/metric/healthcheck`
func (h *MetricHandler) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, web.CommonResponse{
		ApiVersion: config.Get().AppVersion,
		Status:     "success",
		IsSuccess:  true,
		Message:    "health check request is received and is healthy",
		Data:       nil,
	})
}
