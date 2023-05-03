package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to API Ventabook"))

}

func getAllBooks(w http.ResponseWriter, r *http.Request) {

	client := connection(mongoInfo)
	coll := client.Database("ventabookDB").Collection("books")

	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Crear una lista para almacenar los libros
	var books []Book

	// Iterar sobre los documentos encontrados y agregarlos a la lista
	for cursor.Next(context.Background()) {
		var book Book
		if err := cursor.Decode(&book); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		books = append(books, book)
	}

	// Comprobar errores después de iterar
	if err := cursor.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convertir la lista de libros a JSON
	jsonData, err := json.Marshal(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Establecer la respuesta como un JSON
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

	defer cursor.Close(context.Background())

}

func getOneBook(w http.ResponseWriter, r *http.Request) {
	// genero conexion
	client := connection(mongoInfo)
	coll := client.Database("ventabookDB").Collection("books")

	// extraigo variables
	vars := mux.Vars(r)
	fmt.Printf("vars %s", vars)
	id := vars["id"]
	fmt.Printf("\nid %s\n", id)
	objectID, errs := primitive.ObjectIDFromHex(id)

	if errs != nil {
		// Maneja el error si el string no es un ObjectID válido
		panic(errs)
	}
	//genero filtro
	filter := bson.D{{Key: "_id", Value: objectID}}

	var results bson.M
	err := coll.FindOne(context.TODO(), filter).Decode(&results)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			http.NotFound(w, r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Convertir la lista de libros a JSON
	jsonData, err := json.Marshal(results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Establecer la respuesta como un JSON
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	defer client.Disconnect(context.Background())

}

func updateOneBook(w http.ResponseWriter, r *http.Request) {
	// genero conexion
	client := connection(mongoInfo)
	coll := client.Database("ventabookDB").Collection("books")

	// extraigo variables
	vars := mux.Vars(r)
	fmt.Printf("vars %s", vars)
	id := vars["id"]
	objectID, errID := primitive.ObjectIDFromHex(id)
	bodega := vars["bodega"]
	stock, errStock := strconv.Atoi(vars["stock"])

	// Maneja el error si el string no es un ObjectID válido
	if errID != nil {
		http.Error(w, "Invalid id value", http.StatusBadRequest)
		return
	}
	if errStock != nil {
		http.Error(w, "Invalid stock value", http.StatusBadRequest)
	}
	//genero filtro
	filter := bson.M{"_id": objectID, "ubicacion.bodega": bodega}

	update := bson.M{"$inc": bson.M{"ubicacion.$.stock": stock}}
	result, err := coll.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("Documents matched: %v\n", result.MatchedCount)
	fmt.Printf("Documents updated: %v\n", result.ModifiedCount)
	defer client.Disconnect(context.Background())
}
