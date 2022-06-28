package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ojaoferreira/api-go-gin/src/controllers"
)

func HandlerRequests() {
	r := gin.Default()

	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET(":nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)
	r.DELETE("/alunos/:id", controllers.DeletaAlunoPorId)
	r.PATCH("/alunos/:id", controllers.AtualizaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)

	r.Run()
}
