package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

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

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
  

	e.POST("/token",handler.Token)
	
	api := e.Group("/api")

	api.Use(middleware.JWT([]byte(utils.Config.Secret)))

	api.GET("/profile",handler.GetProfile)
	api.GET("/:statement/:frequency/financials", handler.GetFinancials)
	api.GET("/candle",handler.GetCandle)

	e.Logger.Fatal(e.Start(utils.Config.Target.Host))
}

