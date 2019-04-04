package middleware

import (
	"github.com/zainkai/forgetful-bartender/pkg/persistence"
	"github.com/gin-gonic/gin"
)

// Attaches context of all open database connections
// to gin and gonic context of request
func DbConnection() gin.HandlerFunc  {
	dbConn := persistence.Init()
	return func(c *gin.Context) {
		c.Keys["dbConn"] = dbConn

		c.Next()
	}
}