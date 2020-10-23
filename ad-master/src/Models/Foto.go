package Models

import (
	"first-api/Config"

)

//GetAllFotosByComentario etch all propiedad data
func GetAllFotosByComentario(foto *[]Foto, comentarioId string) (err error) {

	if err = Config.DB.Model(&Foto{}).Select("*").Joins("inner join Comentarios on Comentarios.Id = Fotos.id").Where("Fotos.ComentarioId = ?", comentarioId).Find(&foto).Error; err != nil {
			 return err
		}
	return nil
}