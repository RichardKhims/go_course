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

	group.POST("/create", func(c *gin.Context) {
		var input database.Course
		c.ShouldBindJSON(&input)
		if err := s.db.CreateCourse(c.Request.Context(), input); err != nil {
			c.Error(err)
		} else {
			c.Status(200)
		}

	})

	group.DELETE("/delete", func(c *gin.Context) {
		var input database.Course
		c.ShouldBindJSON(&input)
		if err := s.db.DeleteCourse(c.Request.Context(), input); err != nil {
			c.Error(err)
		} else {
			c.Status(200)
		}
	})

	group.GET("/convert", func(c *gin.Context) {
		var input ConvertRequestDTO
		c.ShouldBindJSON(&input)
		course, err := s.db.GetCourse(c.Request.Context(), input.CurrencyFrom, input.CurrencyTo)
		if err != nil {
			c.Error(err)
		}
		if len(course) > 0 {
			result := input.Value * course[0].Mean
			c.JSON(200, gin.H{
				"result": result,
			})
		} else {
			c.String(404, "Pair not found")
		}
	})
}
