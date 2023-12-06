package main

import (
	"fmt"
	"go-rest-api/database"
	"go-rest-api/models"
	"go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{
			ID:       1,
			Nome:     "Nome 1",
			Historia: "História 1",
		},
		{
			ID:       2,
			Nome:     "Nome 2",
			Historia: "Historia 2",
		},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
