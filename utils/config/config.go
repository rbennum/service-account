package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	instance *Configuration
	once     sync.Once
)

type Configuration struct {
	Port                string
	DBConnection        string
	DBConnectionMigrate string
	MigrateFileLocation string
}

func GetConfig() *Configuration {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			return
		}
		instance = &Configuration{
			Port:                getEnv("PORT", "8080"),
			DBConnection:        getDBConnection(),
			DBConnectionMigrate: getDBConnectionMigrate(),
			MigrateFileLocation: getLocationMigrate(),
		}
	})
	return instance
}

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func getDBConnection() string {
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASS", "")
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbName := getEnv("DB_NAME", "postgres")
	databaseurl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	return databaseurl
}

func getDBConnectionMigrate() string {
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "")
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbName := getEnv("DB_NAME", "postgres")
	databaseurl := fmt.Sprintf("pgx://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	return databaseurl
}

func getLocationMigrate() string {
	return "file://" + getEnv("MIGRATION_FILE_PATH", "src/database/migrations")
}
