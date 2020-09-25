package main

import (
	"first-api/Config"
	"first-api/Models"
	"first-api/Routes"
	"fmt"


	"github.com/jinzhu/gorm"
)

//https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/

var err error

func main() {
	 Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})
	Config.DB.LogMode(true)

	//Config.DB.DropTableIfExists(&Models.User{})
	Config.DB.CreateTable(&Models.Propiedad{})
	Config.DB.CreateTable(&Models.Unidad{})
	//Models.NewPropiedad(1,"Av Libertador 5000","Edificio Libertador A","CABA", "Buenos Aires", -34.563957,-58.4383696)
 
	r := Routes.SetupRouter()
	//running
	r.Run()	
}
