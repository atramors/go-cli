package main

import (
	"net/http"

	"github.com/atramors/go-cli/controllers"
	"github.com/atramors/go-cli/database"
	"github.com/gin-gonic/gin"
	// "github.com/atramors/go-cli/routes"
)

func main() {
	postgresClient := database.ConnectPostgreSQL()
	DBConn := database.ConnectToDB(postgresClient.DB)
	defer DBConn.CloseConnection()

	router := gin.Default()
	api := router.Group("/api/v1")
	{
		// api.GET("/cities", GetAllCityWeather)
		api.GET("/cities/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.JSON(200, gin.H{
				"weather info": controllers.GetData(DBConn, name),
			})
		})
	}
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":8080")

	// controllers.GetData(db_conn)

	// routes.StartGin()
}
