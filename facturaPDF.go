package main

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func FacturarPdf() {
	// Leer el archivo HTML
	htmlBytes, err := ioutil.ReadFile("factura1515.html")
	if err != nil {
		log.Fatalf("Error al leer el archivo HTML: %v", err)
	}

	// Crear una nueva instancia de la biblioteca wkhtmltopdf
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatalf("Error al crear una nueva instancia de wkhtmltopdf: %v", err)
	}

	// Establecer el contenido HTML
	pdfg.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader(htmlBytes)))

	// Configurar las opciones del generador de PDF
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	// Generar el archivo PDF
	err = pdfg.Create()
	if err != nil {
		log.Fatalf("Error al generar el archivo PDF: %v", err)
	}

	// Guardar el archivo PDF generado
	err = pdfg.WriteFile("archivo.pdf")
	if err != nil {
		log.Fatalf("Error al guardar el archivo PDF: %v", err)
	}
}
