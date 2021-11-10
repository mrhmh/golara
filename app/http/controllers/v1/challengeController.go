package v1

import (
	"github.com/gin-gonic/gin"
	"golara/app/http/controllers"
	"golara/app/http/requests"
	"golara/app/repositories"
	"golara/app/repositories/contracts"
	"net/http"
)

type challengeController struct {
	controllers.BaseController
	repository contracts.ChallengeRepositoryInterface
}

func (c *challengeController) All(ctx *gin.Context) {

	var req requests.PaginateRequest

	//validate
	if ctx.ShouldBindQuery(&req) != nil {
		c.ValidationError(ctx)
		return
	}

	filters := ctx.Request.URL.Query()
	entities := c.repository.All(filters)

	ctx.JSON(http.StatusOK, entities)
}

func NewChallengeController() *challengeController {
	return &challengeController{
		repository: repositories.NewChallengeRepository(),
	}
}
