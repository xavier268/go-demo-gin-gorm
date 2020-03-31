// Package e2e provides end-to-end testing
package e2e

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
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
	if id := ts.CreateProduct(12, "douze"); id != 2 {
		t.Fatal("Expected ID of 2, but got ", id)
	}
	ts.CreateProduct(13, "treize")
	ts.CreateProduct(14, "quatorze")

	fmt.Println(ts.AllProducts().ToString())

	do(t, "http://localhost:8080/v1/p/2", http.StatusOK, "^.*\\{.*douze.*$")

}
