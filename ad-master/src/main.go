package main

import (
	"first-api/Config"
	"first-api/Models"
	"first-api/Routes"
	"fmt"

	"github.com/qor/validations"
	"github.com/jinzhu/gorm"
)

//https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/

//https://www.soberkoder.com/go-rest-api-mysql-gorm/
//https://stackoverflow.com/questions/34667199/gorm-many-to-many-select-gives-invalid-association-error
//https://github.com/jinzhu/gorm/blob/021d7b33143de37b743d1cf660974e9c8d3f80ea/multi_primary_keys_test.go
//https://medium.com/remotepanda-blog/go-with-gorm-chapter-10-golang-91bc5d01c161
var err error

func main() {
	 Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()
	validations.RegisterCallbacks(Config.DB)
	Config.DB.AutoMigrate(&Models.User{})
	Config.DB.AutoMigrate(&Models.Rubro{})
	Config.DB.LogMode(true)

	//Config.DB.DropTableIfExists(&Models.User{})
	Config.DB.CreateTable(&Models.Propiedad{})
	Config.DB.CreateTable(&Models.Unidad{})
	Config.DB.CreateTable(&Models.User{})
	Config.DB.CreateTable(&Models.Rubro{})

	//Models.NewPropiedad(1,"Av Libertador 5000","Edificio Libertador A","CABA", "Buenos Aires", -34.563957,-58.4383696)
 
	r := Routes.SetupRouter()
	//running
	r.Run()	
}
