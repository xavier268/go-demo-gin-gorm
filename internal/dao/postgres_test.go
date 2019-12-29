package dao

import (
	"testing"
)

func TestPostgres(t *testing.T) {

	UseDriverPostgres()
	d := GetDAO()
	defer d.Close()

	if d.Dialect().GetName() != "postgres" {
		t.Log(d.Dialect().GetName())
		t.Fatal("Unexpected Dialect")
	}
}
