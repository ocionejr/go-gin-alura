package main

import (
	"github.com/ocionejr/go-gin-alura/database"
	"github.com/ocionejr/go-gin-alura/models"
	"github.com/ocionejr/go-gin-alura/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = [] models.Aluno{
		{Nome: "Ocione", CPF: "312313123", RG: "312331323"},
		{Nome: "Leo", CPF: "312313123", RG: "312331323"},
	}
	routes.HandleRequests()
}