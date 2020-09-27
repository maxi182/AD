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
			DB:      Config.DB,
			Page:    page,
			Limit:   limit,
			OrderBy: []string{"id"},
			ShowSQL: true,
		}, &user)
 
 		
		c.JSON(200, paginator)

	}
}

//CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)
	//var rubro []Models.Rubro
	//rubro = user.Rubros

// 	if(rubro !=nil && len(rubro)>0) {
// 	fmt.Println(rubro[1].Id) //Trae el ID del rubro en pos1
// 	rubro[0].Id_rubro = rubro[0].Id
// 	rubro[0].Id_usuario = user.Id
// 	err1 := Models.CreateRubroUsuario(&rubro[0])

// 	if err1 != nil {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"error" : gin.H { 
// 			"status":  400,
// 			"message": err1.Error(),
// 		}})
// 		fmt.Println(err1.Error())
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, user)
// 	}
//  }
	
	var now = time.Now().Unix()
	user.Password = Utils.EncodeBase64(user.Password)
    user.Date_created = Utils.ConvertTimestampToDate(int64(now))
	err := Models.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : gin.H { 
			"status":  400,
			"message": err.Error(),
		}})
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
		fmt.Println("usuario_creado", user.Id)
 
		// 	if(rubro !=nil && len(rubro)>0) {
		// 		for i := 0; i < len(rubro); i++ {
		// 			rubro[i].Id_usuario = user.Id
		// 			Models.CreateRubro(&rubro[i])
		// 		}
		// }
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
		 c.JSON(http.StatusUnauthorized, gin.H{
			"error" : gin.H { 
			"status":  http.StatusUnauthorized,
			"message": "Unathorized",
		}})
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
