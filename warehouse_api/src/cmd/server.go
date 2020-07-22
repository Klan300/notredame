package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"ava.fund/alpha/Post-Covid/warehouse_api/src/internal/handlers"
	"ava.fund/alpha/Post-Covid/warehouse_api/src/internal/utils"
)

func main() {

	if len(os.Args) != 2 {
        fmt.Printf("Usage: %s <config>\n", os.Args[0])
        os.Exit(0)
    }
    utils.LoadConfig(os.Args[1])

	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
    e.Use(middleware.CORS())
    e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
        Format : "${time_rfc3339}: method=${method}, remote=${remote_ip}, domain=${host}, uri=${uri}, status=${status}\n",
        Skipper: middleware.DefaultSkipper,
        Output : utils.Config.Logging.Outputs(),
	}))
	
	e.POST("/token", handlers.Token)
	
	api := e.Group("/api")
	
	api.Use(middleware.JWT([]byte(utils.Config.Authen.Secret)))

	api.GET("/profile", handlers.Profile)
	api.GET("/candle", handlers.Candle)
	api.GET("/financials/:statement/:frequency", handlers.Financials) 
	api.GET("/symbols",handlers.Symbol)
	api.GET("/search", handlers.Search)

	e.Logger.Fatal(e.Start(utils.Config.Target.Host))
}

