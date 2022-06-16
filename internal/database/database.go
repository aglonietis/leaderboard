package database

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Init() error {

	dbUrl := GetDatabaseUrl()

	databaseConnection, err := gorm.Open(postgres.Open(dbUrl), getGormDatabaseConfig())

	if err != nil {
		return err
	}

	db = databaseConnection

	return nil
}

func GetDatabaseUrl() string {
	return GetDatabaseConnection() + "://" + viper.GetString("DB_USER") + ":" + viper.GetString("DB_PASSWORD") +
		"@" + viper.GetString("DB_HOST") + ":" + viper.GetString("DB_PORT") + "/" + GetDatabaseName() + "?sslmode=" + viper.GetString("DB_SSL_MODE")
}

func DbManager() *gorm.DB {
	return db
}

func GetDatabaseName() string {
	return viper.GetString("DB_DATABASE")
}

func GetDatabaseConnection() string {
	return viper.GetString("DB_CONNECTION")
}

func getGormDatabaseConfig() *gorm.Config {
	gormConfig := &gorm.Config{}

	if viper.GetBool("DB_LOGGER_ENABLED") {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	return gormConfig
}
