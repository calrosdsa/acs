package http

import (
	_r "acs/domain/repository"
	"net/http"
	"strconv"

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

func (h *AsistenciaHandler) GetAsistenciasUser(c echo.Context) (err error) {
	ctx := c.Request().Context()
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 1
	}
	size := 20
	chGuid := c.Param("chGuid")
	res, nextPage, count, err := h.asistenciaUCase.GetAsistenciasUser(ctx, chGuid, page, size)
	if err != nil {
		return c.JSON(http.StatusOK, _r.ResponseMessage{Message: err.Error()})
	}
	response := struct {
		NextPage  int             `json:"nextPage"`
		Results   []_r.Asistencia `json:"results"`
		PageCount int             `json:"pageCount"`
		PageSize  int             `json:"pageSize"`
	}{
		NextPage:  nextPage,
		Results:   res,
		PageCount: (count / size) + 1,
		PageSize:  size,
	}
	return c.JSON(http.StatusOK, response)
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
	// err = h.asistenciaUCase.UpdateAsistenciaFromIncomingData(ctx, data)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, _r.ResponseMessage{Message: err.Error()})
	// }
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
