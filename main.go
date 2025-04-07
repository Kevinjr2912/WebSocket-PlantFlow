package main

import (
	"consumersData/infraestructure"
	"consumersData/infraestructure/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Establecemos las dependencias
	infraestructure.StartDependencies()

	// Registramos las rutas
	routes.RegisterRoutes()

	fmt.Println("Server running 8081")

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("Error al iniciarl el servidor")
	}

	
}