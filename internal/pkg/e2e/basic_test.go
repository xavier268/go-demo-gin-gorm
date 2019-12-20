// Package e2e provides end-to-end testing
package e2e

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/xavier268/go-demo-gin-gorm/internal/pkg/dao"
	"github.com/xavier268/go-demo-gin-gorm/internal/pkg/models"
	"github.com/xavier268/go-demo-gin-gorm/internal/pkg/myapp"
)

func TestPing1(t *testing.T) {
	// Flag this as a parralel test ...
	t.Parallel()

	resp, err := http.Get("http://localhost:8080/v1/ping")
	if err != nil || resp.StatusCode != http.StatusOK {
		t.Log("Error : ", err)
		t.Log("Resp  : ", resp)
		t.Fail()
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Log(body)
		t.Fatal(err)
	}
	res := new(models.Ping)
	json.Unmarshal([]byte(body), res)
	fmt.Printf("\n%+v", res)
	if res.Ping != "Pong" {
		t.Fatal(res)
	}

	fmt.Println()
}

func TestPing2(t *testing.T) {
	// Flag this as a parralel test ...
	t.Parallel()

	resp, err := http.Get("http://localhost:8080/v1/ping/test")
	if err != nil || resp.StatusCode != http.StatusOK {
		t.Log("Error : ", err)
		t.Log("Resp  : ", resp)
		t.Fail()
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Log(body)
		t.Fatal(err)
	}
	res := new(models.Ping)
	json.Unmarshal([]byte(body), res)
	fmt.Printf("\n%+v", res)
	if res.Ping != "test" {
		t.Fatal(res)
	}

	fmt.Println()
}

func TestPing3(t *testing.T) {
	// Flag this as a parralel test ...
	t.Parallel()

	resp, err := http.Get("http://localhost:8080/v1/sleep")
	if err != nil || resp.StatusCode != http.StatusOK {
		t.Log("Error : ", err)
		t.Log("Resp  : ", resp)
		t.Fail()
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Log(body)
		t.Fatal(err)
	}
	res := new(models.Ping)
	json.Unmarshal([]byte(body), res)
	fmt.Printf("\n%+v", res)
	if res.Ping != "Sleep" {
		t.Fatal(res)
	}

	fmt.Println()
}

func TestProducts(t *testing.T) {

	t.Parallel()

	dao := dao.GetDAO()

	dao.CreateProduct(11, "onze")
	dao.CreateProduct(12, "douze")
	dao.CreateProduct(13, "treize")
	dao.CreateProduct(14, "quatorze")

	fmt.Println(dao.AllProducts().ToString())

	resp, err := http.Get("http://localhost:8080/v1/p/2")
	if err != nil || resp.StatusCode != http.StatusOK {
		t.Log("Error : ", err)
		t.Log("Resp  : ", resp)
		t.Fail()
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Log(body)
		t.Fatal(err)
	}
	res := new(models.Product)
	json.Unmarshal([]byte(body), res)
	fmt.Printf("\n%+v", res.ToString())

}

// ====================================================

// TestMain is a wrapper around tests, that ensures server is started and then killed ecah time.
func TestMain(m *testing.M) {
	a := myapp.New()
	go a.Run()
	e := m.Run()
	a.Shutdown()
	if dao.GetDAO().Close() != nil {
		panic("Error while closing DAO !?")
	}
	os.Exit(e)
}
