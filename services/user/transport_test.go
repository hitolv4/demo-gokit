package user

import (
	"bytes"
	"net/http"
	"testing"
)

func TestHTTPCreateUser(t *testing.T) {
	var jsonStr = []byte(`{"name":"testing http","email":"test@http.com,"password":"httppassword"}`)
	req, _ := http.NewRequest("POST", "/user/new", bytes.NewBuffer(jsonStr))
	req.Header.Set("content-type", "application/json")

}
func TestEmptyTable(t *testing.T) {

}
func TestHTTPGETAllUsers(t *testing.T) {}
func TestHTTPGETtUser(t *testing.T)    {}

func TestHTTPUpdateUser(t *testing.T) {}

func TestHTTPDeleteUser(t *testing.T) {}
