package core

import (
	"database/sql"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_reporteHttp "acs/core/modules/reporte/delivery/http"
	_reporteGenerator "acs/core/modules/reporte/generator"
	_reporteRepo "acs/core/modules/reporte/repository/sql"
	_reporteUsecase "acs/core/modules/reporte/usecase"
	_reporteUtil "acs/core/modules/reporte/util"

	_asistenciaHttp "acs/core/modules/asistencia/delivery/http"
	_asistenciaRepo "acs/core/modules/asistencia/repository/sql"
	_asistenciaUsecase "acs/core/modules/asistencia/usecase"

	_util "acs/core/modules/util"
	_logger "acs/core/modules/util/logger"
	_material "acs/core/modules/util/material"
	_locale "acs/core/modules/util/locale"
)

func InitServer(db *sql.DB) {

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		// AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,echo.HeaderAccessControlAllowCredentials},
	}))
	timeoutContext := time.Duration(20) * time.Second

	logger := _logger.New()
	util := _util.New()
	colorPalette := _material.NewPaletteColor()
	textSizes  := _material.NewTextSize()
	locale := _locale.New()
	//Asistencia
	asistenciaRepo := _asistenciaRepo.NewRepository(db)
	assistenciaUcase := _asistenciaUsecase.NewUseCase(timeoutContext, asistenciaRepo, logger, util)
	_asistenciaHttp.NewHandler(e, assistenciaUcase)

	//Reporte
	reporteUtil := _reporteUtil.New(colorPalette,textSizes,locale)
	reporteGenerator := _reporteGenerator.New(reporteUtil, logger,locale)
	reporteRepo := _reporteRepo.NewRepoReporte(db)
	reporteUcase := _reporteUsecase.NewUseCase(
		timeoutContext,
		reporteRepo,
		reporteGenerator,
		assistenciaUcase,
		logger,
	)

	_reporteHttp.NewHandler(e, reporteUcase)

	e.Start("0.0.0.0:9090")
}
