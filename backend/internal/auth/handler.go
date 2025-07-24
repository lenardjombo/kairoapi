package auth

import (
	// "net/http"

	"github.com/gin-gonic/gin"
	// "github.com/lenardjombo/kairoapi/models"
)

type Handler struct {
	AuthService
}

func NewHandler(h *AuthService) *Handler {
	return &Handler{AuthService: *h}
}

func (h *Handler) RegisterUser(ctx *gin.Context) {
	// var u models.CreateUserReq

	// if err := ctx.ShouldBindJSON(&u); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// res, err := h.AuthService.RegisterUser(ctx.Request.Context(), &u)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// ctx.JSON(http.StatusOK, res)
}

func (h *Handler) LoginUser(ctx *gin.Context){

}