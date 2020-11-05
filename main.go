package main

import (
	"gee"
	"net/http"
)

func main() {
	myGee := gee.NEW()

	myGee.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>hello ramon<h1>")
	})
	myGee.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello, %s", c.Query("name"))
	})
	myGee.POST("/login", func(c *gee.Context) {

		c.JSON(http.StatusOK, gee.H{
			"username": c.PostFrom("username"),
			"password": c.PostFrom("password"),
		})
	})

	myGee.RUN()
}
