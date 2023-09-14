package testutil

import (
	"app/src/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type SqlHandler struct {
	DB *gorm.DB
}

func NewSqlHandler() *SqlHandler {
	sqlHandler := new(SqlHandler)
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	connection := fmt.Sprintf("%s:%s@tcp(%s)/deep_track_test?charset=utf8&parseTime=True&loc=Local", user, pass, host)
	DB, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		log.Fatalln(connection + "database can't connect")
	}
	sqlHandler.DB = DB
	DB.AutoMigrate(model.Todos{})
	return sqlHandler
}

func TruncateTodoTable(handler SqlHandler) {
	handler.DB.Exec("TRUNCATE TABLE todos")
}
