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
	//if err := db.CreateTable(new(models.Account)).Error; err != nil {
	//	return fmt.Errorf("Error creating AccountUser table: %s", err)
	//}

	return nil
}
