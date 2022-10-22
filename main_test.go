package main

import (
	"testing"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"fmt"
)

func TestHttp(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request){
		io.WriteString(w, "25.0")
	}

	req := httptest.NewRequest("GET", "http://localhost:6000/cpu/temp", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))
}
