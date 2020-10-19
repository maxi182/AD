package Models

import (
	"first-api/Config"
	_ "github.com/go-sql-driver/mysql"
	//http://gorm.io/es_ES/docs/query.html
)

//GetUserByID ... Fetch only one propiedad by Id
func GetReclamoByID(reclamo *Reclamo, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).Find(reclamo).Error; err != nil {
		return err
	}
	return nil
}

//GetAllPropiedades Fetch all propiedad data
func GetAllReclamos(reclamo *[]Reclamo) (err error) {

	if err = Config.DB.Preload("Reclamos").Find(&reclamo).Error; err != nil {
		return err
	}
	return nil
}


//GetAllPropiedades Fetch all propiedad data
func GetAllReclamosByUser(reclamo *[]Reclamo, userId string) (err error) {

  if err = Config.DB.Model(&Reclamo{}).Preload("Comentarios").Select("*").Joins("inner join Comentarios on Comentarios.reclamo_id = Reclamos.id").Where("Reclamos.user_id = ?", userId).Find(&reclamo).Error; err != nil {
	   	return err
	  }
  return nil
}

func CreateReclamo(reclamo *Reclamo) (err error) { 


	if err = Config.DB.Omit("Comentarios", "Fotos","Unidad").Create(reclamo).Error; err != nil {
	  return err
	}
  if err =  Config.DB.Model(reclamo).Association("Comentarios").Replace(reclamo.Comentarios).Error; err != nil {
	   return err
	}
	if err =  Config.DB.Model(reclamo).Association("Unidad").Replace(reclamo.Unidad).Error; err != nil {
	   return err
	}



   return nil
}