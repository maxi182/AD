package Controllers

import (
	"first-api/Models"
	"fmt"
	"net/http" //https://golang.org/pkg/net/http/
	"github.com/gin-gonic/gin"
 )

func UploadFotos(c *gin.Context) {
	var fotos []Models.Foto
	c.BindJSON(&fotos)

 	err := Models.UploadFotos(&fotos)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"error" : gin.H { 
			"status":  400,
			"message": err.Error(),
		}})
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
		})
		
		//fmt.Println("reclamo_creado", rec.ID)
	}
}