package routesV1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	constants "github.com/zainkai/forgetful-bartender/pkg"
	"github.com/zainkai/forgetful-bartender/pkg/logger"
	"github.com/zainkai/forgetful-bartender/pkg/persistence"
	"github.com/zainkai/forgetful-bartender/pkg/validation"
)

const pkgName string = "routesV1/create"

func CreateDrinkEndpoint(c *gin.Context) {
	validation.Drink(c)
	if c.IsAborted() {
		return
	}

	dbConn, ok := c.Keys[constants.DB_CONN].(*persistence.DB)
	reqBody, _ := c.Keys[constants.REQ_BODY].(persistence.Drink)
	if !ok {
		logger.Err(pkgName, "Cannot find database connection", nil)
		return
	}
	result, _ := dbConn.CreateDrink(reqBody)

	// format result for user
	fmt.Println(result)

	statusCode := http.StatusOK
	resp := gin.H{
		"Message":    http.StatusText(statusCode),
		"InsertedID": result.InsertedID}
	c.JSON(statusCode, resp)
}
