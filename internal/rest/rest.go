package rest

import (
	"auth-service/internal/rest/handlers/authHandler"
	"auth-service/internal/rest/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	mv      *middleware.Middleware
	handler authHandler.AuthHandler
	mux     *gin.Engine
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func NewServer(handler authHandler.AuthHandler, mux *gin.Engine, mv *middleware.Middleware) *Server {
	mux.Use(gin.Recovery())
	mux.Use(gin.Logger())
	return &Server{
		mv:      mv,
		handler: handler,
		mux:     mux,
	}
}

func (s *Server) Run() {
	const baseUrl = "api/auth"
	baseGroup := s.mux.Group(baseUrl)
	baseGroup.POST("/login", s.handler.Login)
	baseGroup.POST("/register", s.handler.Register)
	baseGroup.POST("/logout", s.mv.AccessValidation(), s.handler.Logout)
	baseGroup.POST("/refresh", s.mv.RefreshValidation(), s.handler.Refresh)

}
