package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Db() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb+srv://naveengoud:Naveen@cluster0.fkhjj.mongodb.net/Borderfree?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("mongodb connected")
	return client
}

var ProductCollection = Db().Database("Borderfree").Collection("products")
var UserCollection = Db().Database("Borderfree").Collection("users")