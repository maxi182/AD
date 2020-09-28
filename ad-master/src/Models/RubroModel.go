package Models


type Rubro struct {
 
	Id      uint   `gorm:"primary_key" json:"rubro_id"`
	Descripcion string  `json:"descripcion"`
 
}

func (b *Rubro) TableName() string {
	return "Rubros"
}
