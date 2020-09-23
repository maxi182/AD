package Models
import (
	"github.com/jinzhu/gorm"
)
type User struct {
	gorm.Model
	Id         uint            `gorm:"primary_key"`
	Usertype   uint            `json:"usertype"`
	Nombre     string          `json:"nombre"`
	Apellido   string          `json:"apellido"`
	Email      string          `json:"email"`
	Image      string          `json:"image"`
	Fechanac   string          `json:"fechanac"`
	Dni        string          `json:"dni"`
	//Rubros     []RubroUsuario  `gorm:"ForeignKey:Id_usuario" json:"rubros"`
	Rubros []RubroUsuario        `gorm:"many2many:RubroUsuario;foreignKey:Id_usuario" json:"rubros"`
 
	Date_created string        `json:"date_created"`
	Password     string        `json:"password"`
	Is_active bool             `json:"is_active"`
	Is_first_login bool        `json:"is_first_login"`
}

func (b *User) TableName() string {
	return "Usuarios"
}
