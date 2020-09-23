
package Models

import (
	"first-api/Config"

	_ "github.com/go-sql-driver/mysql"
	//http://gorm.io/es_ES/docs/query.html
)
//CreateUser ... Insert New data
func CreateRubroUsuario(rubroUsuario *RubroUsuario) (err error) {
	if err = Config.DB.Create(rubroUsuario).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id    user *[]User
func GetRubroByUserID(rubroUsuario *[]RubroUsuario, id string) (err error) {
	if err = Config.DB.Where("id_usuario = ?", id).Find(rubroUsuario).Error; err != nil {
		return err
	}
	return nil
}

