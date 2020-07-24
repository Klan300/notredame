package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"ava.fund/alpha/Notre-Dame/datamart_api/src/internal/utils"
)



func Update(c echo.Context) error {

    expert := strings.ToLower(c.QueryParam("expert"))
    tag    := strings.ToLower(c.QueryParam("tag"))

    if expert == "" {
        return c.NoContent(http.StatusBadRequest)
    }

    utils.Debug("[api.go] Read scores from request")
    var scores []utils.Score
    err := json.NewDecoder(c.Request().Body).Decode(&scores)
    if err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }

    if tag != "" {
        for i := range scores {
            scores[i].Exchange = strings.ToLower(scores[i].Exchange)
            scores[i].Symbol   = strings.ToLower(scores[i].Symbol)
            scores[i].Expert   = expert
            scores[i].Tag      = tag
        }
    }

    if tag != "latest" {
        for _, score := range scores {
            score.Tag = "latest"
            scores = append(scores, score)
        }
    }

    database, ctx  := utils.Database()
    defer utils.Debug("[api.go] Disconnect from database server")
    defer database.Client().Disconnect(ctx)

    
    utils.Debug("[api.go] Update scores to database")
    var operations []mongo.WriteModel
    for _, score := range scores {

        filter := bson.M{
            "$and": []bson.M{
                bson.M{"symbol"  : score.Symbol},
                bson.M{"exchange": score.Exchange},
                bson.M{"expert"  : score.Expert},
                bson.M{"tag"     : score.Tag},
            },
        }

        operation := mongo.NewUpdateOneModel()
        operation.SetFilter(filter)
        operation.SetUpdate(score)
        operation.SetUpsert(true)

        operations = append(operations, operation)

    }

    collectionName     := "scores"
    collectionInstance := database.Collection(collectionName)
    result, err := collectionInstance.BulkWrite(ctx, operations)
    if err != nil {
        utils.Debug("[api.go] %v", err)
        return c.JSON(http.StatusInternalServerError, err)
    } else {
        utils.Debug("[api.go] BulkWrite: %d record upserted", result.UpsertedCount)
        return c.NoContent(http.StatusOK)
    }

    

}


func Replace(c echo.Context) error {

    expert := strings.ToLower(c.QueryParam("expert"))
    tag    := strings.ToLower(c.QueryParam("tag"))

    if expert == "" {
        return c.NoContent(http.StatusBadRequest)
    }

    utils.Debug("[api.go] Read scores from request")
    var scores []utils.Score
    err :=  json.NewDecoder(c.Request().Body).Decode(&scores)
    if err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }

    if tag != "" {
        for i := range scores {
            scores[i].Exchange = strings.ToLower(scores[i].Exchange)
            scores[i].Symbol   = strings.ToLower(scores[i].Symbol)
            scores[i].Expert   = expert
            scores[i].Tag      = tag
        }
    }

    if tag != "latest" {
        for _, score := range scores {
            score.Tag = "latest"
            scores = append(scores, score)
        }
    }

    database, ctx  := utils.Database()
    defer utils.Debug("[api.go] Disconnect from database server")
    defer database.Client().Disconnect(ctx)

    
    utils.Debug("[api.go] Replace scores in database")
    var operations []mongo.WriteModel
    for _, score := range scores {

        filter := bson.M{
            "$and": []bson.M{
                bson.M{"symbol"  : score.Symbol},
                bson.M{"exchange": score.Exchange},
                bson.M{"expert"  : score.Expert},
                bson.M{"tag"     : score.Tag},
            },
        }

        operation := mongo.NewReplaceOneModel()
        operation.SetFilter(filter)
        operation.SetReplacement(score)
        operation.SetUpsert(true)

        operations = append(operations, operation)

    }

    collectionName     := "scores"
    collectionInstance := database.Collection(collectionName)
    result, err := collectionInstance.BulkWrite(ctx, operations)
    if err != nil {
        utils.Debug("[api.go] %v", err)
        return c.JSON(http.StatusInternalServerError, err)
    } else {
        utils.Debug("[api.go] BulkWrite: %d record upsert", result.UpsertedCount)
        return c.NoContent(http.StatusOK)    
    }
}


func Find(c echo.Context) error {

    expert   := strings.ToLower(c.QueryParam("expert"))
    tag      := strings.ToLower(c.QueryParam("tag"))
    exchange := strings.ToLower(c.QueryParam("exchange"))
    symbol   := strings.ToLower(c.QueryParam("symbol"))

    if expert == "" {
        return c.NoContent(http.StatusBadRequest)
    }
    if tag == "" {
        tag = "latest"
    }

    database, ctx := utils.Database()
    defer utils.Debug("[api.go] Disconnect from database server")
    defer database.Client().Disconnect(ctx)

    utils.Debug("[api.go] Find scores from %s:%s", expert, tag)
    conditions := []bson.M{
        {"expert": expert},
        {"tag"   : tag},
    }

    if exchange != "" {
        conditions = append(conditions, bson.M{"exchange": exchange})
    }
    if symbol != "" {
        conditions = append(conditions, bson.M{"symbol": symbol})
    }

    collectionName := "scores"
    filter         := bson.M{ "$and": conditions}
    cursor, err    := database.
        Collection(collectionName).
        Find(ctx, filter)

    if err != nil {
        utils.Debug("[api.go] %v", err)
        return c.JSON(http.StatusInternalServerError, err)
    }
    
    var data []utils.Score
    cursor.All(ctx, &data)
    if err != nil {
        utils.Debug("[api.go] %v", err)
        return c.JSON(http.StatusInternalServerError, err)
    }

    if data == nil {
        data = make([]utils.Score, 0)
    }

    return c.JSON(http.StatusOK, data)
}

