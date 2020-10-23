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
	Config.DB.AutoMigrate(&Models.User{},&Models.Rubro{},&Models.Propiedad{},&Models.Unidad{},&Models.SharedArea{})

	Config.DB.CreateTable(&Models.User{})
	Config.DB.CreateTable(&Models.Propiedad{})
	Config.DB.CreateTable(&Models.Unidad{})
	Config.DB.CreateTable(&Models.Rubro{})
	Config.DB.CreateTable(&Models.SharedArea{})
	 Config.DB.CreateTable(&Models.Reclamo{})
	 Config.DB.CreateTable(&Models.Comentario{})
	 Config.DB.CreateTable(&Models.Notificacion{})
	//Config.DB.CreateTable(&Models.Foto{})

	  rubros := []Models.Rubro{{Descripcion: "Electricista"},{Descripcion: "Plomero"},{Descripcion: "Gasista"}}
	  for _, r := range rubros {
		Config.DB.Create(&r)
	 }

	 sharedAreas :=  []Models.SharedArea{{Descripcion:"Pileta"},{Descripcion:"Sum"},{Descripcion:"GYM"}}
	//  for _, shared := range sharedAreas {
	// 	Config.DB.Create(&shared)
	//  }

	 propiedades := []Models.Propiedad{{Direccion:"Av libertador 101",Nombre:"Al Rio", Localidad:"Vicente Lopez", Provincia:"Buenos Aires", Image: "https://media.gettyimages.com/photos/modern-condominiums-picture-id952954132?s=2048x2048" ,Lat:-34.563957, Lon:-58.4383696,SharedAreas:sharedAreas},
	 {Direccion:"Av Libertador 7050",Nombre:"Chateau libertador 2", Localidad:"CABA", Provincia:"Buenos Aires",Image: "https://images.unsplash.com/photo-1589095093845-93a387ebc7ee?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=750&q=80" ,Lat:-34.563957, Lon:-58.4383696,SharedAreas:sharedAreas},
	 {Direccion:"Av Figueroa Alcorta 3600",Nombre:"Le Parc", Localidad:"CABA", Provincia:"Buenos Aires", Image: "https://lh3.googleusercontent.com/proxy/6MG-osqJE_roLB90JhBkuike-jqixSx4qvLuMUhAZuNY2vL7-MkOYzUBfnIlCsn1mEjQGYmVkIg9lJShfvky3uS1xjTphhfdGCxYO3zLFXOsIFFP2Q3qREjNn0PXVL_cQkU4WipcLjGfbJ9a5G7q" ,Lat:-34.563957, Lon:-58.4383696},
	 {Direccion:"Azucena Villaflor 559",Nombre:"Edificio Alvear", Localidad:"Puerto Madero", Provincia:"Buenos Aires", Image: "https://images.unsplash.com/photo-1535010827831-0f324b6f6908?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=321&q=80" ,Lat:-34.563957, Lon:-58.4383696},
	 {Direccion:"Azucena Villaflor 559",Nombre:"Edificio Alvear", Localidad:"Puerto Madero", Provincia:"Buenos Aires", Image: "https://images.unsplash.com/photo-1535010827831-0f324b6f6908?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=321&q=80" ,Lat:-34.563957, Lon:-58.4383696},
	 {Direccion:"Azucena Villaflor 559",Nombre:"Edificio Alvear", Localidad:"Puerto Madero", Provincia:"Buenos Aires", Image: "https://images.unsplash.com/photo-1535010827831-0f324b6f6908?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=321&q=80" ,Lat:-34.563957, Lon:-58.4383696},
	 {Direccion:"Azucena Villaflor 559",Nombre:"Edificio Alvear", Localidad:"Puerto Madero", Provincia:"Buenos Aires", Image: "https://images.unsplash.com/photo-1535010827831-0f324b6f6908?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=321&q=80" ,Lat:-34.563957, Lon:-58.4383696},
	 {Direccion:"Azucena Villaflor 559",Nombre:"Edificio Alvear", Localidad:"Puerto Madero", Provincia:"Buenos Aires", Image: "https://images.unsplash.com/photo-1535010827831-0f324b6f6908?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=321&q=80" ,Lat:-34.563957, Lon:-58.4383696},
	 {Direccion:"Azucena Villaflor 559",Nombre:"Edificio Alvear", Localidad:"Puerto Madero", Provincia:"Buenos Aires", Image: "https://images.unsplash.com/photo-1535010827831-0f324b6f6908?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=321&q=80" ,Lat:-34.563957, Lon:-58.4383696}}
	 
	 
	 
	 for _, p := range propiedades {
		Config.DB.Create(&p)
	 }

	 Config.DB.Table("RubroUsuario").AddForeignKey("user_id", "Usuarios(id)", "CASCADE", "CASCADE")
	 Config.DB.Table("RubroUsuario").AddForeignKey("rubro_id", "Rubros(id)", "CASCADE", "CASCADE")

	 Config.DB.Table("UnidadUsuario").AddForeignKey("user_id", "Usuarios(id)", "CASCADE", "CASCADE")
	 Config.DB.Table("UnidadUsuario").AddForeignKey("unidad_id", "Unidades(id)", "CASCADE", "CASCADE")

	 Config.DB.Table("SharedReclamo").AddForeignKey("reclamo_id", "Reclamos(id)", "CASCADE", "CASCADE")
	 Config.DB.Table("SharedReclamo").AddForeignKey("shared_area_id", "Shared(id)", "CASCADE", "CASCADE")



	 fmt.Println("Status:", propiedades[0].Nombre)
	 unidades := []Models.Unidad{{Piso:1, Depto:"A", PropiedadId:1},{Piso:1, Depto:"B",PropiedadId:1},{Piso:5, Depto:"C",PropiedadId:1},
	                            {Piso:2, Depto:"A", PropiedadId:1},{Piso:2, Depto:"B",PropiedadId:1},{Piso:2, Depto:"C",PropiedadId:1},
								{Piso:3, Depto:"A", PropiedadId:1},{Piso:3, Depto:"B",PropiedadId:1},{Piso:3, Depto:"C",PropiedadId:1},
								{Piso:4, Depto:"A", PropiedadId:1},{Piso:4, Depto:"B",PropiedadId:1},{Piso:4, Depto:"C",PropiedadId:1},
								{Piso:1, Depto:"A", PropiedadId:2},{Piso:1, Depto:"B",PropiedadId:2},{Piso:5, Depto:"C",PropiedadId:2},
	                            {Piso:2, Depto:"A", PropiedadId:2},{Piso:2, Depto:"B",PropiedadId:2},{Piso:2, Depto:"C",PropiedadId:2},
								{Piso:3, Depto:"A", PropiedadId:2},{Piso:3, Depto:"B",PropiedadId:2},{Piso:3, Depto:"C",PropiedadId:2},
								{Piso:4, Depto:"A", PropiedadId:2},{Piso:4, Depto:"B",PropiedadId:2},{Piso:4, Depto:"C",PropiedadId:2},
								{Piso:1, Depto:"A", PropiedadId:3},{Piso:1, Depto:"B",PropiedadId:3},{Piso:5, Depto:"C",PropiedadId:3},
	                            {Piso:2, Depto:"A", PropiedadId:3},{Piso:2, Depto:"B",PropiedadId:3},{Piso:2, Depto:"C",PropiedadId:3},
								{Piso:3, Depto:"A", PropiedadId:3},{Piso:3, Depto:"B",PropiedadId:3},{Piso:3, Depto:"C",PropiedadId:3},
								{Piso:4, Depto:"A", PropiedadId:3},{Piso:4, Depto:"B",PropiedadId:3},{Piso:4, Depto:"C",PropiedadId:3}}

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
