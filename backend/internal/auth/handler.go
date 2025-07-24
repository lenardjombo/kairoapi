package auth

import (
	"github.com/gin-gonic/gin"
	// "net/http"

	// "github.com/lenardjombo/kairoapi/models"
)

type Handler struct {
	AuthService
}

func NewHandler(h *AuthService) *Handler {
	return &Handler{AuthService: *h}
}

func (h *Handler) RegisterUser(ctx *gin.Context) {

}
