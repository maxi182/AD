package Models

type User struct {
	Id      uint   `gorm:"primary_key"`
	Usertype uint   `json:"usertype"`
	Nombre    string `json:"nombre"`
	Apellido string `json:"apellido"`
	Email   string `json:"email"`
	Fechanac string `json:"fechanac"`
	Address string `json:"address"`
	Dni string `json:"dni"`
	Date_created string `json:"date_created"`
	Password string `json:"password"`
	Password_expires string  `json:"password_expires,omitempty"`
	Is_active bool  `json:"is_active"`
}

func (b *User) TableName() string {
	return "Usuarios"
}
