package frf

import (
	"net/http"

	"github.com/Haider/AliHosen/example/apitest"
	"github.com/gin-gonic/gin"
)

// type servPostRequest struct {
// 	Title string `json:"title"`
// 	Post  string `json:"post"`
// }

// ServPost for reuse
func ServPost(serv apitest.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		// requestBody := servPostRequest{}
		var req apitest.Item
		err := c.Bind(&req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
			// panic(err)
			return
		}
		// item := apitest.Item{
		// 	Title: requestBody.Title,
		// 	Post:  requestBody.Post,
		// }
		// fmt.Println(req)

		serv.Add(req)

		// c.Status(http.StatusNoContent)
		c.JSON(http.StatusOK, gin.H{"Error": ""})
	}

}
