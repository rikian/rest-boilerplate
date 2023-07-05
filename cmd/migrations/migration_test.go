package migrations_test

import (
	"testing"

	"golang-test/cmd/migrations"
	"golang-test/config"
)

func TestMigration(t *testing.T) {
	config.LoadEnvFile()
	db := config.ConnectDB()
	migrations.RunMigration(db)
}
