package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ocionejr/go-gin-alura/models"
)

func ExibeTodosAlunos(c *gin.Context){
	c.JSON(200, models.Alunos)
}

func Saudacao(c *gin.Context){
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"Api diz:": "E ai " + nome + ", tudo beleza?",
	})
}