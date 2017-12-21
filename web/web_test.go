package web

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeListen(t *testing.T) {
	s := &http.Server{
		Handler: http.HandlerFunc(ServeHTTP),
	}
	// Pick port automatically for parallel tests and to avoid conflicts
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()
	go s.Serve(l)

	res, err := http.Get("http://" + l.Addr().String() + "/?sloths=arecool")
	if err != nil {
		log.Fatal(err)
	}
	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(greeting))
}

func TestServeMemory(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/?sloths=arecool", nil)
	w := httptest.NewRecorder()

	ServeHTTP(w, req)
	greeting, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(greeting))
}
