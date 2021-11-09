package controllers

import "github.com/gin-gonic/gin"

type BaseController struct{}

func (c *BaseController) validationError(ctx *gin.Context) {
	ctx.JSON(422, gin.H{"message": "Validation error"})
	ctx.Abort()
}
