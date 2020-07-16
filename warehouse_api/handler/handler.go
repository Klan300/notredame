package handler

import (
    "net/http"
    // "time"
    "fmt"
    "log"

    "go.mongodb.org/mongo-driver/bson"

    "github.com/dgrijalva/jwt-go"
    "github.com/labstack/echo"

    "ava.fund/alpha/Post-Covid/warehouse_api/database"
)


func Token( c echo.Context) error {
	username := c.QueryParam("username")

	// Throws unauthorized error
	if username != username {
		return echo.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = username
	claims["exp"] = nil

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"username" : username,
		"token": t,
	})
}

func GetProfile( c echo.Context) error {

	symbol := c.QueryParam("symbol")
	exchange := c.QueryParam("exchange")

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

func GetFinancials( c echo.Context) error {

	exchange := c.QueryParam("exchange")

	collectionName := fmt.Sprintf("%s_financials",exchange)
	collection,ctx := helper.ConnectDB(collectionName)

	fmt.Println(collection.Name)

	symbol := c.QueryParam("symbol")
	frequency := c.Param("freq")
	statement := c.Param("statement")

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

func GetCandle( c echo.Context) error {

	symbol := c.QueryParam("symbol")
	exchange := c.QueryParam("exchange")

	collectionName := fmt.Sprintf("%s_candle",exchange)
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

