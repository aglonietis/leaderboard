package database

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {

	dbUrl := GetDatabaseUrl()

	databaseConnection, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		return err
	}

	db = databaseConnection

	return nil
}

func GetDatabaseUrl() string {
	return 	"postgres://"+viper.GetString("DB_USER")+":"+viper.GetString("DB_PASSWORD") +
		"@"+viper.GetString("DB_HOST")+":"+viper.GetString("DB_PORT")+"/"+viper.GetString("DB_DATABASE")
}

func DbManager() *gorm.DB {
	return db
}