package config_test

import (
	"golang-test/config"
	"testing"
)

func TestMigration(t *testing.T) {
	config.LoadEnvFile()
	config.ConnectDB()
}
