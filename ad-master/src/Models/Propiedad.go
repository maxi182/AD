package Models

import (
	"first-api/Config"
	_ "github.com/go-sql-driver/mysql"
	//http://gorm.io/es_ES/docs/query.html
)

//GetUserByID ... Fetch only one propiedad by Id
func GetPropiedadByID(propiedad *Propiedad, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).Find(propiedad).Error; err != nil {
		return err
	}
	return nil
}

//GetAllPropiedades Fetch all propiedad data
func GetAllPropiedades(propiedad *[]Propiedad) (err error) {

	if err = Config.DB.Preload("Unidades").Find(&propiedad).Error; err != nil {
		return err
	}
	return nil
}


//GetAllPropiedades Fetch all propiedad data
func GetAllPropiedadesByUser(propiedad *[]Propiedad, userId string) (err error) {

  if err = Config.DB.Model(&Propiedad{}).Select("*").Joins("inner join PropiedadUsuario on PropiedadUsuario.propiedad_id = Propiedades.id").Where("PropiedadUsuario.user_id = ?", userId).Find(&propiedad).Error; err != nil {
	   	return err
	  }
  return nil
}