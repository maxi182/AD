package Models

import (
	"first-api/Config"
)

type Reclamo struct {
	ID        		uint      		`gorm:"primary_key;auto_increment" json:"reclamo_id"` 
	Date_created    string          `json:"date_created"`
	Date_updated    string          `json:"date_updated"`
	Estado    		uint    		`json:"estado"`
	UsuarioId		uint 			`json:"usuario_id"`
	Usuario			User 			`gorm:"foreignKey:UsuarioId" json:"creator"`  
	Comentarios  	[]Comentario   	`json:"comentarios"`        //`gorm:"many2many:PropiedadUnidad" json:"unidades"`
	UnidadId		uint			`json:"unidad_id"`
	PropiedadId     uint            `json:"propiedad_id"`
	RubroId         uint            `json:"rubro_id"`
	Propiedad 	    Propiedad       `json:"propiedad"`
	SharedAreaId    uint		    `json:"area_id"`
	SharedAreas     SharedArea      `gorm:"foreignKey:SharedAreaId" json:"shared"`  					//	`gorm:"many2many:SharedReclamo" json:"shared"`
}

func (b *Reclamo) TableName() string {
	return "Reclamos"
}

func NewReclamo(reclamo Reclamo) error {
	 rs:=Config.DB.Create(reclamo)
	 return rs.Error
}
