package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func FacturarPdf() {
	// Leer el archivo HTML
	htmlBytes, err := ioutil.ReadFile("factura.html")
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
	err = pdfg.WriteFile("factura.pdf")
	if err != nil {
		log.Fatalf("Error al guardar el archivo PDF: %v", err)
	}
}

func PDFHandler(w http.ResponseWriter, r *http.Request) {
	// Abrir el archivo PDF
	pdfFile, err := os.Open("factura.pdf")
	if err != nil {
		http.Error(w, "No se pudo leer el archivo PDF", http.StatusInternalServerError)
		return
	}
	defer pdfFile.Close()

	// Leer el contenido del archivo PDF en un []byte
	pdfBytes, err := ioutil.ReadAll(pdfFile)
	if err != nil {
		http.Error(w, "No se pudo leer el contenido del archivo PDF", http.StatusInternalServerError)
		return
	}

	// Establecer el encabezado de respuesta como "application/pdf"
	w.Header().Set("Content-Type", "application/pdf")

	// Escribir los datos del archivo PDF en la respuesta HTTP
	_, err = w.Write(pdfBytes)
	if err != nil {
		http.Error(w, "No se pudo escribir el contenido del archivo PDF en la respuesta", http.StatusInternalServerError)
		return
	}
}
