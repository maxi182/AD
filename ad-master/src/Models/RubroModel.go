package Models
import (
	"first-api/Config"

)


type Rubro struct {
 
	Id      uint   `gorm:"primary_key" json:"rubro_id"`
	Descripcion string  `json:"descripcion"`
 
}

func (b *Rubro) TableName() string {
	return "Rubros"
}

func NewRubro(rubro Rubro) error {
	rs:=Config.DB.Create(rubro)
	return rs.Error
}