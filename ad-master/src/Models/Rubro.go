
package Models

import (
	"first-api/Config"

	_ "github.com/go-sql-driver/mysql"
	//http://gorm.io/es_ES/docs/query.html
)
 
func GetAllRubros(rubro *[]Rubro) (err error) {
			
	if err = Config.DB.Find(&rubro).Error; err != nil {
		return err
	}
	return nil
}


//CreateUser ... Insert New data
func CreateRubro(rubro *Rubro) (err error) {
	if err = Config.DB.Create(rubro).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id    user *[]User
func GetRubroByUserID(rubro*[]Rubro, id string) (err error) {
	if err = Config.DB.Where("id_usuario = ?", id).Find(rubro).Error; err != nil {
		return err
	}
	return nil
}

