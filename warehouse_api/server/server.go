package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type handler struct{}

func main() {

	e := echo.New()

	e.POST("/login",login)

	api := e.Group("/api")
	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	
	api.Use(middleware.JWTWithConfig(config))

	api.GET("/:exchange/profile",getProfile)
	api.GET("/:exchange/financials",getFinancials)

	e.Logger.Fatal(e.Start(":8000"))
}

