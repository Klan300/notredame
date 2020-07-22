package utils

import (
    "context"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func Database() (*mongo.Database, context.Context){
   
    Debug("[database.go] Begin")

    Debug("[database.go] Create a database client")
    client, err := mongo.NewClient(options.
        Client().
        ApplyURI(Config.Source.Host))
        // SetAuth(
        // options.Credential{
        //  Username: Config.Source.Username, 
        //  Password: Config.Source.Password,
        // }))

    
    if err != nil {
        Error("[database.go] %v", err)
    }


    Debug("[database.go] Create a connection to %s", Config.Source.Host)
    ctx, _ := context.WithTimeout(context.Background(), 24 * time.Hour)
    err = client.Connect(ctx)
    if err != nil {
        Error("[database.go] %v", err)
    }

    database := client.Database(Config.Source.Database)
    
    return database, ctx
}