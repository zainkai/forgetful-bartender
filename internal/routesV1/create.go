package routesV1

import(
	//"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	//"github.com/zainkai/forgetful-bartender/pkg/persistence"
)

func CreateDrinkEndpoint(c *gin.Context)  {
	// dbConn, ok := c.Keys["dbConn"].(*persistence.DB)
	// if !ok {
	// 	fmt.Println("Cannot find database connection")
	// }
	// dbConn.CreateDrink()

	statusCode := http.StatusOK
	resp := gin.H{"Message": http.StatusText(statusCode)}
	c.JSON(statusCode, resp)
}