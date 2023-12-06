package database

import (
	"gin-api-rest/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	conexao := "host=localhost user=root password=root dbname=root port=55432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(conexao), &gorm.Config{})
	if err != nil {
		log.Println("[FATAL] erro ao abri conexão com o banco de dados:", err)
	}
	DB.AutoMigrate(&models.Aluno{})
}
