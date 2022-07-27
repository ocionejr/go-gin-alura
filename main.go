package main

import (
	"github.com/ocionejr/go-gin-alura/models"
	"github.com/ocionejr/go-gin-alura/routes"
)

func main() {
	models.Alunos = [] models.Aluno{
		{Nome: "Ocione", CPF: "312313123", RG: "312331323"},
		{Nome: "Leo", CPF: "312313123", RG: "312331323"},
	}
	routes.HandleRequests()
}