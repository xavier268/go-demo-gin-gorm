package dao

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDB(t *testing.T) {

	dao := GetDAO()

	if dao.Dialect().GetName() != "sqlite3" {
		t.Log(dao.Dialect().GetName())
		t.Fatal("Unexpected Dialect")
	}

}
func TestCount(t *testing.T) {

	dao := GetDAO()

	if dao.Dialect().GetName() != "sqlite3" {
		t.Log(dao.Dialect().GetName())
		t.Fatal("Unexpected Dialect")
	}

	if count := dao.CountProducts(); count != 0 {
		t.Fatalf("Unexpected count of products : %d", count)
	}

	dao.CreateProduct(100, "cent")

	if count := dao.CountProducts(); count != 1 {
		t.Fatalf("Unexpected count of products : %d", count)
	}

	deux := dao.CreateProduct(200, "deux cents")
	dao.CreateProduct(300, "trois cents")

	fmt.Println(dao.AllProducts().ToString())
	if count := dao.CountProducts(); count != 3 {
		t.Fatalf("Unexpected count of products : %d", count)
	}
	dao.DeleteProduct(deux)
	fmt.Println(dao.AllProducts().ToString())
	if count := dao.CountProducts(); count != 2 {
		t.Fatalf("Unexpected count of products : %d", count)
	}
	dao.DeleteProducts()
	if count := dao.CountProducts(); count != 0 {
		t.Fatalf("Unexpected count of products : %d", count)
	}

}

//====================================

// Ensure closing db after all tests are performed.
func TestMain(m *testing.M) {
	e := m.Run()
	if GetDAO().Close() != nil {
		panic("Error while closing DAO !?")
	}
	os.Exit(e)
}
