//github.com/lenardjombo/kairoapi/auth
package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lenardjombo/kairoapi/models"
)

type Handler struct {
	AuthService
}

func NewHandler(s AuthService) *Handler {
	return &Handler{AuthService: s}
}

func (h *Handler) RegisterUser(ctx *gin.Context) {
	var req models.CreateUserReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload", "details": err.Error()})
		return
	}

	res, err := h.AuthService.RegisterUser(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register user", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func (h *Handler) LoginUser(ctx *gin.Context) {
	var req models.LoginUserReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid login payload", "details": err.Error()})
		return
	}

	res, err := h.AuthService.LoginUser(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "login failed", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
