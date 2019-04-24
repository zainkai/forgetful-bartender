package main

import (
	"github.com/gin-gonic/gin"
	config "github.com/zainkai/forgetful-bartender/configs"
	routes "github.com/zainkai/forgetful-bartender/internal/routesV1"
	"github.com/zainkai/forgetful-bartender/pkg/logger"
	"github.com/zainkai/forgetful-bartender/pkg/middleware"
)

func initGin(c *gin.Context) {
	if c.Keys == nil {
		c.Keys = make(map[string]interface{})
	}
	c.Next()
}

func main() {
	config.LoadConfig()
	if config.Config == nil {
		logger.Err("Main", "Config was not loaded", nil)
	}
	configuration := *config.Config

	r := gin.Default()
	r.Use(initGin)
	r.Use(middleware.DbConnection())

	v1 := r.Group("/api/v1")
	v1.GET("/health", routes.HealthEndpoint)
	v1.POST("/drink", routes.CreateDrinkEndpoint)
	v1.GET("/drink", routes.GetOneDrinkEndpoint)
	v1.GET("/drink/batch", routes.GetMultipleDrinkEndpoint)

	r.NoRoute(routes.NotFoundEndPoint)

	logger.Log("Main", "Running server on http://localhost:"+configuration.Port, nil)
	r.Run(":" + configuration.Port) // listen and serve on 0.0.0.0:8080
}
