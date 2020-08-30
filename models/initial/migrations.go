package initial

import (
	"github.com/davveo/donkey/utils/migrations"
	"github.com/jinzhu/gorm"
)

var (
	list = []migrations.MigrationStage{
		{
			Name:     "initial",
			Function: migrate0001,
		},
	}
)

func MigrateAll(db *gorm.DB) error {
	return migrations.Migrate(db, list)
}

func migrate0001(db *gorm.DB, name string) error {
	//if err := db.CreateTable(new(models.Admin)).Error; err != nil {
	//	return fmt.Errorf("Error creating Admin table: %s", err)
	//}
	//if err := db.CreateTable(new(models.AuthGroup)).Error; err != nil {
	//	return fmt.Errorf("Error creating AuthGroup table: %s", err)
	//}
	//if err := db.CreateTable(new(models.AuthRule)).Error; err != nil {
	//	return fmt.Errorf("Error creating AuthRule table: %s", err)
	//}
	//if err := db.CreateTable(new(models.App)).Error; err != nil {
	//	return fmt.Errorf("Error creating App table: %s", err)
	//}
	//if err := db.CreateTable(new(models.Token)).Error; err != nil {
	//	return fmt.Errorf("Error creating Token table: %s", err)
	//}

	return nil
}
