package routes

import (
	"gin-api-rest/controllers"

	"github.com/gin-gonic/gin"
)

func HanldeRequest() {
	r := gin.Default()
	r.GET("alunos", controllers.ExibeTodosOsAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.GET("alunos/:id", controllers.BuscaAlunoPorID)
	r.PUT("alunos", controllers.EditaAluno)
	r.DELETE("alunos/:id", controllers.DeletaAlunoPorID)
	r.GET("alunos/cpf/:cpf", controllers.BuscarAlunoPorCPF)
	r.POST("alunos", controllers.CriaNovoAluno)
	r.Run()
}
