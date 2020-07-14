package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
	"net/http"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB : This is helper function to connect mongoDB
// If you want to export your function. You must to start upper case function name. Otherwise you won't see your function when you import that on other class.
func ConnectDB(collectionName string) (*mongo.Collection, context.Context) {

	url := "mongodb://localhost:27017" 
	// Set client options
	clientOptions := options.Client().ApplyURI(url)

	ctx, _ := context.WithTimeout(context.Background(), 1000000000*time.Second)
	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("StockExchange").Collection(collectionName)

	fmt.Println("Connected to MongoDB!")

	return collection,ctx
}

// ErrorResponse : This is error model.
type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

// GetError : This is helper function to prepare error model.
// If you want to export your function. You must to start upper case function name. Otherwise you won't see your function when you import that on other class.
func GetError(err error, w http.ResponseWriter) {

	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}