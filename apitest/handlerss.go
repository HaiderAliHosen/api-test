package apitest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAPI for reuse
func GetAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"message": "boss is the leader",
		})
	}

}
