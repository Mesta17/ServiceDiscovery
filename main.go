package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"net/http/httptest"

	"github.com/Mesta17/ServiceDiscovery/xlog"
)

const debugMode = false

type handler func(http.ResponseWriter, *http.Request)

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	xlog.Infof("\n\tMethod:%v \n\tURL:%v", r.Method, r.URL.Path)
	if debugMode {
		h.debugServeHTTP(w, r)
	} else {
		h(w, r)
	}
}

func (h handler) debugServeHTTP(w http.ResponseWriter, r *http.Request) {
	recorder := httptest.NewRecorder()
	h(recorder, r)
	response := recorder.Body
	body, _ := ioutil.ReadAll(response)
	xlog.Verbosef("\n\tResponse:%v", string(body))
	fmt.Fprintf(w, string(body)+"\n")
}

func loveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", handler(loveHandler))
	http.ListenAndServe(":8080", mux)
}
