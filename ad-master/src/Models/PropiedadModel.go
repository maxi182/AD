package Models

import (
	"first-api/Config"

)

type Propiedad struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"propiedad_id"` 
//	Id        uint      `gorm:"primary_key"`
	Direccion string    `json:"direccion"`
	Nombre    string    `json:"nombre"`
	Localidad string    `json:"localidad"`
	Provincia string    `json:"provincia"`
	Lat       float64   `json:"lat"`
	Lon       float64   `json:"lon"`
	Unidades  []Unidad   `gorm:"foreignKey:propiedad_id" json:"unidades"`        //`gorm:"many2many:PropiedadUnidad" json:"unidades"`
}

func (b *Propiedad) TableName() string {
	return "Propiedades"
}

func NewPropiedad(propiedad Propiedad) error {
	 rs:=Config.DB.Create(propiedad)
	 return rs.Error
}
