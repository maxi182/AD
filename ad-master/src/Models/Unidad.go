package Models

import (
	"first-api/Config"
	_ "github.com/go-sql-driver/mysql"
	//http://gorm.io/es_ES/docs/query.html
)


//GetAllPropiedades Fetch all propiedad data
func GetAllUnidades(unidad *[]Unidad) (err error) {

	if err = Config.DB.Find(&unidad).Error; err != nil {
		return err
	}
	return nil
}


//GetUnidades by user
func GetAllUnidadesByUser(unidad *[]Unidad, userId string) (err error) {

	if err = Config.DB.Model(&Unidad{}).Preload("Propiedad").Select("*").Joins("inner join UnidadUsuario on Unidades.id = UnidadUsuario.unidad_id inner join Propiedades on Propiedades.id=Unidades.propiedad_id").Where("UnidadUsuario.user_id = ?", userId).Find(&unidad).Error; err != nil {
			 return err
		}
	return nil
  }
