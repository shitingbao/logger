package logdriver

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenMysql(user, pas, host, dataBase, port string) (*gorm.DB, error) {
	db, err := openMysql(user, pas, host, dataBase, port)
	if err != nil {
		return nil, err
	}
	return db, nil
}

//SQLOpen sqlopen
func openMysql(user, pas, host, dataBase, port string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("root:12345678@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Println("orm:", err)
		return db, err
	}
	return db, nil
}
