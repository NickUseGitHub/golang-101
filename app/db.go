package app

import (
	"fmt"

	"github.com/NickUseGitHub/golang-101/app/models"
	"github.com/NickUseGitHub/golang-101/configs"

	"github.com/jinzhu/gorm"
	// Register for connection to postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DBApp for initial DB App
type DBApp struct {
	DB *gorm.DB
}

// InitializeDB to initial app's database
func (dbApp *DBApp) InitializeDB(cnfs configs.Configs) error {
	db, err := gorm.Open("postgres", cnfs.GetDbConfigConnection(cnfs))
	if err != nil {
		return err
	}

	fmt.Println("[√]:: DB Connected")
	dbApp.DB = db
	dbApp.migrationDB()

	return nil
}

func (dbApp *DBApp) migrationDB() {
	dbApp.DB.AutoMigrate(&models.Todo{})
}
