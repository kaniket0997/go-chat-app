package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	Service
}

func NewUserHandler(svc Service) *Handler {
	return &Handler{
		Service: svc,
	}
}

func (h *Handler) CreateUser(ctx *gin.Context) {

	var request CreateUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.Service.CreateUser(ctx.Request.Context(), &request)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, res)
	return
}
