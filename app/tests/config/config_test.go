package tests_config

import (
	"battle-of-monsters/app/config"
	tests_utils "battle-of-monsters/app/tests/utils"
	"testing"
)

func TestLoad(t *testing.T){

	tests_utils.LoadEnv()

	port := config.ENV.Port
	driver := config.ENV.DBDriver
	dbName := config.ENV.DBName

	if port != "8080" {
		t.Errorf("Port expected as 8080 but got %v", port)
	}

	if driver != "sqlite" {
		t.Errorf("Driver expected as sqlite but got %v", driver)
	}

	if dbName != "db/db.test.sqlite" {
		t.Errorf("Database name expected as db/db.test.sqlite but got %v", dbName)
	}

}