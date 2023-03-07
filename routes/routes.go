package routes

import (
	"github.com/gin-gonic/gin"
	"rest-api/controller"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/user", controller.ShowAllUsers)
	r.POST("/user", controller.CreateUser)
	r.GET("/user/:code", controller.FindUserByCode)
	r.Run()
}
