package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"rest-api/models"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	connectionString := "host=postgres user=root password=root dbname=root port=5432"
	DB, err = gorm.Open(postgres.Open(connectionString))

	if err != nil {
		log.Panic("Erro de conexão com o banco")
	}
	DB.AutoMigrate(&models.User{})
}
