package database

import (
	"api-go-fit-tracker/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Fatal("Erro ao se conectar ao banco de dados: " + err.Error())
	}
	var treino models.Treino
	DB.AutoMigrate(&treino)
	//automigrate
}
