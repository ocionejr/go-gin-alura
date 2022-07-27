package main

import (
	"github.com/ocionejr/go-gin-alura/database"
	"github.com/ocionejr/go-gin-alura/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}