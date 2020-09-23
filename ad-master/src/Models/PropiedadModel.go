package Models

type Propiedad struct {
	Id        uint      `gorm:"primary_key"`
	Direccion string    `json:"direccion"`
	Nombre    string    `json:"nombre"`
	Localidad string    `json:"localidad"`
	Provincia string    `json:"provincia"`
	Lat       float64   `json:"lat"`
	Lon       float64   `json:"lon"`
	Unidades  []Unidad  `json:"unidades"`

}

func (b *Propiedad) TableName() string {
	return "Propiedades"
}
