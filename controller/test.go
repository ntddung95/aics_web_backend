package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// Response : test struct
type Response struct {
	StatusCode int    `json:"StatusCode"`
	Data       string `json:"Data"`
}

// HelloWorld : test function
func HelloWorld(c echo.Context) error {
	response := Response{StatusCode: 200, Data: "Hello, World"}
	return c.JSON(http.StatusOK, response)
}
