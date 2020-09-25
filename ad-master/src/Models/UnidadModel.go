package Models




type Unidad struct {
	ID            uint     `gorm:"primary_key;auto_increment" json:"id"` 
//	Id            uint     `gorm:"primary_key"`
	Id_propiedad  uint     `json:"id_propiedad"`
	Piso          uint     `json:"piso"`
	Depto         string   `json:"dpto"`
}

func (b *Unidad) TableName() string {
	return "Unidades"
}
