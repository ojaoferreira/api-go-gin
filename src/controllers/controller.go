package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ojaoferreira/api-go-gin/src/database"
	"github.com/ojaoferreira/api-go-gin/src/models"
)

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)

	c.JSON(200, alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")

	c.JSON(200, gin.H{
		"API diz:": "E aí " + nome + ", tudo beleza",
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

func BuscaAlunoPorId(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func DeletaAlunoPorId(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno
	database.DB.First(&aluno, id)
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusNoContent, nil)
}

func AtualizaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorCPF(c *gin.Context) {
	cpf := c.Params.ByName("cpf")
	var aluno models.Aluno
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, aluno)
}
