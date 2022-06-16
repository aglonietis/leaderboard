package main

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
	"leaderboard/internal/database"
	"leaderboard/internal/models"
)

func main() {
	configFileName := ".env"

	// Load Viper
	viper.SetConfigFile(configFileName)
	err := viper.ReadInConfig()

	if err != nil {
		panic("Unable to load configuration file " + configFileName + ":" + err.Error())
	}

	// Load database
	err = database.Init()

	if err != nil {
		panic("Unable to establish database connection: " + err.Error())
	}

	err = MigrateModels()

	if err != nil {
		panic("Failed to migrate Models for database: " + err.Error())
	}

	err = MigrateRaw()

	if err != nil && err.Error() != "no change" {
		panic("Failed to migrate Raw queries for database: " + err.Error())
	}

	fmt.Println("Succesfuly migrated database")
}

func MigrateModels() error {
	db := database.DbManager()

	err := db.AutoMigrate(&models.PlayerScore{})

	if err != nil {
		return err
	}

	return nil
}

//https://pkg.go.dev/github.com/golang-migrate/migrate/v4#section-readme
func MigrateRaw() error {

	//fsrc, err := (&file.File{}).Open("file://migrations")
	//if err != nil {
	//	return err
	//}

	db, err := sql.Open(database.GetDatabaseConnection(), database.GetDatabaseUrl())

	if err != nil {
		return err
	}

	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
"file://database/migrations",
		database.GetDatabaseName(),driver,
	)

	if err != nil {
		return err
	}

	step, dirty, err := m.Version()

	// TODO: Remove after development
	m.Down()

	if dirty {
		err := m.Force(int(step-1))

		if err != nil {
			return err
		}
	}

	return m.Up()
}
