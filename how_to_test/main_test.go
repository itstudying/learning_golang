package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestHelloHandleFunc(t *testing.T) {
	rw := httptest.NewRecorder()
	name := "zouying"
	req := httptest.NewRequest(http.MethodPost, "/hello?name="+name, nil)
	handleHello(rw, req)

	if rw.Code != http.StatusOK {
		t.Errorf("status code not ok, status code is %v", rw.Code)
	}

	if len(counter) != 1 {
		t.Errorf("counter len not correct")
	}

	if counter[name] != 1 {
		t.Errorf("counter value is error: visitor=%s count=%v", name, counter[name])
	}
}

func TestHTTPServer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handleHello))
	defer ts.Close()

	logrus.Infof("server url: %s", ts.URL)

}
