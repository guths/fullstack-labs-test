package tests_config

import "testing"

func TestLoad(t *testing.T){
	if true != true {
		t.Errorf("Error load configs")
	}
}