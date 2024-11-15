package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))
	r.GET("/ping", SendCurrentTime)
	return r
}

func SendCurrentTime(c *gin.Context) {
	respBody := struct {
		CurrentTime string `json:"datetime"`
	}{
		CurrentTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	c.IndentedJSON(http.StatusOK, respBody)
}
