package http

import (
	_r "acs/domain/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AsistenciaHandler struct {
	asistenciaUCase _r.AsistenciaUseCase
}

func NewHandler(e *echo.Echo, asustenciaUCase _r.AsistenciaUseCase) {
	handler := AsistenciaHandler{
		asistenciaUCase: asustenciaUCase,
	}
	e.POST("/v1/asistencia/update/", handler.UpdateAsistenciaFromIncomingData)
	e.POST("/v1/asistencia/recover/", handler.RecoverAsistenciaAllUsers)
}

func (h *AsistenciaHandler) UpdateAsistenciaFromIncomingData(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var data _r.TMarcacionAsistencia
	err = c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, _r.ResponseMessage{Message: err.Error()})
	}
	err = h.asistenciaUCase.InsertMarcacion(ctx, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, _r.ResponseMessage{Message: err.Error()})
	}
	err = h.asistenciaUCase.UpdateAsistenciaFromIncomingData(ctx, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, _r.ResponseMessage{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, nil)
}

func (h *AsistenciaHandler) RecoverAsistenciaAllUsers(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var data struct {
		Fecha string `json:"fecha"`
	}
	err = c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, _r.ResponseMessage{Message: err.Error()})
	}
	// err = h.asistenciaUCase.InsertMarcacion(ctx, data)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, _r.ResponseMessage{Message: err.Error()})
	// }
	err = h.asistenciaUCase.RevocerAsistenciaAllUsers(ctx, data.Fecha)
	if err != nil {
		return c.JSON(http.StatusBadRequest, _r.ResponseMessage{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, _r.ResponseMessage{Message: "Successfully"})
}
