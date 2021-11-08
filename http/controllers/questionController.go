package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golara/app/auth"
	"golara/app/database"
	"golara/models"
	"net/http"
	"strconv"
)

type QuestionController struct{}

func (q *QuestionController) Index(c *gin.Context) {

	fmt.Println("Controller")
	//var data requests.PaginateRequest
	//
	//if c.BindJSON(&data) != nil {
	//	c.JSON(422, gin.H{"message": "Validation error"})
	//	c.Abort()
	//	return
	//}
	var entities []models.ChallengeQuestion
	database.DB().Limit(5).Find(&entities)

	if auth.Check() {
		fmt.Println(auth.User())
	}

	c.JSON(http.StatusOK, entities)
}

func (q *QuestionController) Show(c *gin.Context) {

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	var entity models.ChallengeQuestion
	database.DB().Where("id = ?", id).First(&entity)

	c.JSON(http.StatusOK, entity)
}
