package Models


type Unidad struct {
	Id            uint     `json:"id"` 
	Id_propiedad  uint     `json:"id_propiedad"`
	Piso          uint     `json:"piso"`
	Depto         string   `json:"dpto"`
}

func (b *Unidad) TableName() string {
	return "Unidades"
}
