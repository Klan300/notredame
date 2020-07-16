package main

import (
    "net/http"
    "time"
    "fmt"
    "log"

    "go.mongodb.org/mongo-driver/bson"

    "github.com/dgrijalva/jwt-go"
    "github.com/labstack/echo"

    "ava.fund/alpha/Post-Covid/warehouse_api/database"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}


func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Throws unauthorized error
	if username != "jon" || password != "shhh!" {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		"Jon Snow",
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func getProfile( c echo.Context) error {

	symbol := c.QueryParam("symbol")
	exchange := c.Param("exchange")

	collectionName := fmt.Sprintf("%s_profile",exchange)
	collection,ctx := helper.ConnectDB(collectionName)

	filter := bson.M{"symbol" : symbol}
	data := bson.M{}
	err := collection.FindOne( ctx, filter).Decode(&data)

	if err != nil {
		fmt.Println("err")
		return c.NoContent(http.StatusNotFound)
	}

	err = collection.Database().Client().Disconnect(ctx)

	if err != nil {
		log.Panicln(err)
	}

	return c.JSON(http.StatusOK, data)
}

func getFinancials( c echo.Context) error {

	exchange := c.Param("exchange")

	collectionName := fmt.Sprintf("%s_financials",exchange)
	collection,ctx := helper.ConnectDB(collectionName)

	fmt.Println(collection.Name)

	symbol := c.QueryParam("symbol")
	frequency := c.QueryParam("freq")
	statement := c.QueryParam("statement")

	filter := bson.M{
		"$and": []bson.M{
			bson.M{"symbol"   : symbol},
			bson.M{"statement": statement},
			bson.M{"frequency": frequency},
		}}

	data := bson.M{}
	err := collection.FindOne(ctx, filter).Decode(&data)

	if err != nil {
		fmt.Println("err")
		return c.NoContent(http.StatusNotFound)
	}

	err = collection.Database().Client().Disconnect(ctx)

	if err != nil {
		log.Panicln(err)
	}

	return c.JSON(http.StatusOK, data)
}

