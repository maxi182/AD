package Models

import (
	"bytes"
	"first-api/Config"
	"first-api/Utils"
	"time"
	"strings"
	"github.com/jlaffaye/ftp"
)

//GetAllFotosByComentario etch all propiedad data
func GetAllFotosByComentario(foto *[]Foto, comentarioId string) (err error) {

	if err = Config.DB.Model(&Foto{}).Select("*").Joins("inner join Comentarios on Comentarios.Id = Fotos.id").Where("Fotos.ComentarioId = ?", comentarioId).Find(&foto).Error; err != nil {
		return err
	}
	return nil
}

func UploadFoto(foto *FotoUpload) (err error) {
	c, err := ftp.Dial("190.55.95.131:21", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		return err
	}

	err = c.Login("reclamo@cerfoglio.com", "Uade12345")
	if err != nil {
		return err
	}

	
	fotoBytes := Utils.DecodeBase64(strings.ReplaceAll(foto.Data, "data:image/jpeg;base64,", ""))
	data := bytes.NewBuffer(fotoBytes)
	err = c.Stor(foto.Filename, data)

	if err != nil {
		return err
	}

	return nil
}