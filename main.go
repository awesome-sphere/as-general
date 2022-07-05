package main

import (
	"github.com/awesome-sphere/as-general/assets"
	"github.com/awesome-sphere/as-general/models"
	"github.com/awesome-sphere/as-general/service"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	server := gin.Default()

	models.ConnectDatabase()
	assets.ConnectMinio()

	server.GET("/get-all-movies", service.GetAllMovies)

	server.POST("/add-movie", service.AddMovie)

	server.Run(":9002")
}
