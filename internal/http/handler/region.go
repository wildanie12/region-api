package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/wildanie12/region-api/config"
	"github.com/wildanie12/region-api/internal/entity/web"
	"github.com/wildanie12/region-api/internal/service"
)

// RegionHandler concrete implementation object.
type RegionHandler struct {
	regionService *service.RegionService
}

// NewRegion constructs new instance of region handler.
func NewRegion(rs *service.RegionService) *RegionHandler {
	return &RegionHandler{
		regionService: rs,
	}
}

// Provinces handles `GET` `/provinces` endpoint.
func (h *RegionHandler) Provinces(c echo.Context) error {
	// get list of provinces
	data, err := h.regionService.ListProvince(context.Background())
	if err != nil {
		return responseInternalError(c, err)
	}

	// mapping data
	res := []web.Province{}
	err = copier.Copy(&res, data)
	if err != nil {
		return responseInternalError(c, err)
	}

	return c.JSON(http.StatusOK, web.CommonResponse{
		ApiVersion: config.Get().AppVersion,
		Status:     "success",
		IsSuccess:  true,
		Message:    "success getting list of provinces",
		Data:       res,
	})
}

// Regencies handles `GET` `/regencies` endpoint.
func (h *RegionHandler) Regencies(c echo.Context) error {
	// validate parameter
	provinceID, err := strconv.Atoi(c.QueryParam("province_id"))
	if err != nil {
		return responseBadRequest(c, "province_id parameter must be a valid province id number")
	}

	// get list of regencies
	data, err := h.regionService.ListRegency(context.Background(), uint(provinceID))
	if err != nil {
		return responseInternalError(c, err)
	}

	// mapping data
	res := []web.Regency{}
	err = copier.Copy(&res, data)
	if err != nil {
		return responseInternalError(c, err)
	}

	return c.JSON(http.StatusOK, web.CommonResponse{
		ApiVersion: config.Get().AppVersion,
		Status:     "success",
		IsSuccess:  true,
		Message:    "success getting list of regencies",
		Data:       res,
	})
}
