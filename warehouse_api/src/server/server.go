package main

import (
	"os"
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"ava.fund/alpha/Post-Covid/warehouse_api/src/handler"
	"ava.fund/alpha/Post-Covid/warehouse_api/src/utils"
)

func main() {

	if len(os.Args) != 2 {
        fmt.Printf("Usage: %s <config>\n", os.Args[0])
        os.Exit(0)
    }
    utils.LoadConfig(os.Args[1])

	e := echo.New()

	e.POST("/token",handler.Token)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Debug = true
	
	api := e.Group("/api")

	// Configure middleware with the custom claims type
	api.Use(middleware.JWT([]byte(utils.Config.Secret)))

	api.GET("/profile",handler.GetProfile)
	api.GET("/:balancesheet/:frequency/financials", handler.GetFinancials)
	api.GET("/candle",handler.GetCandle)

	e.Logger.Fatal(e.Start(":8000"))
}

