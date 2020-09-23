package Models

type Unidad struct {
	Id            uint     `gorm:"primary_key"`
	Id_propiedad  uint     `json:"id_propiedad"`
	Piso          uint     `json:"piso"`
	Depto         string   `json:"dpto"`
	
}

func (b *Unidad) TableName() string {
	return "Unidades"
}
