package services

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func init() {
	Logger.Info("Initializing the DB")
	dbConfig := ApplicationConfig.Database
	datasource := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Database, dbConfig.Password)
	Logger.Info(datasource)
	db, err := gorm.Open("postgres", datasource)
	if err != nil {
		Logger.Fatal("Error connecting the db", err)
	}
	Logger.Info("Database Initialized!")
	DB = db
}
