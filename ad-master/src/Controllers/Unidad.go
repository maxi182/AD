package Controllers
import (
   "first-api/Config"
   "first-api/Models"
   "fmt"
   "strconv"
   "github.com/biezhi/gorm-paginator/pagination"
   "log"
   "net/http" //https://golang.org/pkg/net/http/
   "github.com/gin-gonic/gin"
)


//Get all propiedades 
func GetUnidades(c *gin.Context) {
	var unidad []Models.Unidad
	var u Models.Unidad
	err := Models.GetAllUnidades(&unidad)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : gin.H { 
			"status":  http.StatusNotFound,
			"message": err.Error(),
		}})
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		log.Println("====== Bind By Query String ======")
		log.Println(u.Piso)
	 
		fmt.Println(c.Request.URL.Query())
		 page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		 limit, _ := strconv.Atoi(c.DefaultQuery("limit", "3"))
	
		  paginator := pagination.Paging(&pagination.Param{
			DB:      Config.DB.Preload("Propiedades"),
			Page:    page,
			Limit:   limit,
			OrderBy: []string{"id"},
			ShowSQL: true,
		},  &unidad)
 
		c.JSON(200, paginator)

	}
}


//Get all propiedades for the user.
func GetUnidadesByUser(c *gin.Context) {
	var unidad []Models.Unidad
	params, ok := c.Request.URL.Query()["userId"]
	
	if(!ok){
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : gin.H { 
			"status":  http.StatusBadRequest,
			"message": "Invalid param",
		}})
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err := Models.GetAllUnidadesByUser(&unidad, string(params[0]))
	if (err != nil || !ok ) {
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
			DB:      Config.DB.Model(&unidad).Preload("Propiedad").Preload("Propiedad.SharedAreas").Select("*").Joins("inner join UnidadUsuario on Unidades.id = UnidadUsuario.unidad_id inner join Propiedades on Propiedades.id=Unidades.propiedad_id").Where("UnidadUsuario.user_id = ?",params).Find(&unidad),
			Page:    page,
			Limit:   limit,
			OrderBy: []string{"Unidades.id"},
			ShowSQL: true,
		},  &unidad)
		c.JSON(http.StatusOK, paginator)
	}
}
