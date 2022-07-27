package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ocionejr/go-gin-alura/controllers"
)

func HandleRequests(){
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.Run()
}