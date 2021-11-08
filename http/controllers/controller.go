package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.String(http.StatusOK, "The available groups are [...]")
}
