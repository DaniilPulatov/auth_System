package authHandler

import (
	"auth-service/internal/domain/token"
	"auth-service/internal/domain/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h handler) Register(c *gin.Context) {
	var (
		req user.Input
	)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.authService.Register(c.Request.Context(), &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User successfully registered"})
}

func (h handler) Login(c *gin.Context) {
	var req user.Input

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tokens, err := h.authService.Login(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{
		AccessToken:  tokens["access"],
		RefreshToken: tokens["refresh"],
	})

}

func (h handler) Logout(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id required"})
		return
	}
	if err := h.authService.Logout(c.Request.Context(), userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User successfully logged out"})
}

func (h handler) Refresh(c *gin.Context) {
	var req token.SingleTokenInput
	req.ProvidedToken = c.GetString("refresh_token")
	if req.ProvidedToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "refresh_token required"})
		return
	}
	tokens, err := h.authService.Refresh(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, LoginResponse{
		AccessToken:  tokens["access"],
		RefreshToken: tokens["refresh"],
	})
}
