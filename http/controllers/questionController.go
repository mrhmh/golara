package controllers

import (
	"github.com/labstack/echo/v4"
	"golara/app/database"
	"golara/models"
	"net/http"
	"strconv"
)

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "INDEX PAGE")
}

func IndexQuestion(c echo.Context) error {

	var entities []models.ChallengeQuestion
	database.DB().Limit(5).Find(&entities)

	return c.JSON(http.StatusOK, entities)
}

func ShowQuestion(c echo.Context) error {

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	var entity models.ChallengeQuestion
	database.DB().Where("id = ?", id).First(&entity)

	return c.JSON(http.StatusOK, entity)
}
