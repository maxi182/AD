package Models


type SharedArea struct {
  
	Id          uint  		    `gorm:"primary_key;auto_increment" json:"area_id"` 
	Descripcion string  	    `json:"descripcion"`
	Image		string			`json:"image"`
 
}

func (b *SharedArea) TableName() string {
	return "Shared"
}