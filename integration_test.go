package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ts *httptest.Server

func TestMain(m *testing.M) {
	ts = httptest.NewServer(http.HandlerFunc(loveHandler))
	defer ts.Close()
	m.Run()
}

func TestServer(t *testing.T) {
	resp, err := http.Get(ts.URL + "/dog")
	assert.Nil(t, err)

	message, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	assert.Nil(t, err)

	assert.Equal(t, "Hi there, I love dog!", string(message))
}
