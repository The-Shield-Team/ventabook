package main

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func FacturarHtml(f Factura) {

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

	doc.Find("#tipo").SetHtml(f.Tipo)
	doc.Find("#numero").SetHtml(strconv.Itoa(f.Numero))
	doc.Find("#fecha").SetHtml(f.Fecha)
	doc.Find("#cliente").SetHtml(f.Cliente)
	doc.Find("#retira").SetHtml(f.Retira)
	doc.Find("#rut").SetHtml(f.Rut)
	doc.Find("#direccion").SetHtml(f.Direccion)
	doc.Find("#email").SetHtml(f.Email)
	doc.Find("#nombreLibro").SetHtml(f.NombreLibro)
	doc.Find("#total").SetHtml(strconv.Itoa(f.Total))

	htmlString, err = doc.Html()
	if err != nil {
		panic(err)
	}
	// Escribir la cadena de texto actualizada de vuelta al archivo HTML
	facturaSalida := "factura.html"

	err = ioutil.WriteFile(facturaSalida, []byte(htmlString), 0644)
	if err != nil {
		panic(err)
	}

}
