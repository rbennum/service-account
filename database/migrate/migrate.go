package migrate

import (
	"github.com/golang-migrate/migrate"
	"github.com/rbennum/service-account/utils/config"
)

var appConfig = config.GetConfig()

func Migrate() error {
	migrated, err := migrate.New(appConfig.MigrateFileLocation, appConfig.DBConnectionMigrate)
	if err != nil {
		return err
	}

	if err := migrated.Up(); err != nil {
		if err.Error() == "no change" {
		} else {
			return err
		}
	}
	return nil
}
