package Models

import (
	"first-api/Config"
	"fmt"
	
	_ "github.com/go-sql-driver/mysql"
	//http://gorm.io/es_ES/docs/query.html
)
 
//GetAllUsers Fetch all user data
func GetAllUsers(user *[]User) (err error) {
			
 
	// if err = Config.DB.Model(&User{}).Select("*").Joins("join RubroUsuario on RubroUsuario.id_usuario = Usuarios.Id").Find(&User{}).Error; err != nil {

	// 	return err
	// }
	// return nil
	
	if err = Config.DB.Preload("Rubros").Find(&user).Error; err != nil {
		return err
	}
	return nil
}

//CreateUser ... Insert New data
func CreateUser(user *User) (err error) { 


     if err = Config.DB.Omit("Rubros", "Propiedades").Create(user).Error; err != nil {
	   return err
     }
   if err =  Config.DB.Model(user).Association("Rubros").Replace(user.Rubros).Error; err != nil {
		return err
	 }

	if err =  Config.DB.Model(user).Association("Propiedades").Replace(user.Propiedades).Error; err != nil {
		return err
	 }
 
	return nil
}

 
//GetUserByID ... Fetch only one user by Id
func GetUserByID(user *User, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).Find(user).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser ... Update user
func UpdateUser(user *User, id string) (err error) {
	fmt.Println(user)
	Config.DB.Save(user)
	return nil
}

//DeleteUser ... Delete user
func DeleteUser(user *User, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}

//LoginUser
func LoginUser(user *User, email string) (err error) {
	if err = Config.DB.Where("email = ?", email).Find(user).Error; err != nil {
		return err
	}

	return nil
}

//UpdateUser ... Update user
func UpdateUserById(user *User, id string) (err error) {
	fmt.Println(user)
	Config.DB.Where("id = ?", id).Save(&user)
	return nil
}

func UpdateUserByEmail(user *User, email string, psw string, isFirstLogin bool) (err error) {

	if err = Config.DB.Model(user).Where("email = ?", email).Find(user).Updates(map[string]interface{}{"password": psw, "is_first_login": isFirstLogin}).Error; err != nil {
		return err
	}
	return nil
}