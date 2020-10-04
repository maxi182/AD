package Models

import (
	 "time"
)
 
type User struct {
	Id             uint            `gorm:"primary_key"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time       `sql:"index"`
	Usertype       uint            `json:"usertype"`
	Nombre         string          `json:"nombre" valid:"length(6|20)"`
	Apellido       string          `json:"apellido"`
	Email          string          `json:"email"`
	Image          string          `json:"image"`
	Fechanac       string          `json:"fechanac"`
	Dni            string          `json:"dni"`
	Rubros         []Rubro         `gorm:"many2many:RubroUsuario"  json:"rubros"`
	Unidades       []Unidad        `gorm:"many2many:UnidadUsuario" json:"unidades"`
	Date_created   string          `json:"date_created"`
	Password       string          `json:"password" validator:"min=4"`
	Is_active      bool            `json:"is_active"`
	Is_first_login bool            `json:"is_first_login" gorm:"type:boolean"`
}

func (b *User) TableName() string {
	return "Usuarios"
}

 