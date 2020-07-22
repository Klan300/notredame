package handlers

import (
	"net/http"
	"time"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"ava.fund/alpha/Post-Covid/warehouse_api/src/internal/utils"
)

func Token(c echo.Context) error {

    username := c.QueryParam("username")
    if !utils.Config.Authen.Exists(username) {
        return c.NoContent(http.StatusUnauthorized)
    }


    claims             := jwt.MapClaims{}
    claims["username"] = username
    claims["exp"]      = time.Date(2030,01,1,0,0,0,0,time.Local).Unix()
    time.Now().Day()

    token, err := jwt.
        NewWithClaims(jwt.SigningMethodHS256, claims).
        SignedString([]byte(utils.Config.Authen.Secret))

    if err != nil {
        log.Panicf("[token.go] %s\n", err)
    }

    return c.JSON(http.StatusOK, echo.Map{
        "username" : username,
        "token"    : token,
    })
}