package config_test

import (
	"testing"

	"battle-of-monsters/app/config"
	testsutils "battle-of-monsters/app/tests/utils"
)

func TestLoad(t *testing.T) {
	testsutils.LoadEnv()

	if config.ENV.Port != "8080" {
		t.Errorf("Port expected as 8080 but got %v", config.ENV.Port)
	}

	if config.ENV.DBDriver != "sqlite" {
		t.Errorf("Driver expected as sqlite but got %v", config.ENV.DBDriver)
	}

	if config.ENV.DBName != "db/db.test.sqlite" {
		t.Errorf("Database name expected as db/db.test.sqlite but got %v", config.ENV.DBName)
	}
}
