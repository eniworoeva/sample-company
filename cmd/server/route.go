package server

import (
	"time"

	"github.com/eniworoeva/sample-company/internal/api"
	"github.com/eniworoeva/sample-company/internal/ports"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter is where router endpoints are called
func SetupRouter(handler *api.HTTPHandler, repository ports.Repository) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r := router.Group("/")
	{
		r.GET("/", handler.Readiness)
		r.POST("/computers", handler.CreateComputer)
		r.GET("/computers", handler.GetAllComputers)
		r.GET("/computers/employee/:abbr", handler.GetComputersByEmployee)
		r.GET("/computers/:id", handler.GetComputerByID)
		r.PUT("/computers/:id", handler.UpdateComputer)
		r.DELETE("/computers/:id", handler.DeleteComputer)
		r.PUT("/computers/:id/assign", handler.AssignComputer)
	}

	return router
}
