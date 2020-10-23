package Controllers
import (
   "first-api/Config"
   "first-api/Models"
   "fmt"
   "time"
   "strconv"
   "first-api/Utils"
   "github.com/biezhi/gorm-paginator/pagination"
   "log"
   "net/http" //https://golang.org/pkg/net/http/
   "github.com/gin-gonic/gin"
)

//GetReclamoByID ... Get the claim by id
func CreateReclamo(c *gin.Context) {
	var rec Models.Reclamo
	c.BindJSON(&rec)

	var now = time.Now().Unix()
	rec.Date_created = Utils.ConvertTimestampToDate(int64(now))

 	err := Models.CreateReclamo(&rec)
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
			"data" :  gin.H{
				"reclamo_id" : rec.ID,
				"user_id" : rec.UsuarioId,
				"estado":  rec.Estado,
			},
			"status":  http.StatusOK,
		})
		fmt.Println("reclamo_creado", rec.ID)
	}
}

//GetReclamoByID ... Get the claim by id
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

//Get reclamos by propiedad
func GetReclamos(c *gin.Context) {

	var rec []Models.Reclamo
	var r Models.Reclamo
	params, ok := c.Request.URL.Query()["propId"]
	if(!ok){
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : gin.H { 
			"status":  http.StatusBadRequest,
			"message": "Invalid param",
		}})
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err := Models.GetAllReclamosByPropiedad(&rec,string(params[0]))
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
			DB:   Config.DB.Model(&rec).Preload("Comentarios").Preload("Usuario").Preload("Comentarios.Usuario").Preload("Propiedad").Select("*").Joins("inner join Comentarios on Comentarios.reclamo_id = Reclamos.id").Where("Reclamos.propiedad_id = ?", params).Find(&rec),
			Page:    page,
			Limit:   limit,
			OrderBy: []string{"Reclamos.id"},
			ShowSQL: true,
		},  &rec)
 
		c.JSON(200, paginator)

	}
}

//Get reclamos by user
func GetReclamosByUser(c *gin.Context) {

	var rec []Models.Reclamo
	var r Models.Reclamo
	params, ok := c.Request.URL.Query()["userId"]
	if(!ok){
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : gin.H { 
			"status":  http.StatusBadRequest,
			"message": "Invalid param",
		}})
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err := Models.GetAllReclamosByUser(&rec,string(params[0]))
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
			DB:      Config.DB.Model(&rec).Preload("Comentarios").Preload("Usuario").Preload("Comentarios.Usuario").Preload("Propiedad").Select("*").Joins("inner join Usuarios on Reclamos.usuario_id = Usuarios.id").Where("Usuarios.id = ?", params).Find(&rec),
			Page:    page,
			Limit:   limit,
			OrderBy: []string{"Reclamos.id"},
			ShowSQL: true,
		},  &rec)
 
		c.JSON(200, paginator)

	}
}

func UpdateReclamo(c *gin.Context) {
	var rec Models.Reclamo
	id := c.Params.ByName("id")
	fmt.Println("id", id)
	err := Models.GetReclamoByID(&rec, id)
	if err != nil {
		c.JSON(http.StatusNotFound,  gin.H{
			"error" : gin.H { 
			"status":  http.StatusNotFound,
			"message": "Not Found",
		}})
		return
	} else {
	c.BindJSON(&rec)
	
	err = Models.UpdateReclamo(&rec, id)
	if err != nil {
	//	c.AbortWithStatus(http.StatusNotFound)
		c.JSON(http.StatusNotFound,  gin.H{
			"data":gin.H { 
			"error" : gin.H { 
			"status":  http.StatusBadRequest,
			"message": "Can´t update reclamo",
		}}})
	} else {
		 c.JSON(http.StatusOK, rec)
		return
	}
}
}

func UpdateEstadoReclamo(c *gin.Context) {
	var rec Models.Reclamo
	
	if err := c.ShouldBindJSON(&rec); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

	} else {
 
	err := Models.UpdateEstadoReclamo(&rec, rec.ID, rec.Estado)
	   if err != nil {
		c.JSON(http.StatusNotFound,  gin.H{
			"error" :   gin.H { 
			"status":   http.StatusNotFound,
			"message": "Not Found",
		}})
		return
	   } else {
		c.JSON(http.StatusOK, gin.H{"status": true})
	   }
	  }
	}
