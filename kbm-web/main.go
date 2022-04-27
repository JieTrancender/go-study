package main

import (
	"log"
	"net/http"

	"kbm"
)

func main() {
	engine := new(kbm.Engine)
	log.Fatal(http.ListenAndServe("localhost:8000", engine))
}
