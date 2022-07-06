package service

import (
	"net/http"

	"github.com/awesome-sphere/as-general/assets"
	"github.com/awesome-sphere/as-general/models"

	"github.com/gin-gonic/gin"
)

func GetAllMovies(c *gin.Context) {
	var movies []models.Movie

	if err := models.DB.Find(&movies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	for i, m := range movies {
		m.Poster = assets.GetPoster(m.Poster)
		m.Trailer = assets.GetTrailer(m.Trailer)
		movies[i] = m
	}

	c.JSON(http.StatusOK, gin.H{
		"movies": movies,
	})
}
