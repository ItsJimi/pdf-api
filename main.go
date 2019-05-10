package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	e := echo.New()
	e.GET("/", home)
	e.Logger.Fatal(e.Start(":" + port))
}
