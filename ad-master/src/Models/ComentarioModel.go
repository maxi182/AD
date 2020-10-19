package Models

import (
	"first-api/Config"
	"time"

)

type Comentario struct {
	ID        		uint      `gorm:"primary_key;auto_increment" json:"comentario_id"` 
	ReclamoId		uint	 `json:"reclamo_id"`
	Fecha 			time.Time `json:"fecha"`
	Texto    		string    `json:"texto"`
	Usuario			User	  `gorm:"foreignKey:id" json:"usuario"`  
	Fotos  			[]Foto   `gorm:"foreignKey:foto_id" json:"fotos"`        //`gorm:"many2many:PropiedadUnidad" json:"unidades"`
}

func (b *Comentario) TableName() string {
	return "Comentarios"
}

func NewComentario(comentario Comentario) error {
	 rs:=Config.DB.Create(comentario)
	 return rs.Error
}
