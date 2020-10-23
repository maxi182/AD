package Models

import (
	"first-api/Config"
	_ "github.com/go-sql-driver/mysql"
	//http://gorm.io/es_ES/docs/query.html
)

//GetReclamosByID ... Fetch only one propiedad by Id
func GetReclamoByID(reclamo *Reclamo, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).Find(reclamo).Error; err != nil {
		return err
	}
	return nil
}

//GetAllReclamos 
func GetAllReclamosByPropiedad(reclamo *[]Reclamo, propiedadId string) (err error) {

	if err = Config.DB.Model(&Reclamo{}).Preload("Comentarios").Preload("Propiedad").Select("*").Joins("inner join Comentarios on Comentarios.reclamo_id = Reclamos.id").Where("Reclamos.propiedad_id = ?", propiedadId).Find(&reclamo).Error; err != nil {
		return err
   }
return nil
}


//GetAllReclamosByUser Fetch all propiedad data
func GetAllReclamosByUser(reclamo *[]Reclamo, userId string) (err error) {
	
  if err = Config.DB.Model(&Reclamo{}).Preload("Comentarios").Preload("Usuario").Preload("Comentarios.Usuario").Preload("Propiedad").Select("*").Joins("inner join Usuarios on Reclamos.usuario_id = Usuarios.id").Where("Usuarios.id = ?", userId).Find(&reclamo).Error; err != nil {
	   	return err
	  }
  return nil
}

//UpdateUser ... Update user
func UpdateReclamo(reclamo *Reclamo, id string) (err error) {
	Config.DB.Save(&reclamo)
	return nil
}

func UpdateEstadoReclamo(reclamo *Reclamo, reclamoId uint, estado uint) (err error) {

	if err = Config.DB.Model(reclamo).Where("ID = ?", reclamoId).Find(reclamo).Updates(map[string]interface{}{"estado": estado}).Error; err != nil {
		return err
	}
	return nil
}

func CreateReclamo(reclamo *Reclamo) (err error) { 

	
	if err = Config.DB.Omit("Unidad","SharedAreas").Create(reclamo).Error; err != nil {
	  return err
	}
	if err =  Config.DB.Model(reclamo).Association("SharedAreas").Replace(reclamo.SharedAreas).Error; err != nil {
		return err
	 }
	if err =  Config.DB.Model(reclamo).Association("Comentarios").Replace(reclamo.Comentarios).Error; err != nil {
		return err
	 }
 
   return nil
}