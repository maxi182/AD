package Models

type RubroUsuario struct {
	Id      uint   `gorm:"primary_key"`
	Id_rubro uint   `json:"id_rubro"`
	Id_usuario uint   `json:"id_usuario"`
}

func (b *RubroUsuario) TableName() string {
	return "RubroUsuario"
}
