package handlers

import (
	"net/http"
	"strings"

	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"ava.fund/alpha/Post-Covid/warehouse_api/src/internal/utils"
)


func Profile( c echo.Context) error {

    symbol := c.QueryParam("symbol")

    filter := bson.M{"symbol" : strings.ToLower(symbol)}
    data := bson.M{}

    exchange := c.QueryParam("exchange")
    collectionName := fmt.Sprintf("%s_profile",strings.ToLower(exchange))
    utils.Debug("[api.go] find in %s",collectionName)

    database, ctx := utils.Database()
    err := database.
        Collection(collectionName).
        FindOne( ctx, filter).
        Decode(&data)

    if err != nil {
        fmt.Println(err)
        return c.NoContent(http.StatusNotFound)
    }

    err = database.Client().Disconnect(ctx)

    if err != nil {
        log.Panicln(err)
    }

    return c.JSON(http.StatusOK, data)
}

func Financials( c echo.Context) error {

    symbol := c.QueryParam("symbol")
    frequency := c.Param("frequency")
    statement := c.Param("statement")

    filter := bson.M{
        "$and": []bson.M{
            bson.M{"symbol"   : strings.ToLower(symbol)},
            bson.M{"statement": strings.ToLower(statement)},
            bson.M{"frequency": strings.ToLower(frequency)},
        }}

    data := bson.M{}

    exchange := c.QueryParam("exchange")
    collectionName := fmt.Sprintf("%s_profile",strings.ToLower(exchange))
    utils.Debug("[api.go] find in %s",collectionName)
    
    database, ctx := utils.Database()
    err := database.
        Collection(collectionName).
        FindOne(ctx, filter).
        Decode(&data)

    if err != nil {
        fmt.Println("err")
        return c.NoContent(http.StatusNotFound)
    }

    err = database.Client().Disconnect(ctx)

    if err != nil {
        log.Panicln(err)
    }

    return c.JSON(http.StatusOK, data)
}

func Candle( c echo.Context) error {

    symbol := c.QueryParam("symbol")

    filter := bson.M{"symbol" : strings.ToLower(symbol)}
    data := bson.M{}

    exchange := c.QueryParam("exchange")
    collectionName := fmt.Sprintf("%s_candle",strings.ToLower(exchange))
    utils.Debug("[api.go] find in %s",collectionName)

    database, ctx := utils.Database()
    err := database.
        Collection(collectionName).
        FindOne( ctx, filter).
        Decode(&data)

    if err != nil {
        utils.Error("[api.go] get candle", err)
        return c.NoContent(http.StatusNotFound)
    }

    err = database.Client().Disconnect(ctx)

    if err != nil {
        log.Panicln(err)
    }

    return c.JSON(http.StatusOK, data)
}

func Symbol ( c echo.Context) error{

    exchange := c.QueryParam("exchange")

    collectionName := fmt.Sprintf("%s_securities",strings.ToLower(exchange))

    filter := bson.M{}

    findOptions := options.Find()
    findOptions.SetProjection(bson.M{"symbol":1,"_id" : 0})

    database, ctx := utils.Database()
    symbols,err := database.Collection(collectionName).Find(ctx,filter,findOptions)
    if err != nil {
        utils.Error("[api.go] get symbol", err)
    }

    var data  []bson.M
    symbols.All(ctx,&data)
    
    err = database.Client().Disconnect(ctx)

    if err != nil {
        log.Panicln(err)
    }

    return c.JSON(http.StatusOK, data)
}

