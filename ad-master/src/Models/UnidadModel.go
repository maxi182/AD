package Models


type Unidad struct {  
	Id            uint       `gorm:"primary_key" json:"unidad_id"`
	Piso          uint       `json:"piso"`
	Depto         string     `json:"dpto"`
	PropiedadId   uint       `json:"propiedad_id"`
	Propiedad 	  Propiedad  `json:"propiedad"`
}



func (b *Unidad) TableName() string {
	return "Unidades"
}
