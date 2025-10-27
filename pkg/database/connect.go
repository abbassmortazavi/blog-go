package database

import (
	"blog/pkg/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	cfg := config.Get()
	fmt.Println(cfg.DBNAME, cfg.DBHOST, cfg.DBPORT, cfg.DBUSERNAME, cfg.DBPASSWORD)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUSERNAME,
		cfg.DBPASSWORD,
		cfg.DBHOST,
		cfg.DBPORT,
		cfg.DBNAME,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to database", err.Error())
		return
	}
	log.Println("Successfully connected to database!")
	DB = db
}
