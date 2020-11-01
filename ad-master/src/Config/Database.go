package Config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

//   func BuildDBConfig() *DBConfig {
//   	dbConfig := DBConfig{
// 		Host:     "localhost",
//   		Port:     3306,
//   		User:     "root",
//         Password: "123456",
// 		 DBName:   "testdb",

// 	}
//  	return &dbConfig
//  }

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "10.0.6.3",
		Port:     3306,
		User:     "tpuser",
		Password: "Uade$12345",
		DBName:   "tpdb",
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
