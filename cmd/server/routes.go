package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kjasuquo/jobslocation/internal/api"
	"time"
)

//SetupRouter is where router endpoints are called
func SetupRouter(handler *api.HTTPHandler) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r := router.Group("/api/v1")
	{
		r.GET("/ping", handler.PingHandler)
		r.GET("/", handler.SearchByLocation)
		r.GET("/title", handler.SearchByTitle)

	}

	return router
}
