package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)


func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// OS
	r.GET("/os", func(c *gin.Context) {
		c.String(http.StatusOK, "os")
	})
	

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
