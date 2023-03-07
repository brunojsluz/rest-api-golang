package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rest-api/database"
	"rest-api/models"
)

func ShowAllUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(200, users)
}

func FindUserByCode(c *gin.Context) {
	var user models.User
	var id = c.Params.ByName("code")
	database.DB.First(&user, id)
	c.JSON(200, user)
}

func CreateUser(c *gin.Context) {
	var user models.User
	var err = c.ShouldBindJSON(&user)

	if err != nil {
		log.Panic("Erro catastrofico!!!")
	}
	database.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}
