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

type Factura struct {
	ID          primitive.ObjectID `bson:"_id"`
	Tipo        string             `bson:"tipo"`
	Numero      int                `bson:"numero"`
	Fecha       string             `bson:"fecha"`
	Cliente     string             `bson:"cliente"`
	Retira      string             `bson:"retira"`
	Rut         string             `bson:"rut"`
	Direccion   string             `bson:"direccion"`
	Email       string             `bson:"email"`
	NombreLibro string             `bson:"nombre"`
	Precio      int                `bson:"precio"`
	Cantidad    int                `bson:"cantidad"`
	Total       int                `bson:"total"`
}

type FacturaPost struct {
	Tipo        string `bson:"tipo"`
	Cliente     string `bson:"cliente"`
	Retira      string `bson:"retira"`
	Rut         string `bson:"rut"`
	Direccion   string `bson:"direccion"`
	Email       string `bson:"email"`
	NombreLibro string `bson:"nombre"`
	Precio      int    `bson:"precio"`
	Cantidad    int    `bson:"cantidad"`
	Total       int    `bson:"total"`
}

type Response struct {
	Mensaje string `bson:"mensaje"`
}
