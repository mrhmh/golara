package contracts

import (
	"github.com/gin-gonic/gin"
)

type Route struct {
	Method     string
	Path       string
	Controller func(c *gin.Context)
}
