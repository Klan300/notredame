package workers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"ava.fund/alpha/Post-Covid/warehouse_cloning/src/internal/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)



func process(request *http.Request, exchange string) []Security {
    client := http.Client{}

    response, err := client.Do(request)
    if err != nil {
        utils.Error("[reader.go] %v", err)
    }
    
    var securities []Security
    data, _ := ioutil.ReadAll(response.Body)
    err = json.Unmarshal(data, &securities)
    if err != nil {
        utils.Error("[reader.go] %v", err)
    }

    for i := range securities {
        securities[i].Exchange = strings.ToLower(exchange)
        securities[i].Symbol = strings.ToLower(securities[i].Symbol)
    }

    return securities

}


func RetrieveSecurities() []Security {

    utils.Debug("[reader.go] Begin")
    database, ctx := utils.Database()
    defer utils.Debug("[reader.go] Disconnect from database server")
    defer database.Client().Disconnect(ctx)
    
    
    var operations []mongo.WriteModel
    var securities []Security
    for _, exchange := range utils.Config.Exchanges {

        endpoint := fmt.Sprintf(
            utils.Endpoints["symbol"], 
            utils.Config.Source.Host, 
            exchange, 
            utils.Config.Source.Token)


        request, err := http.NewRequest("GET", endpoint, nil)
        if err != nil {
            utils.Error("[reader.go] %v", err)
        }

        securitiesForExchange := process(request, exchange)
        securities = append(securities, securitiesForExchange...)

        var operationsForExchange []mongo.WriteModel
        for _, security := range securitiesForExchange {
            
            filter := bson.M{"symbol": security.Symbol}

            operation := mongo.NewReplaceOneModel()
            operation.SetFilter(filter)
            operation.SetReplacement(security)
            operation.SetUpsert(true)

            operationsForExchange = append(operationsForExchange, operation)

        }
        
        collectionName := fmt.Sprintf("%s_securities", strings.ToLower(exchange))
        collectionInstance := database.Collection(collectionName)

        operations = append(operations, operationsForExchange...)

        result, err := collectionInstance.BulkWrite(ctx, operationsForExchange)        
        if err != nil {
            utils.Error("[reader.go] %v", err)
        } else {
            utils.Debug("[reader.go] BulkWrite: %v", result)
        }

    }
    collectionAllSecurity := database.Collection("securities")

    result, err := collectionAllSecurity.BulkWrite(ctx, operations)        
        if err != nil {
            utils.Error("[reader.go] %v", err)
        } else {
            utils.Debug("[reader.go] BulkWrite: %v", result)
        }

    utils.Debug("[reader.go] End")

    return securities
}