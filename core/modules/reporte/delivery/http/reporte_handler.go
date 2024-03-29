package http

import (
	_r "acs/domain/repository"
	"bytes"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type reporteHandler struct {
	reporteUcase _r.ReporteUseCase
}

func NewHandler(e *echo.Echo, reporteUcase _r.ReporteUseCase) {
	handler := reporteHandler{
		reporteUcase: reporteUcase,
	}
	e.POST("/v1/reporte/empleados/", handler.GetReporte)
}



func (h *reporteHandler) GetReporte(c echo.Context) (err error) {
	log.Println("GET REPORTE")
	ctx := c.Request().Context()
	var buffer bytes.Buffer
	var data _r.ReporteRequest
	err = c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, _r.ResponseMessage{Message: err.Error()})
	}
	err = h.reporteUcase.GetReporte(ctx, data, &buffer)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.Blob(http.StatusOK, "reporte.xlsx", buffer.Bytes())
}
