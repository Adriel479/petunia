package main

import (
	"gin-api-rest/database"
	"gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HanldeRequest()
}
