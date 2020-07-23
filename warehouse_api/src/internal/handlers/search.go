package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"ava.fund/alpha/Post-Covid/warehouse_api/src/internal/utils"
)

func Search(c echo.Context) error {

	symbol := strings.ToLower(c.QueryParam("symbol"))
	text := strings.ToLower(c.QueryParam("text"))
	limit := c.QueryParam("limit")

	collectionName := "securities"
	database, ctx := utils.Database()
	defer utils.Debug("[search.go] Disconnect from database server")
	defer database.Client().Disconnect(ctx)

	utils.Debug("[search.go] Search for symbols in %s", collectionName)
	var filter bson.M
	switch {

	case symbol == "" && text == "":
		utils.Debug("[search.go] Empty {symbol} and {text}")
		return c.NoContent(http.StatusBadRequest)

	case symbol != "" && text == "":
		filter = bson.M{"symbol": bson.M{"$regex": symbol}}

	case symbol == "" && text != "":
		filter = bson.M{"$or": []interface{}{
			bson.M{"symbol": bson.M{"$regex": text}},
			bson.M{"description": bson.M{"$regex": text}},
		}}

	case symbol != "" && text != "":
		filter = bson.M{"$and": []interface{}{
			bson.M{"symbol": bson.M{"$regex": symbol}},
			bson.M{"description": bson.M{"$regex": text}},
		}}

	}

	sorter := options.Find()
	sorter.SetSort(bson.M{"symbol": 1})
	if limit != "" {
		limit, err := strconv.ParseInt(limit, 10, 64)
		if err != nil {
			utils.Debug("[search.go] %v", err)
			return c.NoContent(http.StatusBadRequest)
		}
		sorter.SetLimit(limit)
	}

	cursor, err := database.
		Collection(collectionName).
		Find(ctx, filter, sorter)
	if err != nil {
		utils.Debug("[search.go] %v", err)
		return c.NoContent(http.StatusNotFound)
	}

	var data []bson.M
	cursor.All(ctx, &data)
	if err != nil {
		utils.Error("[search.go] %v", err)
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, data)
}
