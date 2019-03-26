package routesV1

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthEndpoint(c *gin.Context) {
	statusCode := http.StatusOK
	resp := gin.H{"Message": http.StatusText(statusCode)}
	c.JSON(statusCode, resp)
}

func NotFoundEndPoint( c *gin.Context) {
	statusCode := http.StatusNotFound
	resp := gin.H{"Message": http.StatusText(statusCode)}
	c.JSON(statusCode, resp)
}