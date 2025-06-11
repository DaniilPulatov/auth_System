package middleware

import (
	"auth-service/internal/domain/token"
	"auth-service/internal/repository/authpin"
	"auth-service/internal/tokenOps"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Middleware struct {
	pin authpin.PinRepository
}

func NewMiddleware(pin authpin.PinRepository) *Middleware {
	return &Middleware{
		pin: pin,
	}
}

func (m *Middleware) AccessValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerVal := c.GetHeader("Authorization")
		if headerVal == "" || !strings.HasPrefix(headerVal, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization header format is invalid"})
		}
		accessToken := strings.TrimPrefix(headerVal, "Bearer ")
		if err := tokenOps.ValidateToken(accessToken); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		}
		claims, err := tokenOps.ExtractClaims(token.SingleTokenInput{ProvidedToken: accessToken})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		}
		val, err := m.pin.Get(c.Request.Context(), claims.UserID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		}
		if val != claims.LogoutPin {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to access this resource, logout pin mismatch"})
		}
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}

func (m *Middleware) RefreshValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input token.SingleTokenInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		}
		if err := tokenOps.ValidateToken(input.ProvidedToken); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		}
		claims, err := tokenOps.ExtractClaims(input)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		}
		val, err := m.pin.Get(c.Request.Context(), claims.UserID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		}
		if val != claims.LogoutPin {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to access this resource, logout pin mismatch"})
		}
		c.Set("refresh_token", input.ProvidedToken)
		c.Next()
	}
}
