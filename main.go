package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	b, err := ioutil.ReadFile("/index.html")
	if err != nil {
		panic(err)
	}
	err = http.ListenAndServe("0.0.0.0:9000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(b)
	}))
	if err != nil {
		panic(err)
	}
}
