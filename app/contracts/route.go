package contracts

import "github.com/labstack/echo/v4"

type Route struct {
	Method     string
	Path       string
	Controller func(c echo.Context) error
}
