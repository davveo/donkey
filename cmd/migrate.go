package cmd

import (
	"github.com/davveo/donkey/models/initial"
	"github.com/davveo/donkey/utils/migrations"
)

func Migrate(configFile string) error {
	_, db, err := initConfig(configFile)
	if err != nil {
		return err
	}
	defer db.Close()

	// Bootstrap migrations
	if err := migrations.Bootstrap(db); err != nil {
		return err
	}

	// Run migrations for the oauth service
	if err := initial.MigrateAll(db); err != nil {
		return err
	}
	return nil
}
