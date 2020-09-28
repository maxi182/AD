package Models


type Rubro struct {
 
	Id      uint   `json:"rubro_id"`
	Descripcion string   `gorm:"not null" json:"descripcion"`
 
}

func (b *Rubro) TableName() string {
	return "Rubros"
}
