package controller

import (
	"AICS_WebBackend/services"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	logger "github.com/sirupsen/logrus"
)

// UserLogin UserLogin
type UserLogin struct {
	User       string `json:"user"`
	Permission string `json:"permission"`
}

// ResponseUserLogin ResponseUserLogin
type ResponseUser struct {
	StatusCode int       `json:"StatusCode"`
	Data       interface{} `json:"Data"`
}

// LoginUser : controller login user
func LoginUser(c echo.Context) error {
	m := echo.Map{}
	var response ResponseUser
	if err := c.Bind(&m); err != nil {
		response = ResponseUser{StatusCode: 9999, Data: "Body invalid"}
		return c.JSON(http.StatusBadRequest, response)
	}

	if m["user"] == nil || m["pass"] == nil {
		response = ResponseUser{StatusCode: 9999, Data: "Body invalid"}
		return c.JSON(http.StatusBadRequest, response)
	}
	user := fmt.Sprintf("%v", m["user"])
	pass := fmt.Sprintf("%v", m["pass"])
	success, permission := services.LoginUser(user, pass)
	if !success {
		response = ResponseUser{StatusCode: 401, Data: "Not authorized"}
		return c.JSON(http.StatusBadRequest, response)
	}
	loginInfo := UserLogin{User: user, Permission: permission}
	responseInfo := ResponseUser{StatusCode: 200, Data: loginInfo}
	return c.JSON(http.StatusOK, responseInfo)
}

// UserRegister :
func UserRegister(c echo.Context) error {
        m := echo.Map{}
        var response ResponseUser
        if err := c.Bind(&m); err != nil {
                response = ResponseUser{StatusCode: 9999, Data: "Body invalid"}
                return c.JSON(http.StatusBadRequest, response)
        }

        if m["user"] == nil || m["pass"] == nil || m["permission"] == nil{
                response = ResponseUser{StatusCode: 9999, Data: "Body invalid"}
                return c.JSON(http.StatusBadRequest, response)
        }
        user := fmt.Sprintf("%v", m["user"])
        pass := fmt.Sprintf("%v", m["pass"])
	permission:= fmt.Sprintf("%v", m["permission"])
	success, err := services.UserRegister(user, pass, permission)
	if err != nil{
		response = ResponseUser{StatusCode:9998, Data:err.Error()}
		return c.JSON(http.StatusBadRequest, response)
	}
	if !success {
		response = ResponseUser{StatusCode:9998, Data:"Not sucess"}
		return c.JSON(http.StatusBadRequest, response)

	}
	responseInfo := ResponseUser{StatusCode: 200, Data: "OK"}
        return c.JSON(http.StatusOK, responseInfo)
}

func UserChangePass(c echo.Context) error {
	user := c.Param("username")
	logger.Debug("User: ", user)
	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseUser{StatusCode:9999, Data: err.Error()})
	}
	if m["oldpass"] == nil || m["newpass"] == nil {
		response := ResponseUser{StatusCode: 9999, Data: "Body invalid"}
		return c.JSON(http.StatusBadRequest, response)
	}
	oldpass := fmt.Sprintf("%v", m["oldpass"])
	newpass := fmt.Sprintf("%v", m["newpass"])

	err := services.UserChangePass(user, oldpass, newpass)

	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseUser{StatusCode:9999, Data: err.Error()})
	}
	return c.JSON(http.StatusOK, ResponseUser{StatusCode: 200, Data: "OK"})
}
