package api

import (
	"api/controller"
	"fmt"

	"github.com/labstack/echo/v4"
)

func NewServer(port int) {
	// r := gin.New()
	// r.Use(gin.Recovery())
	// r.GET("/", controller.Hello)
	// r.GET("/users", controller.GetUsers)
	// r.Run(fmt.Sprintf(":%v", port))

	r := echo.New() // New Engine
	r.GET("/", controller.Hello)
	r.GET("/users", controller.GetUsers)
	r.Start(fmt.Sprintf(":%v", port))
}
