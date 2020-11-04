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

const PRODUCT string = "Canada"

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
	Config.DB.CreateTable(&Models.Foto{})

	Config.DB.Delete(&Models.Rubro{})
	  rubros := []Models.Rubro{{Id: 1,Descripcion: "Electricista"},{Id:2,Descripcion: "Plomero"},{Id:3, Descripcion: "Gasista"},{Id:4, Descripcion: "Piletero"},{Id:5, Descripcion: "Limpieza"},{Id:6, Descripcion: "Albañileria"},{Id:7, Descripcion: "Otros"}}
	  for _, r := range rubros {
		Config.DB.Create(&r)
	 }
	 Config.DB.Delete(&Models.SharedArea{})

	 sharedAreas :=  []Models.SharedArea{{Id:1, Descripcion:"General", Image:"https://images.unsplash.com/photo-1527232165582-78c982a1cad1?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1000&q=80"},{Id:2, Descripcion:"Pileta", Image:"https://images.unsplash.com/photo-1585077017412-1b54a783b8ac?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=750&q=80"},{Id:3, Descripcion:"SUM",Image:"https://images.unsplash.com/photo-1531973968078-9bb02785f13d?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=677&q=80"},{Id:4, Descripcion:"GYM", Image:"https://images.unsplash.com/photo-1578874691223-64558a3ca096?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=375&q=80"},{Id:5, Descripcion:"CINE",Image:"https://images.unsplash.com/photo-1581250586548-0b44aeac2ad5?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=750&q=80"}}
	 sharedAreas2 :=  []Models.SharedArea{{Id:1, Descripcion:"General", Image:"https://images.unsplash.com/photo-1527232165582-78c982a1cad1?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1000&q=80"},{Id:2, Descripcion:"Pileta", Image:"https://images.unsplash.com/photo-1585077017412-1b54a783b8ac?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=750&q=80"},{Id:3, Descripcion:"SUM",Image:"https://images.unsplash.com/photo-1531973968078-9bb02785f13d?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=677&q=80"},{Id:4, Descripcion:"GYM", Image:"https://images.unsplash.com/photo-1578874691223-64558a3ca096?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=375&q=80"}}
	 sharedAreas3 :=  []Models.SharedArea{{Id:1, Descripcion:"General", Image:"https://images.unsplash.com/photo-1527232165582-78c982a1cad1?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1000&q=80"},{Id:2, Descripcion:"Pileta", Image:"https://images.unsplash.com/photo-1585077017412-1b54a783b8ac?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=750&q=80"},{Id:3, Descripcion:"SUM",Image:"https://images.unsplash.com/photo-1531973968078-9bb02785f13d?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=677&q=80"}}
	 sharedAreas4 :=  []Models.SharedArea{{Id:1, Descripcion:"General", Image:"https://images.unsplash.com/photo-1527232165582-78c982a1cad1?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1000&q=80"},{Id:2, Descripcion:"Pileta", Image:"https://images.unsplash.com/photo-1585077017412-1b54a783b8ac?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=750&q=80"},{Id:3, Descripcion:"SUM",Image:"https://images.unsplash.com/photo-1531973968078-9bb02785f13d?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=677&q=80"}}


	 for _, s := range sharedAreas {
		Config.DB.Create(&s)
	 }

	 propiedades := []Models.Propiedad{{Direccion:"Av libertador 101",Nombre:"Al Rio", Localidad:"Vicente Lopez", Provincia:"Buenos Aires", Image: "https://media.gettyimages.com/photos/modern-condominiums-picture-id952954132?s=2048x2048" ,Lat:-34.563957, Lon:-58.4383696, SharedAreas:sharedAreas2},
	 {Direccion:"Av Libertador 7050",Nombre:"Chateau libertador 2", Localidad:"CABA", Provincia:"Buenos Aires",Image: "https://images.unsplash.com/photo-1589095093845-93a387ebc7ee?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=750&q=80" ,Lat:-34.563957, Lon:-58.4383696,SharedAreas:sharedAreas},
	 {Direccion:"Av Figueroa Alcorta 3600",Nombre:"Le Parc", Localidad:"CABA", Provincia:"Buenos Aires", Image: "https://images.unsplash.com/photo-1535010827831-0f324b6f6908?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=321&q=80" ,Lat:-34.563957, Lon:-58.4383696, SharedAreas:sharedAreas3},
	 {Direccion:"Azucena Villaflor 559",Nombre:"Edificio Alvear", Localidad:"Puerto Madero", Provincia:"Buenos Aires", Image: "https://images.unsplash.com/photo-1535010827831-0f324b6f6908?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=321&q=80" ,Lat:-34.563957, Lon:-58.4383696,  SharedAreas:sharedAreas4}}
	 
	 
	 
	 for _, p := range propiedades {
		Config.DB.Create(&p)
	 }

	 Config.DB.Table("RubroUsuario").AddForeignKey("user_id", "Usuarios(id)", "CASCADE", "CASCADE")
	 Config.DB.Table("RubroUsuario").AddForeignKey("rubro_id", "Rubros(id)", "CASCADE", "CASCADE")

	 Config.DB.Table("UnidadUsuario").AddForeignKey("user_id", "Usuarios(id)", "CASCADE", "CASCADE")
	 Config.DB.Table("UnidadUsuario").AddForeignKey("unidad_id", "Unidades(id)", "CASCADE", "CASCADE")

	 Config.DB.Table("SharedReclamo").AddForeignKey("reclamo_id", "Reclamos(id)", "CASCADE", "CASCADE")
	 Config.DB.Table("SharedReclamo").AddForeignKey("shared_area_id", "Shared(id)", "CASCADE", "CASCADE")


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
								{Piso:4, Depto:"A", PropiedadId:3},{Piso:4, Depto:"B",PropiedadId:3},{Piso:4, Depto:"C",PropiedadId:3},
								{Piso:1, Depto:"A", PropiedadId:4},{Piso:1, Depto:"B",PropiedadId:4},{Piso:5, Depto:"C",PropiedadId:4},
	                            {Piso:2, Depto:"A", PropiedadId:4},{Piso:2, Depto:"B",PropiedadId:4},{Piso:2, Depto:"C",PropiedadId:4},
								{Piso:3, Depto:"A", PropiedadId:4},{Piso:3, Depto:"B",PropiedadId:4},{Piso:3, Depto:"C",PropiedadId:4}}

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
