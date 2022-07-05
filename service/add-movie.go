package service

import (
	"fmt"
	"net/http"

	"github.com/awesome-sphere/as-general/assets"
	"github.com/awesome-sphere/as-general/models"

	"github.com/gin-gonic/gin"
)

func AddMovie(c *gin.Context) {
	var movie models.Movie

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	poster := assets.GetPoster(movie.Poster)
	trailer := assets.GetTrailer(movie.Trailer)

	models.DB.Model(&movie).Updates(models.Movie{
		Trailer: fmt.Sprintf("data:video/mp4;base64,%s", trailer),
		Poster:  fmt.Sprintf("data:image/jpeg;base64,%s", poster),
	})

	if result := models.DB.Create(&movie); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &movie)
}
