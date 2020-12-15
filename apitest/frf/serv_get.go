package frf

import (
	"net/http"

	"github.com/Haider/AliHosen/example/apitest"
	"github.com/gin-gonic/gin"
)

// ServAllGet -- reuse
func ServAllGet(serv apitest.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := serv.GetAll()
		c.JSON(http.StatusOK, results)
	}

}
