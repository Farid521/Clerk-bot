package db

import (
	"context"
	"time"
	"errors"
	"log"
	"fmt"

	"clerk-bot/src/types"
	"clerk-bot/config"


	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MethodRead = "read"
	MethodWrite = "write"
)

func DbAcces (content types.UserMsg, method string) ([]byte,error) {
	// creating client
	clientOptions := options.Client().ApplyURI(config.DbUri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connecting to mongoDb atlas
	client, err := mongo.Connect(ctx, clientOptions)
	if err !=  nil {
		return nil, errors.New("fail in connectig to the mongodb atlas")
	}
	fmt.Println("connected to mongoDB Atlas")
	defer func () {
		if err := client.Disconnect(ctx); err != nil {
			log.Printf("failed in disconnecting the client %v\n", err)
		} 
	}()

	switch method {
	case MethodWrite:
		// inserting data
		collection := client.Database("DiscordBot-Kuliah").Collection("jadwal")
		userInput := content
		result, err := collection.InsertOne(ctx, userInput)
		if err != nil {
			return nil, errors.New("error in inserting data")
		}
		fmt.Printf("succesfully inserted the data. ID: %v", result.InsertedID)
		return nil, nil
	}
	return nil, nil
}