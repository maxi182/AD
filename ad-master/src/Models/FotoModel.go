package Models

import (
	"first-api/Config"
)

type Foto struct {
	Id       		uint      `gorm:"primary_key;auto_increment" json:"foto_id"` 
	ComentarioId	uint 	  `json:"comentario_id"` 
	Fecha 			string    `json:"fecha"`
	Uri    			string    `json:"uri"`
	Filename		string    `json:"filename"`
}

type FotoUpload struct {
	Data			string    `json:"data"`
	Filename		string    `json:"filename"`
	ComentarioId	uint 	  `json:"comentario_id"` 
}

func (b *Foto) TableName() string {
	return "Fotos"
}

func NewFoto(foto Foto) error {
	 rs:=Config.DB.Create(foto)
	 return rs.Error
}