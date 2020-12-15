package main

import (
	"github.com/Haider/AliHosen/example/apitest"
	"github.com/Haider/AliHosen/example/apitest/frf"
	"github.com/gin-gonic/gin"
)

func main() {

	serv := apitest.New()
	r := gin.Default()
	r.GET("/leader", apitest.GetAPI())
	r.GET("/serv", frf.ServAllGet(serv))
	r.POST("/serv", frf.ServPost(serv))

	//fileupload.SetupRoutes()
	r.Run()

}
