package routes

import (
	"api-go-fit-tracker/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/styles", "./styles/")
	r.POST("/treinos", controllers.RegistraNovoTreino)
	r.GET("/treinos", controllers.RetornaTodosOsTreinos)
	r.GET("/treinos/:id", controllers.RetornaTreinoPeloId)
	r.DELETE("/treinos/delete/:id", controllers.DeletaTreino)
	r.GET("/index", controllers.IndexHtml)
	r.PATCH("/treinos/update/:id", controllers.EditaTreino)
	r.Run(":8080")
}
