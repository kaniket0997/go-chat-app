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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, res)
	return
}

func (h *Handler) LoginUser(ctx *gin.Context) {

	var request LoginUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := h.Service.LoginUser(ctx.Request.Context(), &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.SetCookie("jwt", u.AccessToken, 3600, "/", "localhost", false, true)
	response := &LoginUserResponse{
		Id:       u.Id,
		Username: u.Username,
	}
	ctx.JSON(http.StatusOK, response)
	return
}

func (h *Handler) LogoutUser(ctx *gin.Context) {

	ctx.SetCookie("jwt", "", -1, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{"message": "Logout Successful"})
}
