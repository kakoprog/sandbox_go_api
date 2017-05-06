package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/kakoprog/sandbox_go_api/config"
	"strconv"
)

func main() {
	port := config.Read("server", "port").(int)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(port)))
}
