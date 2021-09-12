package main

import (
	"log"

	"github.com/neimv/curso-go/bd"
	"github.com/neimv/curso-go/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la base de datos")
		return
	}

	handlers.Manejadores()
}
