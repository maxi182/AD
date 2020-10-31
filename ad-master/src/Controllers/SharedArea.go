package Controllers

import (
	"first-api/Models"
	"net/http"
	"github.com/gin-gonic/gin"
 )
 
 
 func GetAllshared(c *gin.Context) {
	 var shared []Models.SharedArea
	 propId := c.Query("propId")
	 
	 err := Models.GetAllShared(&shared, propId)
	 if err != nil {
		 c.JSON(http.StatusNotFound, gin.H{
			 "error" : gin.H { 
			 "status":  http.StatusNotFound,
			 "message": err.Error(),
		 }})
		 c.AbortWithStatus(http.StatusNotFound)
	 } else {
  
		 c.JSON(http.StatusOK,gin.H{
			 "data" : shared,
			 "status":  http.StatusOK,
		 })
	 }
 }
 