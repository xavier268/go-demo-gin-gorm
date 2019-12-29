package dao

import (
	"testing"
)

func TestPostgres(t *testing.T) {

	d := NewPostgresSource()
	defer d.Close()

	if d.GetDAO().Dialect().GetName() != "postgres" {
		t.Log(d.GetDAO().Dialect().GetName())
		t.Fatal("Unexpected Dialect")
	}
}
