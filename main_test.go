package main

import (
	"api-go-fit-tracker/controllers"
	"api-go-fit-tracker/database"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func CriaRotas() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas

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
