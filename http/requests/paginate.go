package requests

type PaginateRequest struct {
	Page string `json:"page" binding:"required"`
}
