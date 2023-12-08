package main

import (
	"encoding/json"
	"fmt"
	"gin-api-rest/controllers"
	"gin-api-rest/database"
	"gin-api-rest/models"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	return rotas
}

func TestRotaSaudacaoOK(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest(http.MethodGet, "/bob", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, `{"API diz":"E ai bob, tudo beleza?"}`, mustReadBody(resp.Body))
}

func TestRotaListaAlunosOK(t *testing.T) {
	database.ConectaComBancoDeDados()
	alunoMock := criaAluno()
	defer deletaAluno(alunoMock)
	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosOsAlunos)
	req, _ := http.NewRequest(http.MethodGet, "/alunos", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	var res []*models.Aluno
	json.NewDecoder(resp.Body).Decode(&res)
	assert.Equal(t, alunoMock.Nome, res[len(res)-1].Nome)
	assert.Equal(t, alunoMock.CPF, res[len(res)-1].CPF)
	assert.Equal(t, alunoMock.RG, res[len(res)-1].RG)
}

func TestRotaBuscaAlunoPorIDOK(t *testing.T) {
	database.ConectaComBancoDeDados()
	alunoMock := criaAluno()
	defer deletaAluno(alunoMock)
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/alunos/%d", alunoMock.ID), nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	var res *models.Aluno
	json.NewDecoder(resp.Body).Decode(&res)
	assert.Equal(t, alunoMock.Nome, res.Nome)
	assert.Equal(t, alunoMock.CPF, res.CPF)
	assert.Equal(t, alunoMock.RG, res.RG)
}

func TestRotaBuscaAlunoPorCPFOK(t *testing.T) {
	database.ConectaComBancoDeDados()
	alunoMock := criaAluno()
	defer deletaAluno(alunoMock)
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:cpf", controllers.BuscarAlunoPorCPF)
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/alunos/%s", alunoMock.CPF), nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	var res *models.Aluno
	json.NewDecoder(resp.Body).Decode(&res)
	assert.Equal(t, alunoMock.Nome, res.Nome)
	assert.Equal(t, alunoMock.CPF, res.CPF)
	assert.Equal(t, alunoMock.RG, res.RG)
}

func TestRotaDeletar(t *testing.T) {
	database.ConectaComBancoDeDados()
	alunoMock := criaAluno()
	defer deletaAluno(alunoMock)
	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.DeletaAlunoPorID)
	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/alunos/%d", alunoMock.ID), nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, `{"data":"aluno deletado com sucesos"}`, mustReadBody(resp.Body))
}

func criaAluno() *models.Aluno {
	if database.DB == nil {
		database.ConectaComBancoDeDados()
	}
	aluno := new(models.Aluno)
	aluno.Nome = "Nome aluno teste"
	aluno.CPF = "00000000001"
	aluno.RG = "000000000"
	database.DB.Create(&aluno)
	return aluno
}

func deletaAluno(aluno *models.Aluno) {
	if database.DB == nil {
		database.ConectaComBancoDeDados()
	}
	database.DB.Delete(aluno, aluno.ID)
}

func mustReadBody(in io.Reader) string {
	arr, _ := io.ReadAll(in)
	return string(arr)
}
