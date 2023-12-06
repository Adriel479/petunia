package controllers

import (
	"gin-api-rest/database"
	"gin-api-rest/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ExibeTodosOsAlunos(ctx *gin.Context) {
	var alunos []*models.Aluno
	database.DB.Find(&alunos)
	ctx.JSON(http.StatusOK, alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Param("nome")
	c.JSON(http.StatusOK, gin.H{
		"API diz": "E ai " + nome + ", tudo beleza?",
	})
}

func BuscaAlunoPorID(c *gin.Context) {
	var (
		id, _ = strconv.Atoi(c.Param("id"))
		aluno models.Aluno
	)
	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := aluno.Validar(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Save(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := aluno.Validar(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusCreated, aluno)
}

func DeletaAlunoPorID(c *gin.Context) {
	var (
		id, _ = strconv.Atoi(c.Param("id"))
		aluno models.Aluno
	)
	database.DB.First(&aluno, id)
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{
		"data": "aluno deletado com sucesos",
	})
}

func BuscarAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	aluno.CPF = c.Param("cpf")
	database.DB.Where(&aluno).First(&aluno)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}
