package main

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/lukinairina90/L_Scraper/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// ENV Carloos
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	fmt.Printf("CONFIGGGG  %+v\n", cfg)

	// смотрите https://github.com/go-sql-driver/mysql#dsn-data-source-name для подробностей
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Login, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	// Миграция схем
	if err = db.AutoMigrate(&models.Property{}, &models.PropertyCharacteristic{}); err != nil {
		return
	}

	sc := scraper{
		cfg: cfg,
		db:  db,
	}

	if err := sc.collectData(); err != nil {
		fmt.Printf("error has happened  %v\n", err)
	}

	fmt.Println("success!")
}
