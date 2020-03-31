package e2e

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/xavier268/go-demo-gin-gorm/internal/dao"
	"github.com/xavier268/go-demo-gin-gorm/internal/myapp"
)

// do will send url, compare body to expectedBody (regex),
// and compare status to expectedStatus.
// Will return the body as string for further testing.
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
