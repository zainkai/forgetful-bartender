package main

import (
	"github.com/gin-gonic/gin"
  routes "github.com/zainkai/forgetful-bartender/internal/routesV1"
  "github.com/zainkai/forgetful-bartender/configs"
  "fmt"
)

func main () {
  config.LoadConfig()
  configuration := *config.Config

	r := gin.Default()
  v1 := r.Group("/api/v1")
  v1.GET("/health", routes.HealthEndpoint)
  
  r.NoRoute(routes.NotFoundEndPoint)

  fmt.Println("Running server on http://localhost:" + configuration.Port)
	r.Run(":" + configuration.Port) // listen and serve on 0.0.0.0:8080
}