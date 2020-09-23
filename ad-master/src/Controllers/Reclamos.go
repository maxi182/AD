package Controllers

import (
	"first-api/Models"
	"fmt"
	"net/http" //https://golang.org/pkg/net/http/
	"github.com/gin-gonic/gin"
)


//CreateUser ... Create User
func CreateReclamo(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)

	err := Models.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}