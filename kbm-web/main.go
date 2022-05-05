package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"kbm"
)

func onlyForV2() kbm.HandlerFunc {
	return func(c *kbm.Context) {
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := kbm.Default()
	r.Use(kbm.Logger())
	r.GET("/index", func(c *kbm.Context) {
		c.HTML(http.StatusOK, "kbm.css", "<h1>Index Page</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *kbm.Context) {
			c.HTML(http.StatusOK, "kbm.css", "<h1>Hello KBM </h1>")
		})

		v1.GET("/hello", func(c *kbm.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
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

	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	r.GET("/panic", func(c *kbm.Context) {
		names := []string{"Kbm"}
		c.String(http.StatusOK, names[100])
	})

	log.Fatal(r.Run(":8888"))
}

func indexHandler(c *kbm.Context) {
	c.HTML(http.StatusOK, "css.tmpl", nil)
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
