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

func CreateReclamo(c *gin.Context) {
	var rec Models.Reclamo
	c.BindJSON(&user)

	err := Models.CreateReclamo(&rec)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, rec)
	}
}

//GetPropiedadByID ... Get the property by id
func GetReclamoByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var rec Models.Reclamo
	err := Models.GetReclamoByID(&rec, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : gin.H { 
			"status":  http.StatusNotFound,
			"message": err.Error(),
		}})
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, rec)
	}
}

//Get all propiedades 
func GetReclamos(c *gin.Context) {
	var rec []Models.Reclamo
	var r Models.Reclamo
	err := Models.GetAllReclamos(&rec)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : gin.H { 
			"status":  http.StatusNotFound,
			"message": err.Error(),
		}})
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		log.Println("====== Bind By Query String ======")
		log.Println(r.ID)
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
		},  &rec)
 
		c.JSON(200, paginator)

	}
}

//Get all propiedades for the user.
/*func GetReclamosByUser(c *gin.Context) {
	var rec []Models.Reclamo
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
	err := Models.GetAllReclamosByUser(&rec, string(params[0]))
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
		},  &rec)
		c.JSON(http.StatusOK, paginator)
	}
}*/

//TENGO QUE HACER:

//GET BY EDIFICIO
//GET BY PENDIENTE ACEPTAR