package main

import (
	"AICS_WebBackend/controller"
	"fmt"
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	fmt.Println("Hello world")
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:[]string{"*"},
		AllowMethods:[]string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.GET("/api/aicsweb/v1/test", controller.HelloWorld)

	// user
	e.POST("/api/aicsweb/v1/user/login", controller.LoginUser)
	e.POST("/api/aicsweb/v1/user/register", controller.UserRegister)
	e.PUT("/api/aicsweb/v1/user/info/:username", controller.UserChangePass)

	// File and folder
	e.GET("/api/aicsweb/v1/files/list", controller.SysFileGetList)
	e.POST("/api/aicsweb/v1/files/download", controller.SysFileDownload)

	//e.Logger.Fatal(e.Start("0.0.0.0:9001"))
	e.Logger.Fatal(e.Start("0.0.0.0:9000"))
}
