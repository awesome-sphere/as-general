package main

import "github.com/gin-gonic/gin"

func main() {
	gin.SetMode(gin.DebugMode)
	server := gin.Default()
	server.GET("/authentication", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "auth",
		})
	})
	server.Run(":9000")
}
