package service

import (
	"net/http"

	"github.com/awesome-sphere/as-general/models"

	"github.com/gin-gonic/gin"
)

func GetMovie(c *gin.Context) {
	var movie models.Movie
	id := c.Param("id")

	if err := models.DB.Find(&movie).Where("id", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"movie": movie,
	})
}
