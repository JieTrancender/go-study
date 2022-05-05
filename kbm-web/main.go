package main

import (
	"log"
	"net/http"

	"kbm"
)

func main() {
	r := kbm.New()
	r.GET("/index", func(c *kbm.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *kbm.Context) {
			c.HTML(http.StatusOK, "<h1>Hello KBM </h1>")
		})

		v1.GET("/hello", func(c *kbm.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *kbm.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *kbm.Context) {
			c.JSON(http.StatusOK, kbm.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	r.GET("/", indexHandler)
	r.GET("/hello", helloHandler)
	r.GET("/hello/:name", func(c *kbm.Context) {
		c.String(http.StatusOK, "hello %s, you'are at %s\n", c.Param("name"), c.Path)
	})
	r.POST("/login", loginHandler)
	r.GET("/asserts/*filepath", func(c *kbm.Context) {
		c.JSON(http.StatusOK, kbm.H{"filepath": c.Param("filepath")})
	})

	log.Fatal(r.Run(":8888"))
}

func indexHandler(c *kbm.Context) {
	c.HTML(http.StatusOK, "<h>Hello Kbm</h1>")
}

func helloHandler(c *kbm.Context) {
	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
}

func loginHandler(c *kbm.Context) {
	c.JSON(http.StatusOK, kbm.H{
		"username": c.PostForm("username"),
		"password": c.PostForm("password"),
	})
}
