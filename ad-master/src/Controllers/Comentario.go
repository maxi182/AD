package Controllers
import (
   "first-api/Models"
   "fmt"
   "time"
   "first-api/Utils"
   "net/http" //https://golang.org/pkg/net/http/
   "github.com/gin-gonic/gin"
)


func CreateComentario(c *gin.Context) {
	var coment Models.Comentario
	
	c.BindJSON(&coment)
	var now = time.Now().Unix()
 	coment.Date_created = Utils.ConvertTimestampToDate(int64(now))

	err := Models.CreateComentario(&coment)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : gin.H { 
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		}})
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK,gin.H{
			"data" : coment,
			"status":  http.StatusCreated,
		})
		saveNotificationComent(coment)
		fmt.Println("comentario_creado", coment.Id)
	}
}
