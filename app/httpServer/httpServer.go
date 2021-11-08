package httpServer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golara/app/config"
	"golara/http/middleware"
	"golara/routes"
	"reflect"
)

func InitHttpServer() {
	r := gin.Default()

	r.Use(middleware.AuthToken)

	v1 := r.Group("/ar/api/v1")
	for _, route := range routes.GetRoutes() {

		if reflect.TypeOf(route).String() != "contracts.Route" {
			panic("Type of route is mismatch!")
		}

		switch route.Method {
		case "GET":
			v1.GET(route.Path, route.Controller)
		case "POST":
			v1.POST(route.Path, route.Controller)
		case "PUT":
			v1.PUT(route.Path, route.Controller)
		}
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found!!"})
	})

	r.Run(fmt.Sprintf(":%d", config.Config().App.Port))
}
