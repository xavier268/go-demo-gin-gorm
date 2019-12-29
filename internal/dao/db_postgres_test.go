package dao

import (
	"testing"
)

func TestPostgres(t *testing.T) {

	d := NewPostgresSource()
	defer d.Close()

	if d.GetDB().Dialect().GetName() != "postgres" {
		t.Log(d.GetDB().Dialect().GetName())
		t.Fatal("Unexpected Dialect")
	}
}
