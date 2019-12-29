package dao

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDB(t *testing.T) {

	dao := ts.GetDAO()

	if dao.Dialect().GetName() != "sqlite3" && dao.Dialect().GetName() != "postgres" {
		t.Log(dao.Dialect().GetName())
		t.Fatal("Unexpected Dialect")
	}

}
func TestCount(t *testing.T) {

	// initial count ...
	c := ts.CountProducts()

	if count := ts.CountProducts(); count != c {
		t.Fatalf("Unexpected count of products : %d", count)
	}

	ts.CreateProduct(100, "cent")

	if count := ts.CountProducts(); count != c+1 {
		t.Fatalf("Unexpected count of products : %d", count)
	}

	deux := ts.CreateProduct(200, "deux cents")
	ts.CreateProduct(300, "trois cents")

	fmt.Println(ts.AllProducts().ToString())
	if count := ts.CountProducts(); count != c+3 {
		t.Fatalf("Unexpected count of products : %d", count)
	}
	ts.DeleteProduct(deux)
	fmt.Println(ts.AllProducts().ToString())
	if count := ts.CountProducts(); count != c+2 {
		t.Fatalf("Unexpected count of products : %d", count)
	}
	ts.DeleteProducts()
	if count := ts.CountProducts(); count != 0 {
		t.Fatalf("Unexpected count of products : %d", count)
	}

}

//====================================

// ts is the data source for this test
var ts *Source = NewMemorySource()

// Ensure closing db after all tests are performed.
func TestMain(m *testing.M) {

	e := m.Run()
	if ts.GetDAO().Close() != nil {
		panic("Error while closing DAO !?")
	}
	os.Exit(e)
}
