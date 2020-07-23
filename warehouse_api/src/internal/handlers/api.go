package handlers

import (
	"net/http"
	"strings"

	"fmt"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"ava.fund/alpha/Post-Covid/warehouse_api/src/internal/utils"
)

func Profile(c echo.Context) error {
	symbol := strings.ToLower(c.QueryParam("symbol"))
	exchange := strings.ToLower(c.QueryParam("exchange"))

	collectionName := fmt.Sprintf("%s_profile", exchange)
	database, ctx := utils.Database()
	defer utils.Debug("[api.go] Disconnect from database server")
	defer database.Client().Disconnect(ctx)

	utils.Debug("[api.go] Find %s in %s", symbol, collectionName)
	filter := bson.M{"symbol": symbol}
	data := bson.M{}
	err := database.
		Collection(collectionName).
		FindOne(ctx, filter).
		Decode(&data)

	if err != nil {
		utils.Debug("[api.go] %v", err)
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, data)
}

func Financials(c echo.Context) error {

	symbol := strings.ToLower(c.QueryParam("symbol"))
	exchange := strings.ToLower(c.QueryParam("exchange"))
	frequency := strings.ToLower(c.Param("frequency"))
	statement := strings.ToLower(c.Param("statement"))

	collectionName := fmt.Sprintf("%s_financials", exchange)
	database, ctx := utils.Database()
	defer utils.Debug("[api.go] Disconnect from database server")
	defer database.Client().Disconnect(ctx)

	utils.Debug("[api.go] Find %s/%s/%s in %s", statement, frequency, symbol, collectionName)
	filter := bson.M{
		"$and": []bson.M{
			bson.M{"symbol": symbol},
			bson.M{"statement": statement},
			bson.M{"frequency": frequency},
		},
	}
	data := bson.M{}
	err := database.
		Collection(collectionName).
		FindOne(ctx, filter).
		Decode(&data)

	if err != nil {
		utils.Debug("[api.go] %v", err)
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, data)
}

func Candle(c echo.Context) error {

	symbol := strings.ToLower(c.QueryParam("symbol"))
	exchange := strings.ToLower(c.QueryParam("exchange"))

	collectionName := fmt.Sprintf("%s_candle", exchange)
	database, ctx := utils.Database()
	defer utils.Debug("[api.go] Disconnect from database server")
	defer database.Client().Disconnect(ctx)

	utils.Debug("[api.go] Find %s in %s", symbol, collectionName)
	filter := bson.M{"symbol": symbol}
	data := bson.M{}
	err := database.
		Collection(collectionName).
		FindOne(ctx, filter).
		Decode(&data)

	if err != nil {
		utils.Debug("[api.go] %v", err)
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, data)
}

func Symbols(c echo.Context) error {

	exchange := strings.ToLower(c.QueryParam("exchange"))

	collectionName := fmt.Sprintf("%s_securities", exchange)
	database, ctx := utils.Database()
	defer utils.Debug("[api.go] Disconnect from database server")
	defer database.Client().Disconnect(ctx)

	utils.Debug("[api.go] List symbols in %s", collectionName)
	filter := bson.M{}
	selecter := options.Find().SetProjection(bson.M{
		"_id":    0,
		"symbol": 1,
	})

	cursor, err := database.
		Collection(collectionName).
		Find(ctx, filter, selecter)
	if err != nil {
		utils.Debug("[api.go] %v", err)
		return c.NoContent(http.StatusNotFound)
	}

	var data []bson.M
	cursor.All(ctx,&data)
	if err != nil {
		utils.Error("[api.go] %v", err)
		return c.NoContent(http.StatusNotFound)
	}

	var symbols []interface{}
	for _, element := range data {
		for _, value := range element {
            symbols = append(symbols,value)
        }
	}

	return c.JSON(http.StatusOK, symbols)
}
