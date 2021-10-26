package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"golara/routes"
	"log"
	"reflect"
)

var e = echo.New()

type HttpServer struct {
	public echo.Group
}

func httpServerServiceProvider() {

	for _, route := range routes.GetRoutes() {

		if reflect.TypeOf(route).String() != "contracts.Route" {
			log.Fatalf("Type of route is mismatch!")
		}

		switch route.Method {
		case "GET":
			e.GET(route.Path, route.Controller)
		case "POST":
			e.POST(route.Path, route.Controller)
		case "PUT":
			e.PUT(route.Path, route.Controller)
		}
	}

	e.Start(fmt.Sprintf(":%d", Config.App.Port))
}
