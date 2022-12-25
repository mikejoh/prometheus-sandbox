package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func defaultServe(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer req.Body.Close()
	var buf bytes.Buffer
	if err := json.Indent(&buf, b, " >", "  "); err != nil {
		log.Fatal(err)
	}
	log.Println(buf.String())
}

func responder(w http.ResponseWriter, req *http.Request) {
	values := req.URL.Query()

	code, err := strconv.Atoi(values["code"][0])
	if err != nil {
		log.Fatal(err)
	}
	body := values["body"][0]

	r := &http.Response{
		Status:     http.StatusText(code),
		StatusCode: code,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Body:       ioutil.NopCloser(bytes.NewBufferString(body)),
	}

	w.WriteHeader(code)

	if err := r.Write(w); err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/alertreceiver", defaultServe)
	http.HandleFunc("/responder", responder)
	err := http.ListenAndServe(":5001", nil)
	if err != nil {
		log.Fatal(err)
	}
}
