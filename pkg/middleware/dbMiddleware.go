package middleware

import (
	"github.com/gin-gonic/gin"
	constants "github.com/zainkai/forgetful-bartender/pkg"
	"github.com/zainkai/forgetful-bartender/pkg/logger"
	"github.com/zainkai/forgetful-bartender/pkg/persistence"
)

const pkgName string = "middleware/db"

// Attaches context of all open database connections
// to gin and gonic context of request
func DbConnection() gin.HandlerFunc {
	dbConn := persistence.Init()
	logger.Log(pkgName, "persistence conn created", nil)
	return func(c *gin.Context) {
		c.Keys[constants.DB_CONN] = dbConn
		c.Next()
	}
}
