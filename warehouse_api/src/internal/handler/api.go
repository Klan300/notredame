package handlers

import (
	"net/http"

	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"github.com/labstack/echo/v4"

	"ava.fund/alpha/Post-Covid/warehouse_api/src/internal/utils"
)


func GetProfile( c echo.Context) error {

	symbol := c.QueryParam("symbol")

	filter := bson.M{"symbol" : symbol}
	data := bson.M{}

	exchange := c.QueryParam("exchange")
	collectionName := fmt.Sprintf("%s_profile",exchange)
	utils.Debug(collectionName)

	client, Database, ctx := utils.Database()
	err := Database.Collection(collectionName).FindOne( ctx, filter).Decode(&data)

	if err != nil {
		fmt.Println(err)
		return c.NoContent(http.StatusNotFound)
	}

	err = client.Disconnect(ctx)

	if err != nil {
		log.Panicln(err)
	}

	return c.JSON(http.StatusOK, data)
}

func GetFinancials( c echo.Context) error {

	symbol := c.QueryParam("symbol")
	frequency := c.Param("frequency")
	statement := c.Param("statement")

	filter := bson.M{
		"$and": []bson.M{
			bson.M{"symbol"   : symbol},
			bson.M{"statement": statement},
			bson.M{"frequency": frequency},
		}}

	data := bson.M{}

	exchange := c.QueryParam("exchange")
	collectionName := fmt.Sprintf("%s_profile",exchange)
	utils.Debug(collectionName)
	
	client, database, ctx := utils.Database()
	err := database.Collection(collectionName).FindOne(ctx, filter).Decode(&data)

	if err != nil {
		fmt.Println("err")
		return c.NoContent(http.StatusNotFound)
	}

	err = client.Disconnect(ctx)

	if err != nil {
		log.Panicln(err)
	}

	return c.JSON(http.StatusOK, data)
}

func GetCandle( c echo.Context) error {

	symbol := c.QueryParam("symbol")

	filter := bson.M{"symbol" : symbol}
	data := bson.M{}

	exchange := c.QueryParam("exchange")
	collectionName := fmt.Sprintf("%s_candle",exchange)
	utils.Debug(collectionName)

	client, Database, ctx := utils.Database()
	err := Database.Collection(collectionName).FindOne( ctx, filter).Decode(&data)

	if err != nil {
		fmt.Println(err)
		return c.NoContent(http.StatusNotFound)
	}

	err = client.Disconnect(ctx)

	if err != nil {
		log.Panicln(err)
	}

	return c.JSON(http.StatusOK, data)
}

