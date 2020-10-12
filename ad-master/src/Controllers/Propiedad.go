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
//GetPropiedadByID ... Get the property by id
func GetPropiedadByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var prop Models.Propiedad
	err := Models.GetPropiedadByID(&prop, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : gin.H { 
			"status":  http.StatusNotFound,
			"message": err.Error(),
		}})
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, prop)
	}
}

//Get all propiedades 
func GetPropiedades(c *gin.Context) {
	var prop []Models.Propiedad
	var p Models.Propiedad
	err := Models.GetAllPropiedades(&prop)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : gin.H { 
			"status":  http.StatusNotFound,
			"message": err.Error(),
		}})
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		log.Println("====== Bind By Query String ======")
		log.Println(p.Nombre)
		//var rubro []Models.RubroUsuario
	 
		fmt.Println(c.Request.URL.Query())
		 page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		 limit, _ := strconv.Atoi(c.DefaultQuery("limit", "15"))
	
		  paginator := pagination.Paging(&pagination.Param{
			DB:      Config.DB.Preload("Unidades"),
			Page:    page,
			Limit:   limit,
			OrderBy: []string{"id"},
			ShowSQL: true,
		},  &prop)
 
		c.JSON(200, paginator)

	}
}

//Get all propiedades for the user.
func GetPropiedadesByUser(c *gin.Context) {
	var prop []Models.Propiedad
	params, ok := c.Request.URL.Query()["userId"]
	fmt.Println("isok",ok)
	if(!ok){
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : gin.H { 
			"status":  http.StatusBadRequest,
			"message": "Invalid param",
		}})
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err := Models.GetAllPropiedadesByUser(&prop, string(params[0]))
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
			DB:       Config.DB.Model(&prop).Preload("Unidades").Select("*").Joins("inner join PropiedadUsuario on PropiedadUsuario.propiedad_id = Propiedades.id").Where("PropiedadUsuario.user_id = ?", params).Find(&prop),
			Page:    page,
			Limit:   limit,
			OrderBy: []string{"id"},
			ShowSQL: true,
		},  &prop)
		c.JSON(http.StatusOK, paginator)
	}
}