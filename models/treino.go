package models

import "gorm.io/gorm"

type Treino struct {
	gorm.Model
	Nome        string `json:"nome"`
	Descricao   string `json:"descricao"`
	DiaDaSemana string `json:"diaDaSemana"`
}
