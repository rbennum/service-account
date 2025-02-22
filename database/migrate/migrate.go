package migrate

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rbennum/service-account/utils/config"
)

var appConfig = config.GetConfig()

func Migrate() error {
	migrated, err := migrate.New(appConfig.MigrateFileLocation, appConfig.DBConnection)
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
