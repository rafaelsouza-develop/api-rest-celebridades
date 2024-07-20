package main

import (
	"api-rest-celebridades/database"
	"api-rest-celebridades/routes"
	"fmt"
)

func main() {
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando Servidor Rest com Go")
	routes.HandleRequests()
}
