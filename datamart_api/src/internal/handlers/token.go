package handlers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"ava.fund/alpha/Post-Covid/datamart_api/src/internal/utils"
)

func Token(c echo.Context) error {

	utils.Debug("[token.go] Begin")
	username := c.QueryParam("username")
	if !utils.Config.Authen.Exists(username) {
		utils.Debug("[token.go] Token requested for an unauthorized user %s", username)
		return c.NoContent(http.StatusUnauthorized)
	}

	expire, _ := time.Parse("2006-01-02", utils.Config.Authen.Expire)

	claims := jwt.MapClaims{}
	claims["username"] = username
	claims["exp"] = expire.Unix()

	utils.Debug("[token.go] Generate token for user %s", username)
	token, err := jwt.
		NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(utils.Config.Authen.Secret))

	if err != nil {
		utils.Error("[token.go] %v", err)
	}

	utils.Debug("[token.go] End")
	return c.JSON(http.StatusOK, echo.Map{
		"username": username,
		"token":    token,
	})
}
