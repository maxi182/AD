  
package Controllers

import (
	"first-api/Models"
	"net/http" //https://golang.org/pkg/net/http/
	"github.com/gin-gonic/gin"
 )

type fotoStatus struct {
    Filename string
	UploadFailed bool
	Error error
}

func UploadFotos(c *gin.Context) {
	var fotos []Models.FotoUpload
	c.BindJSON(&fotos)


	var statusFotos []fotoStatus
	hayError := false

	for i := 0; i < len(fotos); i++ {
		foto := (fotos)[i]

		err := Models.UploadFoto(&foto)

		statusError := false

		if err != nil{
			statusError = true
			hayError = true
		}

		status := fotoStatus{Filename: foto.Filename, Error: err, UploadFailed: statusError}
		statusFotos = append(statusFotos, status)
	}

	if hayError {
		c.JSON(http.StatusNotFound, gin.H{
			"fotos" : statusFotos,
			"error" : gin.H { 
			"status":  400,
			"message": "La subida de fotos contiene errores",
		}})
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"fotos" : statusFotos,
			"status":  http.StatusOK,
		})
	}
}