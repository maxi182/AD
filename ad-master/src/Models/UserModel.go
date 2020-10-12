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
	Nombre         string          `json:"nombre"`
	Apellido       string          `json:"apellido"`
	Email          string          `json:"email"`
	Image          string          `json:"image"`
	Fechanac       string          `json:"fechanac"`
	Dni            string          `json:"dni"`
	Rubros         []Rubro         `gorm:"many2many:RubroUsuario"  json:"rubros"`
	Propiedades    []Propiedad     `gorm:"many2many:PropiedadUsuario" json:"propiedades"`
	Date_created   string          `json:"date_created"`
	Password       string          `json:"password" validator:"min=4"`
	Is_active      bool            `json:"is_active"`
	Is_first_login bool            `json:"is_first_login" gorm:"type:boolean"`
}

func (b *User) TableName() string {
	return "Usuarios"
}

 