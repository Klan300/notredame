package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"fmt"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"ava.fund/alpha/Post-Covid/datamart_api/src/internal/utils"
)

type Score struct {
    Exchange string
    Symbol string 
    Expert string
    Tag string
    Data interface{}
}

func Update(c echo.Context) error {

    expert := c.QueryParam("expert")
    tag := c.QueryParam("tag")

    if expert == "" || tag == ""{
        return c.NoContent(http.StatusBadRequest)
    }

    collectionName := fmt.Sprintf("scores")
    database, ctx := utils.Database()
    defer utils.Debug("[api.go] Disconnect from database server")
    defer database.Client().Disconnect(ctx)

    var scores []Score
    err :=  json.NewDecoder(c.Request().Body).Decode(&scores)
    if err != nil {
        return c.JSON(http.StatusBadRequest,err)
    }

    var lastestScores []Score
    utils.Debug("[api.go] decorate struct score")
    for i := range scores {

        scores[i].Exchange= strings.ToLower(scores[i].Exchange)
        scores[i].Symbol = strings.ToLower(scores[i].Symbol)
        scores[i].Expert = strings.ToLower(expert)
        scores[i].Tag = "lastest"

        lastestScores = append(lastestScores,scores[i])
        scores[i].Tag = strings.ToLower(tag)

    }

    scores = append(scores,lastestScores...)

    var operationsForScore []mongo.WriteModel
    for _, score := range scores {

        filter := bson.M{
            "$and": []bson.M{
                bson.M{"symbol"   : strings.ToLower(score.Symbol)},
                bson.M{"exchange" : strings.ToLower(score.Exchange)},
                bson.M{"expert"   : strings.ToLower(expert)},
                bson.M{"tag"      : strings.ToLower(score.Tag)},
            },
        }

        operation := mongo.NewUpdateOneModel()
        operation.SetFilter(filter)
        operation.SetUpdate(score)
        operation.SetUpsert(true)

        operationsForScore = append(operationsForScore, operation)

    }

    collectionInstance := database.Collection(collectionName)
    result, err := collectionInstance.BulkWrite(ctx, operationsForScore)
    if err != nil {
        utils.Error("[api.go] %v", err)
    } else {
        utils.Debug("[api.go] BulkWrite: %v", result)
    }

    return c.JSON(http.StatusOK,result)

}

func Replace(c echo.Context) error {

    expert := c.QueryParam("expert")
    tag := c.QueryParam("tag")

    if expert == "" || tag == ""{
        return c.NoContent(http.StatusBadRequest)
    }

    collectionName := fmt.Sprintf("scores")
    database, ctx := utils.Database()
    defer utils.Debug("[api.go] Disconnect from database server")
    defer database.Client().Disconnect(ctx)

    var scores []Score
    err :=  json.NewDecoder(c.Request().Body).Decode(&scores)
    if err != nil {
        utils.Debug("[api.go] %v",err)
        return c.NoContent(http.StatusBadRequest)
    }

    utils.Debug("[api.go] List scores")

    var lastestScores []Score
    for i := range scores {

        scores[i].Exchange= strings.ToLower(scores[i].Exchange)
        scores[i].Symbol = strings.ToLower(scores[i].Symbol)
        scores[i].Expert = strings.ToLower(expert)
        scores[i].Tag = "lastest"

        lastestScores = append(lastestScores,scores[i])
        scores[i].Tag = strings.ToLower(tag)
    }

    scores = append(scores,lastestScores...)

    var operationsForScore []mongo.WriteModel
    for _, score := range scores {
        
        filter := bson.M{
            "$and": []bson.M{
                bson.M{"symbol"   : strings.ToLower(score.Symbol)},
                bson.M{"exchange": strings.ToLower(score.Exchange)},
                bson.M{"expert": strings.ToLower(expert)},
                bson.M{"tag": strings.ToLower(score.Tag)},
            },
        }

        operation := mongo.NewReplaceOneModel()
        operation.SetFilter(filter)
        operation.SetReplacement(score)
        operation.SetUpsert(true)

        operationsForScore = append(operationsForScore, operation)

    }

    collectionInstance := database.Collection(collectionName)
    result, err := collectionInstance.BulkWrite(ctx, operationsForScore)
    if err != nil {
        utils.Debug("[api.go] %v", err)
        return c.NoContent(http.StatusInternalServerError)
    } else {
        utils.Debug("[api.go] BulkWrite: %v", result)
    }

    return c.JSON(http.StatusOK,result)

}

func Find(c echo.Context) error {

    expert 	 := strings.ToLower(c.QueryParam("expert"))
    tag 	 := strings.ToLower(c.QueryParam("tag"))
    exchange := strings.ToLower(c.QueryParam("exchange"))
    symbol   := strings.ToLower(c.QueryParam("symbol"))

    if expert == "" {
        return c.NoContent(http.StatusBadRequest)
    }

    if tag == "" {
        tag = "lastest"
    }

    collectionName := "scores"
    database, ctx := utils.Database()
    defer utils.Debug("[api.go] Disconnect from database server")
    defer database.Client().Disconnect(ctx)

    var andCondition []bson.M

    switch {
        case exchange != "":
            andCondition = append(andCondition,
                            bson.M{"exchange": exchange})
            fallthrough
        case symbol != "":
            andCondition = append(andCondition, 
                            bson.M{"symbol": symbol})
    }

    andCondition = append(andCondition, bson.M{"tag": tag}, bson.M{"expert": expert})

    filter := bson.M{
        "$and": andCondition,
    }

    cursor, err := database.
            Collection(collectionName).
            Find(ctx,filter)
    if err != nil {
        utils.Debug("[api.go] %v",err)
        return c.NoContent(http.StatusNotFound)
    }
    
    var data []Score
    cursor.All(ctx, &data)
    if err != nil {
        utils.Debug("[api.go] %v", err)
        return c.NoContent(http.StatusInternalServerError)
    }

    if data == nil {
        data = make([]Score, 0)
    }

    return c.JSON(http.StatusOK, data)
}

