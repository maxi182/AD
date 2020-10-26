package Models

import (
	"first-api/Config"

)
func CreateNotification(notificacion *Notificacion) (err error) { 
	if err = Config.DB.Create(notificacion).Error; err != nil {
	  return err
	}
   return nil
}

//GetAllReclamosByUser Fetch all propiedad data
func GetAllNotificationsByUser(notificacion *[]Notificacion, userId string) (err error) {
	
	if err = Config.DB.Model(&notificacion).Preload("Usuario").Select("*").Where("Notificaciones.usuario_id = ?", userId).Find(&notificacion).Error; err != nil {
			 return err
		}
	return nil
  }
  