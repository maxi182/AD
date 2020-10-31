package Controllers
import (
   "first-api/Models"
   "net/http"
   "github.com/gin-gonic/gin"
)


func GetAllRubros(c *gin.Context) {
	var rubro []Models.Rubro
	err := Models.GetAllRubros(&rubro)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : gin.H { 
			"status":  http.StatusNotFound,
			"message": err.Error(),
		}})
		c.AbortWithStatus(http.StatusNotFound)
	} else {
 
		c.JSON(http.StatusOK,gin.H{
			"data" : rubro,
			"status":  http.StatusOK,
		})
	}
}
