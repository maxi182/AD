package Controllers

import (
	"first-api/Config"
	"first-api/Models"
	"first-api/Utils"
	"fmt"
	"log"
	"net/http" //https://golang.org/pkg/net/http/
	"strconv"
	"time"

	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
			"error": gin.H{
				"status":  400,
				"message": err.Error(),
			}})
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"reclamo_id": rec.ID,
				"comentario": rec.Comentarios,
				"user_id":    rec.UsuarioId,
				"estado":     rec.Estado,
			},
			"status": http.StatusOK,
		})
		saveNotificationUpdateEstado(rec)
		fmt.Println("reclamo_creado", rec.ID)
	}
}

//GetReclamoByID ... Get the claim by id
func GetReclamoByID(c *gin.Context) {

	var rec Models.Reclamo
	propId := c.Query("propId")
	reclamoId := c.Query("recId")

	err := Models.GetReclamoByReclamoAndPropID(&rec, reclamoId, propId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": gin.H{
				"status":  http.StatusNotFound,
				"message": err.Error(),
			}})
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":   rec,
			"status": http.StatusOK,
		})
	}
}

//Get reclamos by propiedad
func GetReclamos(c *gin.Context) {

	var rec []Models.Reclamo

	err := Models.GetAllReclamos(&rec)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": gin.H{
				"status":  http.StatusNotFound,
				"message": err.Error(),
			}})
		c.AbortWithStatus(http.StatusNotFound)
	} else {

		fmt.Println(c.Request.URL.Query())
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "15"))

		paginator := pagination.Paging(&pagination.Param{
			DB:      Config.DB.Model(&rec).Preload("Comentarios").Preload("Usuario").Preload("Comentarios.Usuario").Preload("Propiedad").Select("*").Joins("inner join Comentarios on Comentarios.reclamo_id = Reclamos.id").Find(&rec),
			Page:    page,
			Limit:   limit,
			OrderBy: []string{"Reclamos.id"},
			ShowSQL: true,
		}, &rec)

		c.JSON(200, paginator)

	}
}

//Get reclamos by user
func GetReclamosByUser(c *gin.Context) {

	var rec []Models.Reclamo
	var r Models.Reclamo
	userId := c.Query("userId")
	propId := c.Query("propId")
	//	groupByprop := c.Query("unidadId")

	err := Models.GetAllReclamosByUser(&rec, userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": gin.H{
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
			DB:      getQueryReclamo(&rec, userId, propId),
			Page:    page,
			Limit:   limit,
			OrderBy: []string{"Reclamos.id"},
			ShowSQL: true,
		}, &rec)

		c.JSON(200, paginator)

	}
}

func GetReclamosByPropiedadEstado(c *gin.Context) {

	var rec []Models.Reclamo
	//var r Models.Reclamo
	recStatus := c.Query("recStatus")
	recId := c.Query("recId")
	propId := c.Query("propId")


	err := Models.GetAllReclamosByPropiedadEstado(&rec, recStatus, recId, propId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": gin.H{
				"status":  http.StatusNotFound,
				"message": err.Error(),
			}})
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":   rec,
			"status": http.StatusOK,
		})
	}
}

func getQueryReclamo(reclamo *[]Models.Reclamo, userId string, propId string) *gorm.DB {

	fmt.Println("usuario", userId)
	query := Config.DB.Model(&reclamo).Preload("Comentarios").Preload("Usuario").Preload("Comentarios.Usuario").Preload("Propiedad").Select("*").Joins("inner join Usuarios on Reclamos.usuario_id = Usuarios.id").Find(&reclamo)

	if len(propId) > 0 && userId != "0" {
		return query.Where("Usuarios.id = ? AND Reclamos.propiedad_id = ?", userId, propId)
	} else {

		return query.Where("Reclamos.propiedad_id = ?", propId)
	}
}

func UpdateReclamo(c *gin.Context) {
	var rec Models.Reclamo
	id := c.Params.ByName("id")
	fmt.Println("id", id)
	err := Models.GetReclamoByID(&rec, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": gin.H{
				"status":  http.StatusNotFound,
				"message": "Not Found",
			}})
		return
	} else {
		c.BindJSON(&rec)

		err = Models.UpdateReclamo(&rec, id)
		if err != nil {
			//	c.AbortWithStatus(http.StatusNotFound)
			c.JSON(http.StatusNotFound, gin.H{
				"data": gin.H{
					"error": gin.H{
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

	var now = time.Now().Unix()
	var updated = Utils.ConvertTimestampToDate(int64(now))

	if err := c.ShouldBindJSON(&rec); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	} else {

		err := Models.UpdateEstadoReclamo(&rec, rec.ID, rec.Estado, updated)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": gin.H{
					"status":  http.StatusNotFound,
					"message": "Not Found",
				}})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"status": true})
			saveNotificationUpdateEstado(rec)
		}
	}
}

type recStatus struct {
	ReclamoID uint
	Failed    bool
	Error     error
}

func UpdateEstadoReclamos(c *gin.Context) {
	var rec []Models.Reclamo

	var now = time.Now().Unix()
	var updated = Utils.ConvertTimestampToDate(int64(now))

	var statusReclamos []recStatus
	hayError := false

	if err := c.ShouldBindJSON(&rec); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(len(rec))
	for i := 0; i < len(rec); i++ {

		reclamo := (rec)[i]
		fmt.Println(rec)

		err := Models.UpdateEstadoReclamo(&reclamo, reclamo.ID, reclamo.Estado, updated)

			statusError := false

			if err != nil {
				statusError = true
				hayError = true
			}else{
				saveNotificationUpdateEstado(reclamo)
			}

			status := recStatus{ReclamoID: reclamo.ID, Error: err, Failed: statusError}
			statusReclamos = append(statusReclamos, status)

		
	}

	if hayError {
		c.JSON(http.StatusNotFound, gin.H{
			"reclamos": statusReclamos,
			"error": gin.H{
				"status":  400,
				"message": "La actualización de los reclamos contiene errores",
			}})
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"reclamos": statusReclamos,
			"status": http.StatusOK,
		})
	}
}

func UpdateRepairDateReclamo(c *gin.Context) {
	var rec Models.Reclamo

	var now = time.Now().Unix()
	var updated = Utils.ConvertTimestampToDate(int64(now))

	if err := c.ShouldBindJSON(&rec); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	} else {

		err := Models.UpdateRepairDateReclamo(&rec, rec.ID, rec.Date_repair, updated)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": gin.H{
					"status":  http.StatusNotFound,
					"message": "Not Found",
				}})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"status": true})
		}
	}
}

func GetCantReclamosByUser(c *gin.Context) {

	var rec []Models.Reclamo
	userId := c.Query("userId")


	err := Models.GetAllReclamosByUser(&rec, userId)

	cantRec := len(rec)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": gin.H{
				"status":  http.StatusNotFound,
				"message": err.Error(),
			}})
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":   cantRec,
			"status": http.StatusOK,
		})
	}
}