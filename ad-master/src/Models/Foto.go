package Models

import (
	"first-api/Utils"
	"bytes"
	"first-api/Config"
	"github.com/jlaffaye/ftp"
	"time"
	"log"
	//"fmt"
)

//GetAllFotosByComentario etch all propiedad data
func GetAllFotosByComentario(foto *[]Foto, comentarioId string) (err error) {

	if err = Config.DB.Model(&Foto{}).Select("*").Joins("inner join Comentarios on Comentarios.Id = Fotos.id").Where("Fotos.ComentarioId = ?", comentarioId).Find(&foto).Error; err != nil {
			 return err
		}
	return nil
}

func UploadFotos(fotos *[]FotoUpload) (err error) {
	//c, err := ftp.Dial("190.55.95.131:3307", ftp.DialWithTimeout(5*time.Second))
	c, err := ftp.Dial("190.55.95.131:21", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	
	err = c.Login("reclamo@cerfoglio.com", "Uade12345")
	if err != nil {
		log.Fatal(err)
	}


	for i := 0; i < len(*fotos); i++ {
		//timestamp := time.Now().Unix()
		foto := (*fotos)[i]
		fotoBytes := Utils.DecodeBase64(foto.Data)

		data := bytes.NewBuffer(fotoBytes)
		//filename := fmt.Sprint(timestamp) + ".jpg"
		err = c.Stor(foto.Filename, data)
		if err != nil {
			panic(err)
		}
	}

	return nil
}