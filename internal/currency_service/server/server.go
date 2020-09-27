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

	currencyGroup := group.Group("/currency")
	currencyGroup.POST("/create", func(c *gin.Context) {
		var input database.Currency
		c.ShouldBindJSON(&input)
		if err := s.db.CreateCurrency(c.Request.Context(), input.Symbol, input.Name); err != nil {
			c.Error(err)
		}
		c.Status(200)
	})

	currencyGroup.DELETE("/:id", func(c *gin.Context) {
		if err := s.db.DeleteCurrency(c.Request.Context(), c.Param("id")); err != nil {
			c.Error(err)
		}
		c.Status(200)
	})

	courseGroup := group.Group("course")
	courseGroup.GET("/", func(c *gin.Context) {
		var input database.Course
		c.ShouldBindJSON(&input)
		course, err := s.db.GetCourse(c.Request.Context(), input.Currency1, input.Currency2)
		if err != nil {
			c.Error(err)
		}
		c.JSON(200, course)
	})
}
