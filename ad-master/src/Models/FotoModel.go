package Models

import (
	"first-api/Config"
	"time"

)

type Foto struct {
	ID        		uint      `gorm:"primary_key;auto_increment" json:"foto_id"` 
	Fecha 			time.Time `json:"fecha"`
	Uri    			string    `json:"uri"`
}

func (b *Foto) TableName() string {
	return "Fotos"
}

func NewFoto(foto Foto) error {
	 rs:=Config.DB.Create(foto)
	 return rs.Error
}
