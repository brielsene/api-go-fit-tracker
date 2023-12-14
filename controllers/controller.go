package controllers

import (
	"api-go-fit-tracker/database"
	"api-go-fit-tracker/models"
	"fmt"
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

func RetornaTreinoPeloId(c *gin.Context) {
	id := c.Param("id")
	var treino models.Treino
	database.DB.First(&treino, id)

	if treino.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "Treino com id: " + id + ", não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, &treino)

}

func RetornaMsgHtml(c *gin.Context) {
	var treino models.Treino
	database.DB.First(&treino, 1)
	c.HTML(http.StatusOK, "index.html", &treino)
}

func DeletaTreino(c *gin.Context) {
	var treino models.Treino
	id := c.Param("id")
	database.DB.First(&treino, id)
	database.DB.Delete(&treino)
	c.JSON(http.StatusOK, gin.H{
		"TreinoID": id + ", excluido com sucesso!",
	})

}

func EditaTreino(c *gin.Context) {
	var treino models.Treino
	id := c.Param("id")
	database.DB.First(&treino, id)
	if err := c.ShouldBindJSON(&treino); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	database.DB.Model(&treino).UpdateColumns(treino)
	c.JSON(http.StatusOK, &treino)
}

func IndexHtml(c *gin.Context) {
	var treinos []models.Treino
	database.DB.Find(&treinos)
	fmt.Println(treinos)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"treinos": &treinos,
	})
}

func CreateHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "create-treino.html", nil)
}

func NewTreinoHTML(c *gin.Context) {
	var treino models.Treino
	nome := c.PostForm("nome")
	descricao := c.PostForm("descricao")
	diaDaSemana := c.PostForm("diadasemana")
	treino.Nome = nome
	treino.Descricao = descricao
	treino.DiaDaSemana = diaDaSemana
	database.DB.Create(&treino)
	c.Redirect(http.StatusSeeOther, "index")
}

func DeleteTreinoHTML(c *gin.Context) {
	fmt.Println("Entrando no delete controller")
	id := c.Param("id")
	var treino models.Treino
	database.DB.Delete(&treino, id)
	c.Redirect(http.StatusSeeOther, "/index")
}

func GetHtmlEdit(c *gin.Context) {
	var treino models.Treino
	id := c.Param("id")
	database.DB.First(&treino, id)
	c.HTML(http.StatusOK, "edit-treino.html", gin.H{
		"treino": &treino,
	})
}
