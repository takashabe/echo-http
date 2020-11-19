package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// DefaultPort default listen port
const DefaultPort = "8080"

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	p := os.Getenv("ECHO_PORT")
	if p == "" {
		p = DefaultPort
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", p), nil))
}

type response struct {
	URL     string      `json:"url"`
	Method  string      `json:"method"`
	Headers http.Header `json:"headers"`
	Body    []byte      `json:"body"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	resp := &response{
		URL:     r.URL.String(),
		Method:  r.Method,
		Headers: r.Header,
		Body:    b,
	}

	ret, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(ret)
	w.WriteHeader(200)
}
