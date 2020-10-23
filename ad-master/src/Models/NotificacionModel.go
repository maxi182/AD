package Models

type Notificacion struct {

	Id        		uint      `gorm:"primary_key;auto_increment" json:"comentario_id"` 
	ReclamoId		uint	  `json:"reclamo_id"`
	Date_created    string    `json:"date_created"`
	Type    		uint      `json:"type"`
	UsuarioId       uint      `json:"usuario_id"`
	Usuario			User	  `gorm:"foreignKey:UsuarioId" json:"usuario"`  

}

func (b *Notificacion) TableName() string {
	return "Notificaciones"
}