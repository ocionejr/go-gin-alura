package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ocionejr/go-gin-alura/database"
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

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}