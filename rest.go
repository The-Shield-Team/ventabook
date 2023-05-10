package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

func getOneBook(w http.ResponseWriter, r *http.Request, id string) {
	// genero conexion
	client := connection(mongoInfo)
	coll := client.Database("ventabookDB").Collection("books")

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
func getOneBookByNombre(w http.ResponseWriter, r *http.Request, nombre string) {
	// genero conexion
	client := connection(mongoInfo)
	coll := client.Database("ventabookDB").Collection("books")

	// // extraigo variables
	// vars := mux.Vars(r)

	// nombre := vars["nombre"]

	//genero filtro pattern es el patron que usare para hacer una busqueda parcial
	pattern := primitive.Regex{Pattern: ".*" + nombre + ".*", Options: "i"}
	filter := bson.D{{Key: "nombre", Value: pattern}}

	cursor, err := coll.Find(context.TODO(), filter)
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

func getOneBookByAutor(w http.ResponseWriter, r *http.Request, autor string) {
	// genero conexion
	client := connection(mongoInfo)
	coll := client.Database("ventabookDB").Collection("books")

	// // extraigo variables
	// vars := mux.Vars(r)

	// nombre := vars["autor"]

	//genero filtro pattern es el patron que usare para hacer una busqueda parcial
	pattern := primitive.Regex{Pattern: ".*" + autor + ".*", Options: "i"}
	filter := bson.D{{Key: "autor", Value: pattern}}

	cursor, err := coll.Find(context.TODO(), filter)
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

func getBookByAny(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	params := strings.Split(vars["params"], "/")

	var nombre, autor, id string
	for _, p := range params {
		if strings.HasPrefix(p, "nombre=") {
			nombre = strings.TrimPrefix(p, "nombre=")
		} else if strings.HasPrefix(p, "autor=") {
			autor = strings.TrimPrefix(p, "autor=")
		} else if strings.HasPrefix(p, "id=") {
			id = strings.TrimPrefix(p, "id=")
		}
	}

	if nombre == "" && autor == "" && id == "" {
		http.Error(w, "Debe proporcionar al menos un parámetro válido: nombre o autor", http.StatusBadRequest)
		return
	}
	if nombre != "" {
		getOneBookByNombre(w, r, nombre)
		return
	}
	if autor != "" {
		getOneBookByAutor(w, r, autor)
		return
	}
	if id != "" {
		getOneBook(w, r, id)
		return
	} else {
		http.Error(w, "Debe proporcionar al menos un parámetro válido: nombre o autor", http.StatusBadRequest)
		return
	}
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

func Facturar(w http.ResponseWriter, r *http.Request) {
	// genero conexion
	client := connection(mongoInfo)
	coll := client.Database("facturacion").Collection("documentos")

	// extraigo variables
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	// guardo valores en la variable FacturaPost
	var p FacturaPost
	err := decoder.Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// variable factura por crear	var

	var FacturaTipo = p.Tipo
	var FacturaNumero = 1
	var FacturaFecha = "Hoy"
	var FacturaCliente = p.Cliente
	var FacturaRetira = p.Retira
	var FacturaRut = p.Rut
	var FacturaDireccion = p.Direccion
	var FacturaEmail = p.Email
	var FacturaNombreLibro = p.NombreLibro
	var FacturaPrecio = p.Precio
	var FacturaCantidad = p.Cantidad
	var FacturaTotal = p.Total

	factura := Factura{
		ID:          primitive.NewObjectID(),
		Tipo:        FacturaTipo,
		Numero:      FacturaNumero,
		Fecha:       FacturaFecha,
		Cliente:     FacturaCliente,
		Retira:      FacturaRetira,
		Rut:         FacturaRut,
		Direccion:   FacturaDireccion,
		Email:       FacturaEmail,
		NombreLibro: FacturaNombreLibro,
		Precio:      FacturaPrecio,
		Cantidad:    FacturaCantidad,
		Total:       FacturaTotal}

	result, err := coll.InsertOne(context.TODO(), factura)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Documents matched:", result)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(factura)
	defer client.Disconnect(context.Background())
}
