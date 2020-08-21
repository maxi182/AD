package main

import (
	"first-api/Config"
	"first-api/Models"
	"first-api/Routes"
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	 Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	//Config.DB, err = gorm.Open("mysql", "cerfogli_distri:uade12345@ar-caba-sv1.seconline.net.ar/cerfogli_uade?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})
	Config.DB.LogMode(true)

	r := Routes.SetupRouter()
	//running
	r.Run()
}
