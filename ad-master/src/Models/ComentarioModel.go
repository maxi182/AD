package Models

import (
	"first-api/Config"
)

type Comentario struct {
	Id        		uint      `gorm:"primary_key;auto_increment" json:"comentario_id"` 
	ReclamoId		uint	  `json:"reclamo_id"`
	Date_created    string    `json:"date_created"`
	Texto    		string    `json:"texto"`
	UsuarioId       uint      `json:"usuario_id"`
	Usuario			User	  `gorm:"foreignKey:UsuarioId" json:"usuario"`  
	Fotos  			[]Foto    `gorm:"foreignKey:ComentarioId" json:"fotos"`        //`gorm:"many2many:PropiedadUnidad" json:"unidades"`
}

func (b *Comentario) TableName() string {
	return "Comentarios"
}

func NewComentario(comentario Comentario) error {
	 rs:=Config.DB.Create(comentario)
	 return rs.Error
}
