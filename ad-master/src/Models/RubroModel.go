package Models

import (
	"github.com/jinzhu/gorm"
)

type RubroUsuario struct {
	gorm.Model
	Id      uint   `gorm:"primary_key"`
	Id_rubro uint   `json:"id_rubro"`
	Id_usuario uint   `json:"id_usuario"`
}

func (b *RubroUsuario) TableName() string {
	return "RubroUsuario"
}
