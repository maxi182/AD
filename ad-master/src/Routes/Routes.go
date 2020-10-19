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
		//user_route.GET("propiedades/:id", Controllers.GetPropiedadByID)
		user_route.GET("propiedades/all", Controllers.GetPropiedades)
		user_route.GET("propiedades", Controllers.GetPropiedadesByUser)
		user_route.GET("unidades", Controllers.GetUnidadesByUser)
		user_route.POST("reclamos", Controllers.CreateReclamo)
	}

	accont_route := r.Group("/account")
	{
		accont_route.PUT("resetpassword", Controllers.UpdatePassword)
		accont_route.PUT("email", Controllers.SendEmail)
	}

	return r
}
