package api

import (
	"api/controller"
	"fmt"

	"github.com/gin-gonic/gin"
)

func NewServer(port int) {
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/", controller.Hello)
	r.GET("/users", controller.GetUsers)
	r.Run(fmt.Sprintf(":%v", port))
}
