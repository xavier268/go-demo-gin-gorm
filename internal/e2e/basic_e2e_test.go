// Package e2e provides end-to-end testing
package e2e

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/xavier268/go-demo-gin-gorm/internal/dao"
	"github.com/xavier268/go-demo-gin-gorm/internal/models"
	"github.com/xavier268/go-demo-gin-gorm/internal/myapp"
)

func TestPing1Json(t *testing.T) {
	// Flag this as a parralel test ...
	// t.Parallel()

	body := do(t, "http://localhost:8080/v1/ping?json=1", http.StatusOK, "pong")

	res := new(gin.H)
	json.Unmarshal([]byte(body), res)

	if (*res)["Ping"] != "pong" {
		fmt.Printf("\n%+v", res)
		t.Fatal(res)
	}
}

func TestPing1XML(t *testing.T) {
	// Flag this as a parralel test ...
	// t.Parallel()

	body := do(t, "http://localhost:8080/v1/ping?xml=1", http.StatusOK, "<Ping>pong</Ping>")

	if string(body) != "<map><Ping>pong</Ping></map>" {
		log.Fatal(string(body))
		t.Fatal("Unexpected return : ", string(body), " instead of <map><Ping>pong</Ping></map>")
	}

}

func TestPing1HTML(t *testing.T) {
	// Flag this as a parralel test ...
	// t.Parallel()

	do(t, "http://localhost:8080/v1/ping", http.StatusOK, "<html>.*pong.*</html>")

}

func TestPing2Json(t *testing.T) {
	// Flag this as a parralel test ...
	// t.Parallel()

	body := do(t, "http://localhost:8080/v1/ping/test?json=1", http.StatusOK, "test")

	res := new(gin.H)
	json.Unmarshal([]byte(body), res)

	if (*res)["Ping"] != "test" {
		fmt.Printf("\n%+v", res)
		t.Fatal(res)
	}
}

func TestPing3Json(t *testing.T) {
	// Flag this as a parralel test ...
	// t.Parallel()

	body := do(t, "http://localhost:8080/v1/sleep?json=1", http.StatusOK, "sleep 1 second")

	res := new(gin.H)
	json.Unmarshal([]byte(body), res)
	if (*res)["Ping"] != "sleep 1 second" {
		fmt.Printf("\n%+v", res)
		t.Fatal(res)
	}
}

func TestProducts(t *testing.T) {

	// t.Parallel()

	ts.CreateProduct(11, "onze")
	ts.CreateProduct(12, "douze")
	ts.CreateProduct(13, "treize")
	ts.CreateProduct(14, "quatorze")

	fmt.Println(ts.AllProducts().ToString())

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

// do will send url, and compare body to expectedBody (regex),
// and status to expectedStatus. Will return the body as a string.
func do(t *testing.T, url string, expectedStatus int, expectedBody string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("ERROR : %v\n\turl \t %s\n", err, url)
		t.Fatal(err)
	}
	if resp.StatusCode != expectedStatus {
		fmt.Printf("ERROR : \n  expected code : %d\n        got code : %d\n", expectedStatus, resp.StatusCode)
		t.FailNow()
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	body = []byte(strings.ReplaceAll(string(body), "\n", ""))
	match, err := regexp.Match(expectedBody, body)
	if err != nil {
		t.Fatal("Eror trying to match body with regex : ", expectedBody, "\nerror : ", err)
	}
	if !match {
		fmt.Println("Body expected regex : ", expectedBody)
		fmt.Println("Body actual string  : ", string(body))
		t.Fatal("body did not match expectation")
	}
	return string(body)
}

// ts is the test data source (memory)
var ts *dao.Source

// TestMain is a wrapper around tests, that ensures server is started and then killed each time.
func TestMain(m *testing.M) {

	ts = dao.NewMemorySource()
	a := myapp.New(ts)

	go a.Run()
	e := m.Run()
	a.Shutdown()
	if ts.Close() != nil {
		panic("Error while closing DAO !?")
	}
	os.Exit(e)
}
