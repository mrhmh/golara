package requests

type PaginateRequest struct {
	Page string `form:"page" binding:"required"`
}
