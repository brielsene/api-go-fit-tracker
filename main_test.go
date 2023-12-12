package main

import (
	"api-go-fit-tracker/controllers"
	"api-go-fit-tracker/database"
	"api-go-fit-tracker/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func CriaRotas() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas

}

func CriaTreinoMock() {
	treino := models.Treino{Nome: "Treino de perna", Descricao: "Muita intensidade", DiaDaSemana: "Terça"}
	database.DB.Create(&treino)
	ID = int(treino.ID)
}

func DeletaTreinoMock() {
	var treino models.Treino
	database.DB.Delete(&treino, ID)
}

func TestCodeRequest(t *testing.T) {
	database.ConectaComBancoDeDados()
	r := CriaRotas()
	r.GET("/treinos", controllers.RetornaTodosOsTreinos)
	req, _ := http.NewRequest("GET", "/treinos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestDeleteTreino(t *testing.T) {
	database.ConectaComBancoDeDados()
	r := CriaRotas()
	r.DELETE("/treinos/delete/:id", controllers.DeletaTreino)
	CriaTreinoMock()
	path := "/treinos/delete/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", path, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var treino models.Treino
	json.Unmarshal(resposta.Body.Bytes(), &treino)
	var id uint
	id = 0
	assert.Equal(t, id, treino.ID)
}
