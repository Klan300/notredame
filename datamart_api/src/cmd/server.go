package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"ava.fund/alpha/Post-Covid/datamart_api/src/internal/handlers"
	"ava.fund/alpha/Post-Covid/datamart_api/src/internal/utils"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <config>\n", os.Args[0])
		os.Exit(0)
	}
	utils.LoadConfig(os.Args[1])
	utils.Debug("[server.go] Begin")

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:  "${time_rfc3339}: FROM ${remote_ip} ${method} ${uri} RETURN ${status}\n",
		Skipper: middleware.DefaultSkipper,
		Output:  utils.Config.Logging.Outputs(),
	}))
	e.POST("/token", handlers.Token)

	api := e.Group("/api")
	api.Use(middleware.JWT([]byte(utils.Config.Authen.Secret)))
	api.GET("/find", handlers.Find)
	api.PUT("/update", handlers.Update)
	api.PUT("/replace", handlers.Replace)

	e.Logger.Fatal(e.Start(utils.Config.Target.Host))
	utils.Debug("[server.go] End")
}
