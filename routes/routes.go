package routes

import (
	"api-go-fit-tracker/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.POST("/treinos", controllers.RegistraNovoTreino)
	r.GET("/treinos", controllers.RetornaTodosOsTreinos)
	r.Run(":8080")
}
