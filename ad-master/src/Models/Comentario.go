package Models

import (
	"first-api/Config"

)
func CreateComentario(comentario *Comentario) (err error) { 
	if err = Config.DB.Create(comentario).Error; err != nil {
	  return err
	}
   return nil
}