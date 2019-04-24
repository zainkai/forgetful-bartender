package routesV1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	constants "github.com/zainkai/forgetful-bartender/pkg"
	"github.com/zainkai/forgetful-bartender/pkg/logger"
	"github.com/zainkai/forgetful-bartender/pkg/persistence"
	"github.com/zainkai/forgetful-bartender/pkg/validation"
)

func GetOneDrinkEndpoint(c *gin.Context) {
	queryID := c.Query("id")
	validation.QueryId(c, queryID)

	if c.IsAborted() {
		return
	}

	dbConn, ok := c.Keys[constants.DB_CONN].(*persistence.DB)
	if !ok {
		logger.Err("routesV1/get", "Cannot find database connection", nil)
		return
	}

	result, err := dbConn.GetOneDrink(queryID)
	if err != nil {
		statusCode := http.StatusNotFound
		resp := gin.H{"Message": http.StatusText(statusCode)}
		c.AbortWithStatusJSON(statusCode, resp)
		return
	}

	statusCode := http.StatusOK
	resp := gin.H{
		"Message": http.StatusText(statusCode),
		"result":  *result}
	c.JSON(statusCode, resp)
}

func GetMultipleDrinkEndpoint(c *gin.Context) {
	pageNum, pageSize := validation.QueryPagination(c)

	if c.IsAborted() {
		return
	}

	dbConn, ok := c.Keys[constants.DB_CONN].(*persistence.DB)
	if !ok {
		logger.Err("routesV1/get", "Cannot find database connection", nil)
		return
	}

	result, err := dbConn.GetMultipleDrinks(pageSize, pageNum)
	if err != nil {
		statusCode := http.StatusInternalServerError
		resp := gin.H{"Message": http.StatusText(statusCode)}
		c.AbortWithStatusJSON(statusCode, resp)
		return
	}

	statusCode := http.StatusOK
	resp := gin.H{
		"Message": http.StatusText(statusCode),
		"result":  result}
	c.JSON(statusCode, resp)
}
