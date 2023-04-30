package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connection(uri string) *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)

	}

	// Comprobar la conexión
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)

	}

	fmt.Println("Conexión exitosa a MongoDB")
	return client
}
