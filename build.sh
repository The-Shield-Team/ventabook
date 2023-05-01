#!/bin/bash

# Instalar las dependencias de la aplicación
go get -d -v ./...

# Compilar la aplicación
go build