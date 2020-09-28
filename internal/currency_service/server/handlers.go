package server

import (
	"github.com/RichardKhims/go_course/internal/currency_service/database"
	"github.com/gin-gonic/gin"
)

func (s *Server) CreateCourse (c *gin.Context) {
	var input database.Course
	c.ShouldBindJSON(&input)
	if err := s.db.CreateCourse(c.Request.Context(), input); err != nil {
		c.Error(err)
	} else {
		c.Status(200)
	}
}

func (s *Server) DeleteCourse (c *gin.Context) {
	var input database.Course
	c.ShouldBindJSON(&input)
	if err := s.db.DeleteCourse(c.Request.Context(), input); err != nil {
		c.Error(err)
	} else {
		c.Status(200)
	}
}

func (s *Server) ConvertCourse (c *gin.Context) {
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
}