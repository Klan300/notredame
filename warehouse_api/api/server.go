package main

import (
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"

	"ava.fund/alpha/Post-Covid/warehouse_api/database"
	// "ava.fund/alpha/Post-Covid/warehouse_api/model"
	"github.com/labstack/echo"
)

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

func main() {

	e := echo.New()

	e.GET("/api/:exchange/profile",getProfile)

	e.GET("/api/:exchange/financials",getFinancials)

	e.Logger.Fatal(e.Start(":8000"))
}

