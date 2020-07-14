package workers

import (
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	// "ava.fund/alpha/Post-Covid/warehouse_api/model"
	"ava.fund/alpha/Post-Covid/warehouse_cloning/src/internal/utils"
)

func Writer(responses chan *Response) {
    utils.Debug("[writer.go] Begin")

    client, database, ctx := utils.Database()

    go func() {
        defer utils.Debug("[writer.go] Disconnect from database server")
        defer client.Disconnect(ctx)
        
        for {
            select {
            case response, more := <-responses:
                if !more {
                    utils.Debug("[writer.go] Terminate the writer")
                    return
                }

                var data interface{}
                var filter bson.M
                var replace bson.M
                json.Unmarshal(response.Data, &data)


                switch response.Request.Document {
                    case "profile":
                        filter = bson.M{"symbol": response.Request.Symbol}
                        replace = bson.M{
                            "symbol": response.Request.Symbol,
                            "data"  : data}

                    case "financials":
                        filter = bson.M{
                            "$and": []bson.M{
                                bson.M{"symbol"   : response.Request.Symbol},
                                bson.M{"statement": response.Request.Statement},
                                bson.M{"frequency": response.Request.Frequency},
                            }}

                        replace = bson.M{
                            "symbol"   : response.Request.Symbol,
                            "frequency": response.Request.Frequency,
                            "statement": response.Request.Statement,
                            "data"     : data}
                }

                collectionName := fmt.Sprintf("%s_%s", response.Request.Exchange, response.Request.Document)
                collectionInstance := database.Collection(collectionName)

                options := options.Replace().SetUpsert(true)
                result, err := collectionInstance.ReplaceOne(ctx, filter, replace, options)
                if err != nil {
                    utils.Error("[writer.go] %v", err)
                } else {
                    utils.Debug("[writer.go] ReplaceOne: %v", result)
                }

            }
        }
    }()
    utils.Debug("[writer.go] End")
}
