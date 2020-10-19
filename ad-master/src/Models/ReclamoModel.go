package Models

import (
	"first-api/Config"
	"time"

)

type Reclamo struct {
	ID        		uint      		`gorm:"primary_key;auto_increment" json:"reclamo_id"` 
	FechaCreacion 	time.Time    	`json:"fechaCreacion"`
	Estado    		uint    		`json:"estado"`
	Usuario			User 			`gorm:"foreignKey:id" json:"usuario"`  
	UsuarioId		uint 			`json:"usuario_id"`
	Comentarios  	[]Comentario   	`gorm:"foreignKey:comentario_id" json:"comentarios"`        //`gorm:"many2many:PropiedadUnidad" json:"unidades"`
	Unidad			Unidad			`gorm:"foreignKey:unidad_id" json:"unidad"`  
}

func (b *Reclamo) TableName() string {
	return "Reclamos"
}

func NewReclamo(reclamo Reclamo) error {
	 rs:=Config.DB.Create(reclamo)
	 return rs.Error
}
