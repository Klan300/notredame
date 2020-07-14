package workers

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"

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
        securities[i].Exchange = exchange
    }

    return securities

}


func RetrieveSecurities() []Security {

    utils.Debug("[reader.go] Begin")
    client, database, ctx := utils.Database()
    defer utils.Debug("[reader.go] Disconnect from database server")
    defer client.Disconnect(ctx)
    
    

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

        var operations []mongo.WriteModel
        for _, security := range securitiesForExchange {
            
            filter := bson.M{"symbol": security.Symbol}

            operation := mongo.NewReplaceOneModel()
            operation.SetFilter(filter)
            operation.SetReplacement(security)
            operation.SetUpsert(true)

            operations = append(operations,operation)

        }
        
        collectionName := fmt.Sprintf("%s_securities", exchange)
        collectionInstance := database.Collection(collectionName)

        result, err := collectionInstance.BulkWrite(ctx, operations)        
        if err != nil {
            utils.Error("[reader.go] %v", err)
        } else {
            utils.Debug("[reader.go] BulkWrite: %v", result)
        }

    }

    utils.Debug("[reader.go] End")
    return securities
}