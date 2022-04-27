package main

import (
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

func indexHandler(c *kbm.Context) {
	c.HTML(http.StatusOK, "<h>Hello Kbm</h1>")
}

func helloHandler(c *kbm.Context) {
	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
}
