package main

import (
	"fmt"
	"log"
	"net/http"

	"kbm"
)

func main() {
	r := kbm.New()
	r.GET("/", indexHandler)
	r.GET("/hello", helloHandler)

	log.Fatal(r.Run(":8888"))
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.PATH= %q", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
