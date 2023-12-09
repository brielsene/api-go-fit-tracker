package main

import (
	"api-go-fit-tracker/database"
	"api-go-fit-tracker/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
