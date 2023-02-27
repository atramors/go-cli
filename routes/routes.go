package routes

// import (
// 	"net/http"

// 	"github.com/atramors/go-cli/controllers"
// 	"github.com/gin-gonic/gin"
// )

// // StartGin function
// func StartGin() {
// 	router := gin.Default()
// 	api := router.Group("/api/v1")
// 	{
// 		// api.GET("/cities", GetAllCityWeather)
// 		api.GET("/cities/:name", controllers.GetCityWeather)
// 	}
// 	router.NoRoute(func(c *gin.Context) {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	})
// 	router.Run(":8080")
// }
