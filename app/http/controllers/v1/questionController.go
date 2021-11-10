package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golara/app/http/controllers"
	"golara/app/http/requests"
	"golara/app/models"
	"golara/core/facades"
	"golara/core/facades/auth"
	"net/http"
	"strconv"
)

type QuestionController struct {
	controllers.BaseController
}

func (c *QuestionController) Index(ctx *gin.Context) {

	var req requests.PaginateRequest

	//validate
	if ctx.ShouldBindQuery(&req) != nil {
		c.ValidationError(ctx)
		return
	}

	var entities []models.ChallengeQuestion
	facades.DB().Limit(5).Find(&entities)

	if auth.Check() {
		fmt.Println(auth.User())
	}

	ctx.JSON(http.StatusOK, entities)
}

func (c *QuestionController) Show(ctx *gin.Context) {

	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	var entity models.ChallengeQuestion
	facades.DB().Where("id = ?", id).First(&entity)

	ctx.JSON(http.StatusOK, entity)
}
