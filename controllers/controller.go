package controllers

import (
	"api-go-fit-tracker/database"
	"api-go-fit-tracker/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegistraNovoTreino(c *gin.Context) {
	var treino models.Treino

	if err := c.ShouldBindJSON(&treino); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	database.DB.Create(&treino)
	c.JSON(http.StatusCreated, &treino)
}

func RetornaTodosOsTreinos(c *gin.Context) {
	var treinos []models.Treino
	database.DB.Find(&treinos)
	c.JSON(http.StatusOK, &treinos)
}
