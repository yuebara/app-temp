package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// only allow requests from the frontend
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Docker! <3")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	e.GET("/min", func(c echo.Context) error {
		a := 5
		b := 2
		d := 3
		fmt.Println("a:", a, "b:", b, "c:", d)
		return c.JSON(http.StatusOK, struct{ Result int }{Result: IntMin(a, d)})
	})

	e.GET("/max", func(c echo.Context) error {
		a := 5
		b := 2
		d := 3
		fmt.Println("a:", a, "b:", b, "c:", d)
		result := IntMax(a, d)
		return c.JSON(http.StatusOK, struct {
			Result int `json:"result"`
		}{Result: result})
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}

// Simple implementation of an integer minimum
// Adapted from: https://gobyexample.com/testing-and-benchmarking
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func IntMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
