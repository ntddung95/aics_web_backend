package controller

import (
	"AICS_WebBackend/services"
	"net/http"
	"fmt"

	"github.com/labstack/echo"
	logger "github.com/sirupsen/logrus"
)


func SysFileGetList(c echo.Context) error {
	logger.Debug("Sys file get list controller")
	result, err := services.SysFileGetList()
	if err != nil {
		response := ResponseUser{StatusCode: 9999, Data: err.Error()}
		return c.JSON(http.StatusBadRequest, response)
	}
	responseInfo := ResponseUser{StatusCode: 200, Data:result}
	return c.JSON(http.StatusOK, responseInfo)
}

func SysFileDownload(c echo.Context) error {
	m := echo.Map{}
        if err := c.Bind(&m); err != nil {
                response := ResponseUser{StatusCode: 9999, Data: "Body invalid"}
                return c.JSON(http.StatusBadRequest, response)
        }
	if m["file_path"] == nil {
                response := ResponseUser{StatusCode: 9999, Data: "Body invalid"}
                return c.JSON(http.StatusBadRequest, response)
        }
	file_path := fmt.Sprintf("%v", m["file_path"])
	err := services.SysCheckFile(file_path)
	if err != nil{
		response := ResponseUser{StatusCode: 9997, Data: err.Error()}
                return c.JSON(http.StatusBadRequest, response)
	}
	return c.File(file_path)

}

func SysFileUpload(c echo.Context) error {
	absolute_path := c.FormValue("path")
	err := services.SysCheckUploadPath(absolute_path)
	if err != nil {
		response := ResponseUser{StatusCode: 9997, Data: err.Error()}
                return c.JSON(http.StatusBadRequest, response)
	}
	file, err := c.FormFile("file")
	if err != nil {
		response := ResponseUser{StatusCode: 9997, Data: err.Error()}
                return c.JSON(http.StatusBadRequest, response)
	}
	err = services.SysFileUpload(absolute_path, file)
	if err != nil {
		response := ResponseUser{StatusCode: 9997, Data: err.Error()}
                return c.JSON(http.StatusBadRequest, response)
	}
	response := ResponseUser{StatusCode: 9999, Data: "OK"}
	return c.JSON(http.StatusOK, response)

}
