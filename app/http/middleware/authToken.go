package middleware

import (
	"github.com/gin-gonic/gin"
	"golara/app/models"
	"golara/core/facades/auth"
	"net/http"
	"strings"
)

func AuthToken(c *gin.Context) {

	token := c.GetHeader("Authorization")
	if token == "" {
		unauthorizedError(c)
		return
	}

	splitToken := strings.Split(token, "Bearer ")
	if len(splitToken) != 2 {
		unauthorizedError(c)
		return
	}
	token = splitToken[1]
	if token == "" {
		unauthorizedError(c)
		return
	}

	var user models.User
	err := user.GetUserByToken(token)
	if err != nil {
		unauthorizedError(c)
		return
	}

	auth.LoggedInUser(user)

	c.Next()
}

func unauthorizedError(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
}
