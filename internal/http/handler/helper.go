package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wildanie12/region-api/config"
	"github.com/wildanie12/region-api/internal/entity/web"
)

// responseInternalError can help handler to return response shape for any internal error.
func responseInternalError(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, web.CommonResponse{
		ApiVersion: config.Get().AppVersion,
		Status:     "error",
		IsSuccess:  false,
		Message:    "internal server error: " + err.Error(),
	})
}
