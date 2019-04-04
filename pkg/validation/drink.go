package validation

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/zainkai/forgetful-bartender/pkg/persistence"
)

func Drink (c *gin.Context) {
	var drink persistence.Drink
	if err := c.ShouldBindJSON(&drink); err != nil {
		fmt.Println("invalid drink request")
		
		statusCode := http.StatusBadRequest
		resp := gin.H{"Message": http.StatusText(statusCode)}
		c.AbortWithStatusJSON(statusCode, resp)
		return
	}

	fmt.Println(drink)
	c.Keys["BODY"] = drink
}