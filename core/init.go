package core

import (
	"database/sql"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_reporteHttp "acs/core/modules/reporte/delivery/http"
	_reporteRepo "acs/core/modules/reporte/repository/sql"
	_reporteUsecase "acs/core/modules/reporte/usecase"
)

func InitServer(db *sql.DB) {

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		// AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,echo.HeaderAccessControlAllowCredentials},
	}))
	timeoutContext := time.Duration(20) * time.Second

	reporteRepo := _reporteRepo.NewRepoReporte(db)
	reporteUcase := _reporteUsecase.NewUseCase(timeoutContext,reporteRepo)
	_reporteHttp.NewHandler(e,reporteUcase)


	e.Start("0.0.0.0:9090")
}