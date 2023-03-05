package routes

import (
	"net/http"

	"github.com/atramors/go-cli/controllers"
	"github.com/atramors/go-cli/models"
	"github.com/gin-gonic/gin"
)

type Getter interface {
	GetInfo(arg string) (mod []models.WeatherInfo)
}

func GetWeather(db Getter, input string) []models.WeatherInfo {
	return db.GetInfo(input)
}

// StartGin function
func StartGin(db Getter) {
	router := gin.Default()
	api := router.Group("/api/v1")
	{
		// api.GET("/cities", GetAllCityWeather)
		api.GET("/cities/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.JSON(200, gin.H{
				"weather info": controllers.GetData(db, name),
			})
		})
	}
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":8080")
}
