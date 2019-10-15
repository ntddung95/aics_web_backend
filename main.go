package main

import (
	"AICS_WebBackend/controller"
	"fmt"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	fmt.Println("Hello world")
	e := echo.New()
	e.GET("/api/aicsweb/v1/test", controller.HelloWorld)

	// user
	e.POST("/api/aicsweb/v1/user/login", controller.LoginUser)

	e.Logger.Fatal(e.Start("127.0.0.1:9000"))
}
