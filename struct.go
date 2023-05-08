package main

import "go.mongodb.org/mongo-driver/bson/primitive"

// type con la información de los libros

type Location struct {
	Bodega string `bson:"bodega"`
	Stock  int    `bson:"stock"`
}

type Book struct {
	ID          primitive.ObjectID `bson:"_id"`
	Nombre      string             `bson:"nombre"`
	Autor       string             `bson:"autor"`
	Descripcion string             `bson:"descripcion"`
	Imagen      string             `bson:"imagen"`
	Precio      int                `bson:"precio"`
	Ubicacion   []Location         `bson:"ubicacion"`
}

type AllBooks []Book
