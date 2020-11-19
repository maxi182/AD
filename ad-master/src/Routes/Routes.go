package Routes

import (
	"first-api/Controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	user_route := r.Group("/user-api")
	{
		user_route.GET("users", Controllers.GetUsers)
		user_route.GET("users/:id", Controllers.GetUserByID)
		user_route.POST("users", Controllers.CreateUser)
		user_route.POST("login", Controllers.LoginUser)
		user_route.PUT("users/:id", Controllers.UpdateUser)
		user_route.DELETE("user/:id", Controllers.DeleteUser)
		user_route.GET("propiedades/all", Controllers.GetPropiedades)
		user_route.GET("propiedades", Controllers.GetPropiedadesByUser)
		user_route.GET("unidades", Controllers.GetUnidadesByUser)
	}

	accont_route := r.Group("/account")
	{
		accont_route.PUT("resetpassword", Controllers.UpdatePassword)
		accont_route.PUT("email", Controllers.SendEmail)
	}

	reclamos_route := r.Group("/reclamo")
	{
		reclamos_route.POST("reclamos", Controllers.CreateReclamo)
		reclamos_route.POST("comentario", Controllers.CreateComentario)
		reclamos_route.GET("reclamos/all", Controllers.GetReclamos)
		reclamos_route.GET("reclamos", Controllers.GetReclamosByUser)
		reclamos_route.PUT("reclamo/:id", Controllers.UpdateReclamo)
		reclamos_route.GET("reclamo", Controllers.GetReclamoByID)
		reclamos_route.PUT("estado", Controllers.UpdateEstadoReclamo)
		reclamos_route.PUT("estados", Controllers.UpdateEstadoReclamos)

		reclamos_route.POST("fotos", Controllers.UploadFotos)
	}
	notification_route := r.Group("/notification-api")
	{
		notification_route.GET("notifications", Controllers.GetNotificationsByUser)
	}

	rubros_route := r.Group("/rubros-api")
	{
		rubros_route.GET("rubros", Controllers.GetAllRubros)
	}

	shared_route := r.Group("/propiedades-api")
	{
		shared_route.GET("shared", Controllers.GetAllshared)
	}

	return r
}
