package Controllers
import (
   "first-api/Models"
   "fmt"
   "first-api/Config"
   "strconv"
   "time"
   "first-api/Utils"
   "net/http" //https://golang.org/pkg/net/http/
   "github.com/gin-gonic/gin"
   "github.com/biezhi/gorm-paginator/pagination"
)


//Get reclamos by user
func GetNotificationsByUser(c *gin.Context) {
	var notificacion []Models.Notificacion

	params, ok := c.Request.URL.Query()["userId"]
	if(!ok){
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : gin.H { 
			"status":  http.StatusBadRequest,
			"message": "Invalid param",
		}})
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err := Models.GetAllNotificationsByUser(&notificacion, string(params[0]))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : gin.H { 
			"status":  http.StatusNotFound,
			"message": err.Error(),
		}})
		c.AbortWithStatus(http.StatusNotFound)
	} else {
	 
		fmt.Println(c.Request.URL.Query())
		 page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		 limit, _ := strconv.Atoi(c.DefaultQuery("limit", "15"))
	
		  paginator := pagination.Paging(&pagination.Param{
			DB:      Config.DB.Model(&notificacion).Preload("Usuario").Select("*").Where("Notificaciones.usuario_id = ?", params).Find(&notificacion),
			Page:    page,
			Limit:   limit,
			OrderBy: []string{"Notificaciones.id"},
			ShowSQL: true,
		},  &notificacion)
 
		c.JSON(200, paginator)
	}
}


func saveNotificationComent(comentario Models.Comentario) {
	var now = time.Now().Unix()
	var notification Models.Notificacion
	notification.ReclamoId = comentario.ReclamoId
	notification.Type = 10 //Comentario creado
	notification.Date_created = Utils.ConvertTimestampToDate(int64(now))
	notification.UsuarioId = comentario.UsuarioId
	err := Models.CreateNotification(&notification)
	if err != nil {
		fmt.Println("notification_created_fail", notification.Id)
	}

}

func saveNotificationUpdateEstado(reclamo Models.Reclamo){
	var now = time.Now().Unix()
	var notification Models.Notificacion
	notification.ReclamoId = reclamo.ID
	notification.Type = reclamo.Estado //Reclamo cambio de estado lo guarda como type en notificacion
	notification.Date_created = Utils.ConvertTimestampToDate(int64(now))
	notification.UsuarioId = reclamo.UsuarioId
	err := Models.CreateNotification(&notification)
	if err != nil {
		fmt.Println("notification_estado_fail", notification.Id)
	}
}