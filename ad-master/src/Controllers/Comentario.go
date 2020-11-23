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

type comStatus struct {
	ComentarioID uint
	Failed    bool
	Error     error
}

func CreateComentarios(c *gin.Context) {
	var comentarios []Models.Comentario
	
	var now = time.Now().Unix()

	var statusComentarios []comStatus
	hayError := false

	if err := c.ShouldBindJSON(&comentarios); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i := 0; i < len(comentarios); i++ {

		comentario := (comentarios)[i]
		comentario.Date_created = Utils.ConvertTimestampToDate(int64(now))

		err := Models.CreateComentario(&comentario)

		statusError := false

		if err != nil {
			statusError = true
			hayError = true
		}

		status := comStatus{ComentarioID: comentario.Id, Error: err, Failed: statusError}
		statusComentarios = append(statusComentarios, status)

		
	}

	if hayError {
		c.JSON(http.StatusNotFound, gin.H{
			"reclamos": statusComentarios,
			"error": gin.H{
				"status":  400,
				"message": "La creación de comentarios tiene errores",
			}})
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"reclamos": statusComentarios,
			"status": http.StatusOK,
		})
	}
}
