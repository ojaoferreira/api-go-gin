package main

import (
	"github.com/ojaoferreira/api-go-gin/src/database"
	"github.com/ojaoferreira/api-go-gin/src/models"
	"github.com/ojaoferreira/api-go-gin/src/routes"
)

func main() {
	database.ConectandoComDatabase()

	models.Alunos = []models.Aluno{
		{Nome: "Gui Lima", CPF: "00000000000", RG: "470000000"},
		{Nome: "Ana", CPF: "11111111111", RG: "480000000"},
	}

	routes.HandlerRequests()
}
