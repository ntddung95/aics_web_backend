package main

import (
	"AICS_WebBackend/controller"
	"fmt"
	"io"
	"os"
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
)
func upload(c echo.Context) error {
	// Read form fields
	name := "XXX"
	email := "YYY"
	absolute_path := c.FormValue("path")
	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		panic(err)
		return err
	}
	src, err := file.Open()
	if err != nil {
		panic(err)
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(absolute_path + "/" + file.Filename)
	if err != nil {
		panic(err)
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		panic(err)
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully with fields name=%s and email=%s.</p>", file.Filename, name, email))
}

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
	e.POST("/api/aicsweb/v1/files/upload", controller.SysFileUpload)
	//e.Logger.Fatal(e.Start("0.0.0.0:9001"))
	e.Logger.Fatal(e.Start("0.0.0.0:9000"))
}
