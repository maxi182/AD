package Models

import (
	"first-api/Config"
	_ "github.com/go-sql-driver/mysql"
	//http://gorm.io/es_ES/docs/query.html
)

//GetReclamosByID
func GetReclamoByID(reclamo *Reclamo, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).Find(reclamo).Error; err != nil {
		return err
	}
	return nil
}

//GetReclamosByPropAndReclamoID 
func GetReclamoByReclamoAndPropID(reclamo *Reclamo, id string, propId string) (err error) {
	if err = Config.DB.Where("id = ? AND propiedad_id= ? ", id, propId).Preload("Usuario").Find(reclamo).Error; err != nil {
		return err
	}
	return nil
}

//GetAllReclamos 
func GetAllReclamos(reclamo *[]Reclamo) (err error) {

	if err = Config.DB.Model(&Reclamo{}).Preload("Comentarios").Preload("Propiedad").Select("*").Joins("inner join Comentarios on Comentarios.reclamo_id = Reclamos.id").Find(&reclamo).Error; err != nil {
		return err
   }
return nil
}


//GetAllReclamosByUser Fetch all propiedad data
func GetAllReclamosByUser(reclamo *[]Reclamo, userId string) (err error) {
	
  if err = Config.DB.Model(&Reclamo{}).Preload("Comentarios").Preload("Usuario").Preload("Comentarios.Usuario").Preload("Comentarios.Fotos").Preload("Propiedad").Select("*").Joins("inner join Usuarios on Reclamos.usuario_id = Usuarios.id").Where("Usuarios.id = ?", userId).Find(&reclamo).Error; err != nil {
	   	return err
	  }
  return nil
}

//UpdateUser ... Update user
func UpdateReclamo(reclamo *Reclamo, id string) (err error) {
	Config.DB.Save(&reclamo)
	return nil
}

func UpdateEstadoReclamo(reclamo *Reclamo, reclamoId uint, estado uint, updated string) (err error) {

	if err = Config.DB.Model(reclamo).Where("ID = ?", reclamoId).Find(reclamo).Updates(map[string]interface{}{"estado": estado, "date_updated":updated}).Error; err != nil {
		return err
	}
	return nil
}

func UpdateRepairDateReclamo(reclamo *Reclamo, reclamoId uint, dateRepair string, updated string) (err error) {

	if err = Config.DB.Model(reclamo).Where("ID = ?", reclamoId).Find(reclamo).Updates(map[string]interface{}{"date_repair": dateRepair, "date_updated":updated}).Error; err != nil {
		return err
	}
	return nil
}

func GetAllReclamosByPropiedadEstado(reclamo *[]Reclamo, recStatus string, recId string, propId string) (err error) {
	
	if err = Config.DB.Model(&Reclamo{}).Preload("Comentarios").Preload("Usuario").Preload("Comentarios.Usuario").Preload("Propiedad").Select("*").Where("estado = ? AND propiedad_id = ? AND id != ?", recStatus, propId, recId).Find(&reclamo).Error; err != nil {
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