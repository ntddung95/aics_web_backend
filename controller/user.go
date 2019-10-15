package controller

import (
	"AICS_WebBackend/services"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// UserLogin UserLogin
type UserLogin struct {
	User       string `json:"user"`
	Permission string `json:"permission"`
}

// ResponseUserLogin ResponseUserLogin
type ResponseUserLogin struct {
	StatusCode int       `json:"StatusCode"`
	Data       UserLogin `json:"Data"`
}

// RequestUserLogin RequestUserLogin
type RequestUserLogin struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

// LoginUser : controller login user
func LoginUser(c echo.Context) error {
	m := echo.Map{}
	var response Response
	if err := c.Bind(&m); err != nil {
		response = Response{StatusCode: 9999, Data: "Body invalid"}
		return c.JSON(http.StatusBadRequest, response)
	}

	if m["user"] == nil || m["pass"] == nil {
		response = Response{StatusCode: 9999, Data: "Body invalid"}
		return c.JSON(http.StatusBadRequest, response)
	}
	user := fmt.Sprintf("%v", m["user"])
	pass := fmt.Sprintf("%v", m["pass"])
	success, permission := services.LoginUser(user, pass)
	if !success {
		response = Response{StatusCode: 401, Data: "Not authorized"}
		return c.JSON(http.StatusBadRequest, response)
	}
	loginInfo := UserLogin{User: user, Permission: permission}
	responseInfo := ResponseUserLogin{StatusCode: 200, Data: loginInfo}
	return c.JSON(http.StatusOK, responseInfo)
}
