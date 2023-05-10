package main

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TipoFactura(w http.ResponseWriter, f string) (string, error) {
	FacturaTipo := strings.ToLower(f)
	if FacturaTipo == "boleta" || FacturaTipo == "factura" {
		return strings.Title(FacturaTipo), nil
	}
	return "", errors.New("el tipo de documento no es boleta o factura")

}

func LastDocument(w http.ResponseWriter, r *http.Request, tipo string) (int, error) {
	// Crear conexión a la base de datos de MongoDB
	client := connection(mongoInfo)
	coll := client.Database("facturacion").Collection("documentos")

	// Seleccionar la colección de documentos

	// Obtener el último número de boleta
	var lastDocument Factura
	options := options.FindOne().SetSort(bson.D{{Key: "numero", Value: -1}})
	filter := bson.D{{Key: "tipo", Value: tipo}}
	err := coll.FindOne(context.TODO(), filter, options).Decode(&lastDocument)
	if err != nil {
		return 0, err

	}
	defer client.Disconnect(context.Background())
	return lastDocument.Numero, nil

}
