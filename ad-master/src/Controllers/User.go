package Controllers

import (
	"first-api/Config"
	"strconv"
	"first-api/Models"
	"first-api/Utils"
	"fmt"
	"log"
	"time"
	"net/http" //https://golang.org/pkg/net/http/
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/gin-gonic/gin"
)

const password_expiration_lapse = 7776000 //3 months

//GetUsers ... Get all users
func GetUsers(c *gin.Context) {
	var user []Models.User
	var u Models.User
	err := Models.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		log.Println("====== Bind By Query String ======")
		log.Println(u.Nombre)
 
		fmt.Println(c.Request.URL.Query())
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "3"))
	
		paginator := pagination.Paging(&pagination.Param{
			DB:      Config.DB,
			Page:    page,
			Limit:   limit,
			OrderBy: []string{"id"},
			ShowSQL: true,
		}, &user)
		c.JSON(200, paginator)

	//	c.JSON(http.StatusOK, gin.H{"data": user})
	}
}


//CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)
	var now = time.Now().Unix()
	user.Password = Utils.EncodeBase64(user.Password)
    user.Date_created = Utils.ConvertTimestampToDate(int64(now))
	user.Password_expires = Utils.ConvertTimestampToDate(int64(now + password_expiration_lapse))
	err := Models.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.User
	err := Models.GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//UpdateUser ... Update the user information
func UpdateUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	fmt.Println("id", id)


	err := Models.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = Models.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}

//LoginUser
func LoginUser(c *gin.Context) {
	var user Models.User
	var dbuser Models.User
	 uid := c.Params.ByName("name")
	 fmt.Println("name", uid)
	 firstname := c.Query("firstname")
	 lastname := c.Query("lastname") 
	 fmt.Println("firstname", firstname)
	 fmt.Println("lastname", lastname)

	 	if err := c.ShouldBindJSON(&user); err != nil {
 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	 		return
	 }
	 fmt.Println(user.Email)

	 err := Models.LoginUser(&dbuser, user.Email, Utils.EncodeBase64(user.Password))
	 if err != nil {
		 c.AbortWithStatus(http.StatusNotFound)
		 fmt.Println("Not found")
		 c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		 return
	 } else {
		fmt.Println("User Found")
		if(dbuser.Is_active){
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "user inactive"})
		}
	}
}



//// Example for binding JSON ({"user": "manu", "password": "123"})
	//router.POST("/loginJSON", func(c *gin.Context) {
	//	var json Login
	//	if err := c.ShouldBindJSON(&json); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//		return
	//	}
		
	////	if json.User != "manu" || json.Password != "123" {
	//		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	//		return
	//	} 
		
	//	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
//	})
