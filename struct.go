package main

import "go.mongodb.org/mongo-driver/bson/primitive"

// type con la información de los libros

type Book struct {
	ID        primitive.ObjectID `bson:"_id"`
	Nombre    string             `bson:"nombre"`
	Precio    int                `bson:"precio"`
	Ubicacion []struct {
		Bodega string `bson:"bodega"`
		Stock  int    `bson:"stock"`
	} `bson:"ubicacion"`
}

type AllBooks []Book
