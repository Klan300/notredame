package utils

import (
    "context"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func Database() ( *mongo.Database, context.Context){
    Debug("[database.go] Begin")

    Debug("[database.go] Create a database client")
    client, err := mongo.NewClient(options.
        Client().
        ApplyURI(Config.Target.Host).
        SetAuth(
        options.Credential{
            Username: Config.Target.Username, 
            Password: Config.Target.Password,
        }))
    if err != nil {
        Error("[database.go] %v", err)
    }


    Debug("[database.go] Create a connection to %s", Config.Target.Host)
    ctx, _ := context.WithTimeout(context.Background(), 24 * time.Hour)
    err = client.Connect(ctx)
    if err != nil {
        Error("[database.go] %v", err)
    }

    database := client.Database(Config.Target.Database)
    return database, ctx
}