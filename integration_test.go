package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoot(t *testing.T) {
	req, err := http.NewRequest("GET", "https://localhost:8080/dog", nil)
	if err != nil {
		t.Errorf("Exception creating new request: %s", err.Error())
	}
	fmt.Println(req.URL.Path)
	w := httptest.NewRecorder()
	LoveHandler(w, req)

	resp := w.Body
	body, _ := ioutil.ReadAll(resp)

	fmt.Println(string(body))
}
