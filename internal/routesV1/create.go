package routesV1

import(
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/zainkai/forgetful-bartender/pkg/validation"
	"github.com/zainkai/forgetful-bartender/pkg/persistence"
)

func CreateDrinkEndpoint(c *gin.Context)  {
	validation.Drink(c)
	if c.IsAborted() {
		return
	}

	dbConn, ok := c.Keys["dbConn"].(*persistence.DB)
	reqBody, _ := c.Keys["BODY"].(persistence.Drink)
	if !ok {
		fmt.Println("Cannot find database connection")
	}
	result, _ := dbConn.CreateDrink(reqBody)

	// format result for user
	fmt.Println(result)

	statusCode := http.StatusOK
	resp := gin.H{
		"Message": http.StatusText(statusCode),
		"InsertedID": result.InsertedID}
	c.JSON(statusCode, resp)
}