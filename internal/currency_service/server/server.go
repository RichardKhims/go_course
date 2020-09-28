package server

import (
	"fmt"
	"github.com/RichardKhims/go_course/internal/currency_service/config"
	"github.com/RichardKhims/go_course/internal/currency_service/database"
	"github.com/gin-gonic/gin"
)

// Server struct
type Server struct {
	router *gin.Engine
	port   int
	db     database.Database
}

type ConvertRequestDTO struct {
	CurrencyFrom string
	CurrencyTo string
	Value float64
}

// New create new Server instance
func New(cfg config.ServerConfig, db database.Database) *Server {
	return &Server{
		router: gin.Default(),
		port:   cfg.Port,
		db:     db,
	}
}

// Run server
func (s *Server) Run() {
	s.initHandlers()
	s.router.Run(fmt.Sprintf(":%d", s.port))
}

func (s *Server) initHandlers() {
	group := s.router.Group("/api")

	group.POST("/create", s.CreateCourse)
	group.DELETE("/delete", s.DeleteCourse)
	group.GET("/convert", s.ConvertCourse)
}
