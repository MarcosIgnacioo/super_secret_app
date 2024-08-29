package apis

import (
	"github.com/gin-gonic/gin"
)

func GenericHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"esto": "functiona",
	})
}
