package main

import (
	"github.com/awesome-sphere/as-general/assets"
	"github.com/awesome-sphere/as-general/models"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	server := gin.Default()
	server.GET("/authentication", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "auth",
		})
	})
	models.ConnectDatabase()
	assets.ConnectMinio()
	server.Run(":9002")
}
