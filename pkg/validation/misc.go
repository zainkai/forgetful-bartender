package validation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zainkai/forgetful-bartender/pkg/logger"
)

func QueryId(c *gin.Context, id string) {
	if id == "" {
		logger.Log(pkgName, "invalid drink request", id)

		statusCode := http.StatusBadRequest
		resp := gin.H{"Message": http.StatusText(statusCode)}
		c.AbortWithStatusJSON(statusCode, resp)
		return
	}
}
