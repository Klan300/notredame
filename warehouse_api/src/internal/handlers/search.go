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


func Search( c echo.Context) error {

    database, ctx := utils.Database()
    text := c.QueryParam("text")
    symbol := c.QueryParam("symbol")
    limit := c.QueryParam("limit")
    // exchange := c.QueryParam("exchange")
    var filter bson.M

    switch {
        case symbol != "" && text != "":
            filter = bson.M{"$and": 
            []interface{}{
                bson.M{"description" : bson.M{"$regex" :text}},
                bson.M{"symbol" : bson.M{"$regex" :strings.ToLower(symbol)}},
            },
        }
        case symbol != "" && text == "":
            filter = bson.M{"symbol" : bson.M{"$regex" :strings.ToLower(symbol)}}
        case symbol == "" && text != "":
            filter = bson.M{"$or": 
                []interface{}{
                    bson.M{"description" : bson.M{"$regex" :text}},
                    bson.M{"symbol" : bson.M{"$regex" :text}},
                },
            }
        case symbol == "" && text == "":
            return c.NoContent(http.StatusBadRequest)
        }

    findOptions := options.Find()
    findOptions.SetSort(bson.M{ "symbol": 1})
    
    if limit != ""{
        limitInt,err := strconv.ParseInt(limit, 10, 64)

        if err != nil {
            utils.Error("[api.go] Get Symbol",err)
        }
        
        findOptions.SetLimit(limitInt)
    }

    securities,err := database.Collection("securities").Find( ctx, filter, findOptions)

    if err != nil {
        utils.Error("[api.go] Get Symbol",err)
    }

    data := []bson.M{}

    err = securities.All(ctx,&data)

    if err != nil {
        utils.Error("[api.go] Get Symbol",err)
    }

	return c.JSON(http.StatusOK, data)
}
