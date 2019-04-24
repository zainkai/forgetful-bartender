package validation

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	constants "github.com/zainkai/forgetful-bartender/pkg"
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

func QueryPagination(c *gin.Context) (int64, int64) {
	pageNum, err1 := strconv.Atoi(c.Query("page"))
	if err1 != nil {
		pageNum = 1
	}

	pageSize, err2 := strconv.Atoi(c.Query("size"))
	if err2 != nil {
		pageSize = int(constants.MAX_PAGE_SIZE)
	}

	return int64(pageNum), int64(pageSize)
}
