package Models


import (
	"github.com/jinzhu/gorm"
)
 
type Rubro struct {
	gorm.Model
	//Id      uint   `gorm:"primary_key"`
	Descripcion string   `json:"descripcion"`
//	Id_usuario uint   `json:"id_usuario"`
}

func (b *Rubro) TableName() string {
	return "Rubros"
}
