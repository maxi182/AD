package Controllers

import (
 	"first-api/Config"
 	"strconv"
	"first-api/Models"
	"first-api/Utils"
	"fmt"
	"log"
	"time"

	"gopkg.in/validator.v2"
	"net/http" //https://golang.org/pkg/net/http/
 	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/gin-gonic/gin"
)

//GetUsers ... Get all users
func GetUsers(c *gin.Context) {
	var user []Models.User
	var u Models.User
	err := Models.GetAllUsers(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : gin.H { 
			"status":  http.StatusNotFound,
			"message": err.Error(),
		}})
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		log.Println("====== Bind By Query String ======")
		log.Println(u.Nombre)
		//var rubro []Models.RubroUsuario
	 
		fmt.Println(c.Request.URL.Query())
		 page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		 limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
	
		  paginator := pagination.Paging(&pagination.Param{
			DB:      Config.DB.Preload("Rubros").Preload("Unidades"),
			Page:    page,
			Limit:   limit,
			OrderBy: []string{"id"},
			ShowSQL: true,
		},  &user)
 
		c.JSON(200, paginator)

	}
}

//CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)

	auth := c.Request.Header.Get("Authorization")
    if auth !=  Utils.GetAuthToken() {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error" : gin.H { 
			"status":  http.StatusUnauthorized,
			"message": "Invalid Token",
		}})
		c.Abort()
		return
	}
	var now = time.Now().Unix()
	nur := Models.User(user)
	 err_password := validator.Validate(nur)
	 if err_password != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : gin.H { 
			"status":  http.StatusNotFound,
			"message": err_password.Error(),
		}})
		fmt.Println(err_password.Error())
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	//user.Password = Utils.EncodeBase64(user.Password)
	user.Date_created = Utils.ConvertTimestampToDate(int64(now))
	err := Models.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : gin.H { 
			"status":  http.StatusNotFound,
			"message": err.Error(),
		}})
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
		fmt.Println("usuario_creado", user.Id)
	}
}

//GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.User
	err := Models.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : gin.H { 
			"status":  http.StatusNotFound,
			"message": err.Error(),
		}})
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK,gin.H{
			"data" : user,
			"status":  http.StatusOK,
		})
	}
}

//UpdateUser ... Update the user information
func UpdateUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	fmt.Println("id", id)
	err := Models.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound,  gin.H{
			"error" : gin.H { 
			"status":  http.StatusNotFound,
			"message": "Not Found",
		}})
		return
	} else {
	c.BindJSON(&user)
	
	err = Models.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		c.JSON(http.StatusNotFound,  gin.H{
			"data":gin.H { 
			"error" : gin.H { 
			"status":  http.StatusBadRequest,
			"message": "Can´t update user",
		}}})
	} else {
		c.JSON(http.StatusOK, user)
	}
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
		c.JSON(http.StatusOK, gin.H{"id":"is deleted"})
	}
}

//LoginUser
func LoginUser(c *gin.Context) {
	var user Models.User
	var dbuser Models.User
 
	auth := c.Request.Header.Get("Authorization")
    if auth !=  Utils.GetAuthToken() {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error" : gin.H { 
			"status":  http.StatusUnauthorized,
			"message": "Invalid Token",
		}})
		c.Abort()
		return
	}

	 	if err := c.ShouldBindJSON(&user); err != nil {
 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	 		return
	 }
	 fmt.Println(user.Email)

	 err := Models.LoginUser(&dbuser, user.Email)
	 if err != nil {
		 c.AbortWithStatus(http.StatusNotFound)
		 fmt.Println("Not found")
		 c.JSON(http.StatusNotFound,  gin.H{
			"error" : gin.H { 
			"status":  http.StatusNotFound,
			"message": "User not found",
		}})
		 return
	 } else {
		fmt.Println("User Found")
		if(dbuser.Password==user.Password){
		if(dbuser.Is_active) {
			c.JSON(http.StatusOK,gin.H{
				"data" : dbuser,
				"status":  http.StatusOK,
			})
		} else {
			c.JSON(http.StatusNotFound,  gin.H{
				"error" : gin.H { 
				"status":  http.StatusUnauthorized,
				"message": "User blocked",
			}})
		}
	}else{
		c.JSON(http.StatusNotFound,  gin.H{
			"error" : gin.H { 
			"status":  http.StatusUnauthorized,
			"message": "Incorrect password",
		}})
	}
   }
}

func SendEmail(c *gin.Context) {

var email Models.Email
var user Models.User

if err := c.ShouldBindJSON(&email); err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"status": false})
		return
} else {
	var code = Utils.StringWithCharset(6)
	Utils.SendEmail(code,email.Email)
	user.Password = code
	err := Models.UpdateUserByEmail(&user, email.Email, code, true)
   if err != nil {
	c.JSON(http.StatusNotFound,  gin.H{
		"error" : gin.H { 
		"status":  http.StatusNotFound,
		"message": "Not Found",
	}})
	return
   } else {
	c.JSON(http.StatusOK, gin.H{"status": true})
   }
  }
}

func UpdatePassword(c *gin.Context) {
	var user Models.User
	
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}else {

		err := Models.UpdateUserByEmail(&user, user.Email, user.Password, false)
	   if err != nil {
		c.JSON(http.StatusNotFound,  gin.H{
			"error" : gin.H { 
			"status":  http.StatusNotFound,
			"message": "Not Found",
		}})
		return
	   } else {
		c.JSON(http.StatusOK, gin.H{"status": true})
	   }
	  }
	}


//UPDATE USER BY EMAIL TO RESET PASSWORD




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
