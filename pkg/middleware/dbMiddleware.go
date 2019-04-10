package middleware

import (
	"github.com/zainkai/forgetful-bartender/pkg/persistence"
	"github.com/zainkai/forgetful-bartender/pkg/logger"
	"github.com/gin-gonic/gin"
)

const pkgName string = "middleware/db"

// Attaches context of all open database connections
// to gin and gonic context of request
func DbConnection() gin.HandlerFunc  {
	dbConn := persistence.Init()
	logger.Log(pkgName, "persistence conn created", nil)
	return func(c *gin.Context) {
		c.Keys["dbConn"] = dbConn

		c.Next()
	}
}