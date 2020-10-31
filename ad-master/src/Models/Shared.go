
package Models

import (
	"first-api/Config"
	_ "github.com/go-sql-driver/mysql"
 
)
 
func GetAllShared(shared *[]SharedArea, propId string) (err error) {
			
	if err = Config.DB.Model(&SharedArea{}).Select("*").Joins("inner join SharedPropiedad on Shared.id = SharedPropiedad.shared_area_id").Where("SharedPropiedad.propiedad_id = ?", propId).Find(&shared).Error; err != nil {
		return err
   }
    return nil
}


 