package Models

import (
	"first-api/Config"

)

type Propiedad struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"propiedad_id"` 
	Direccion string    `json:"direccion"`
	Nombre    string    `json:"nombre"`
	Localidad string    `json:"localidad"`
	Provincia string    `json:"provincia"`
	Lat       float64   `json:"lat"`
	Lon       float64   `json:"lon"`
}

func (b *Propiedad) TableName() string {
	return "Propiedades"
}

func NewPropiedad(propiedad Propiedad) error {
	 rs:=Config.DB.Create(propiedad)
	 return rs.Error
}
