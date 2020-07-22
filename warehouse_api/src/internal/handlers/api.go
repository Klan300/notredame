package handlers

import (
	"net/http"

	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"

	"ava.fund/alpha/Post-Covid/warehouse_api/src/internal/utils"
)


func Profile( c echo.Context) error {

    symbol := c.QueryParam("symbol")

    filter := bson.M{"symbol" : symbol}
    data := bson.M{}

    exchange := c.QueryParam("exchange")
    collectionName := fmt.Sprintf("%s_profile",exchange)
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
            bson.M{"symbol"   : symbol},
            bson.M{"statement": statement},
            bson.M{"frequency": frequency},
        }}

    data := bson.M{}

    exchange := c.QueryParam("exchange")
    collectionName := fmt.Sprintf("%s_profile",exchange)
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

    filter := bson.M{"symbol" : symbol}
    data := bson.M{}

    exchange := c.QueryParam("exchange")
    collectionName := fmt.Sprintf("%s_candle",exchange)
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

func Search( c echo.Context) error {

    database, ctx := utils.Database()

    filterCollection := bson.M{"name": bson.M{"$regex": "securities"}}
    collections, err := database.ListCollectionNames(ctx,filterCollection)
    if err != nil {
        utils.Error("[api.go] Get Symbol",err)
    }

    dataSet := []bson.M{}

    search := c.QueryParam("search")
    // exchange := c.QueryParam("exchange")
    filter := bson.M{"$or": 
                []interface{}{
                    bson.M{"description" : bson.M{"$regex" :search}},
                    bson.M{"symbol" : bson.M{"$regex" :search}√},
                },
            }


    for _, collection := range collections {
        data := []bson.M{}
        securities , err  := database.Collection(collection).Find(ctx,filter)
        if err != nil {
            utils.Error("[api.go] Get Symbol",err)
        }

        for securities.Next(ctx) {
            var security bson.M
            if err =  securities.Decode(&security); err != nil {
                log.Fatal(err)
            }
            data = append(data,security)
        }

        dataSet = append(dataSet,data...)
    }

    return c.JSON(http.StatusOK, dataSet)
}

