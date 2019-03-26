package main

import (
	"github.com/gin-gonic/gin"
  routes "github.com/zainkai/forgetful-bartender/internal/routesV1"
  "fmt"
)

func main () {
	r := gin.Default()
  v1 := r.Group("/api/v1")
  v1.GET("/health", routes.HealthEndpoint)
  
  fmt.Println("Running server on http://localhost:8080")

  r.NoRoute(routes.NotFoundEndPoint)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}