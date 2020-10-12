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
	Config.DB.LogMode(true)
	validations.RegisterCallbacks(Config.DB)
	Config.DB.AutoMigrate(&Models.User{},&Models.Rubro{},&Models.Propiedad{},&Models.Unidad{})

	Config.DB.CreateTable(&Models.User{})
	Config.DB.CreateTable(&Models.Propiedad{})
	Config.DB.CreateTable(&Models.Unidad{})
	Config.DB.CreateTable(&Models.Rubro{})

	  rubros := []Models.Rubro{{Descripcion: "Electricista"},{Descripcion: "Plomero"},{Descripcion: "Gasista"}}
	  for _, r := range rubros {
		Config.DB.Create(&r)
	 }
	 propiedades := []Models.Propiedad{{Direccion:"Av libertador 5060",Nombre:"Chateau libertador", Localidad:"CABA", Provincia:"Buenos Aires",Lat:-34.563957, Lon:-58.4383696, Unidades:nil}}

	 for _, p := range propiedades {
		Config.DB.Omit("Unidades").Create(&p)
	 }

	 Config.DB.Table("RubroUsuario").AddForeignKey("user_id", "Usuarios(id)", "CASCADE", "CASCADE")
	 Config.DB.Table("RubroUsuario").AddForeignKey("rubro_id", "Rubros(id)", "CASCADE", "CASCADE")

	 Config.DB.Table("PropiedadUsuario").AddForeignKey("user_id", "Usuarios(id)", "CASCADE", "CASCADE")
	 Config.DB.Table("PropiedadUsuario").AddForeignKey("propiedad_id", "Propiedades(id)", "CASCADE", "CASCADE")

	 unidades := []Models.Unidad{{Piso:1, Depto:"e",Propiedad_id:1},{Piso:1, Depto:"f",Propiedad_id:1}}

	 for _, u := range unidades {
		Config.DB.Create(&u)
	 }

	 
	//  users := []Models.User{{Usertype:2,Nombre:"Maximiliano",Apellido:"Ferraiuolo",Email:"maxi08@gmail.com",Image:"imagen",Fechanac:"08/08/1985",Dni:"32326616",
	// 		Rubros:rubros,Propiedades:nil,Password:"08/10/2020",Is_active:false,Is_first_login:true}}

	// 	for _, usr := range users {
	// 	Config.DB.Omit("Rubros","Propiedades").Create(&usr)
	
	// 	}
			




	r := Routes.SetupRouter()
	//running
	r.Run()	
}
