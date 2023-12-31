package http

import (
	_reporte "acs/domain/repository/reporte"
	"bytes"
	"net/http"

	"github.com/labstack/echo/v4"
)

type reporteHandler struct{
	reporteUcase  _reporte.ReporteUseCase
}

func NewHandler(e *echo.Echo,reporteUcase _reporte.ReporteUseCase){
	handler := reporteHandler{
		reporteUcase: reporteUcase,
	}
	e.GET("/v1/reporte/empleados/",handler.GetReporteEmpleado)
}

func (h *reporteHandler)GetReporteEmpleado(c echo.Context)(err error){
	ctx := c.Request().Context()
	var buffer bytes.Buffer
	err = h.reporteUcase.GetReporteEmpleado(ctx,&buffer)
	if err != nil {
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.Blob(http.StatusOK, "reporte.xlsx", buffer.Bytes())
}