package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/anna", func(c *gin.Context) {
		c.String(http.StatusOK, "EU TE AMOO!!!!!!!!!!")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and server on 0.0.0.0:8080
}
