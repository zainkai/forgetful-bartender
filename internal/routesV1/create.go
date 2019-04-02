package routesV1

import(
	"github.com/gin-gonic/gin"
	"net/http"
	// "github.com/zainkai/forgetful-bartender/pkg/persistence"
)

func CreateDrinkEndpoint(c *gin.Context)  {
	// dbConn := persistence.Init()
  	// dbConn.CreateDrink()

	statusCode := http.StatusOK
	resp := gin.H{"Message": http.StatusText(statusCode)}
	c.JSON(statusCode, resp)
}