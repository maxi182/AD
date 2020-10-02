package Models


type Unidad struct {
	Id            uint    `gorm:"primary_key" json:"unidad_id"`
	Piso          uint     `json:"piso"`
	Depto         string   `json:"dpto"`
}



func (b *Unidad) TableName() string {
	return "Unidades"
}
