package thot

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// ForwardRequest ...
func ForwardRequest(request Subscriber) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		uri := fmt.Sprintf("%s/%s", request.URL, request.Endpoint)

		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)
			fatal(err)
			fmt.Println(r.Method, uri)
			rr, err := http.NewRequest(r.Method, uri, bytes.NewBuffer(body))
			fatal(err)

			rr.Header.Add("Content-Type", "application/json")
			rr.Header.Add("Content-Length", strconv.Itoa(len(body)))

			client := &http.Client{}
			resp, err := client.Do(rr)
			fatal(err)

			defer resp.Body.Close()
			boby, err := ioutil.ReadAll(resp.Body)
			fatal(err)

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.Write(boby)
		} else if r.Method == "GET" {
			resp, err := http.Get(uri)
			fatal(err)

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			fatal(err)

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.Write(body)
		}
	}
}
