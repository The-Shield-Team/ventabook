package main

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func FacturarHtml() string {

	// cambiar html
	htmlBytes, err := ioutil.ReadFile("factura.html")
	if err != nil {
		panic(err)
	}

	// Convertir los bytes a una cadena de texto
	htmlString := string(htmlBytes)

	// Seleccionar y actualizar el valor del elemento con id="mi-elemento"
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlString))
	if err != nil {
		panic(err)
	}
	numero := 1515
	doc.Find("#tipo").SetHtml("1111111111Nuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valorNuevo valor")
	doc.Find("#numero").SetHtml(strconv.Itoa(numero))

	htmlString, err = doc.Html()
	if err != nil {
		panic(err)
	}
	// Escribir la cadena de texto actualizada de vuelta al archivo HTML
	facturaSalida := "factura" + strconv.Itoa(numero) + ".html"

	err = ioutil.WriteFile(facturaSalida, []byte(htmlString), 0644)
	if err != nil {
		panic(err)
	}
	return facturaSalida
}
