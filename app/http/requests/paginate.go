package requests

type PaginateRequest struct {
	Page    int `form:"page" binding:"required"`
	PerPage int `form:"per_page" binding:"required"`
}
