package models

import (
	"time"

	"gorm.io/gorm"
)

type Aluno struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Nome      string         `json:"nome"`
	RG        string         `json:"rg"`
	CPF       string         `json:"cpf"`
	CreatedAt time.Time      `json:"createAt,omitempty"`
	UpdatedAt time.Time      `json:"updateAt,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"index"`
}

var Alunos []Aluno
