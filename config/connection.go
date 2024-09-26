package config

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"

	"github.com/RaihanMalay21/web-gudang/models"
)

var (
	DB *gorm.DB
)

func DB_Connection() {
	var (
		dbUser = "root"
		dbPassword = "0987"
		dbHostAndPort = "localhost:3306"
		dbName = "manage_warehouse"
	)

	DNS := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbUser, dbPassword, dbHostAndPort, dbName)

	db, err := gorm.Open(mysql.Open(DNS))
	if err != nil {
		fmt.Println(err)
		return
	}

	db.AutoMigrate(&models.Shelf{})
	db.AutoMigrate(&models.Row{})
	db.AutoMigrate(&models.Block{})
	db.AutoMigrate(&models.Barang{})

	DB = db
}