package httpServer

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"golara/app/config"
	"golara/routes"
	"log"
	"reflect"
)

func InitHttpServer() {
	e := echo.New()

	for _, route := range routes.Routes {

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

	e.Start(fmt.Sprintf(":%d", config.Config().App.Port))
}
