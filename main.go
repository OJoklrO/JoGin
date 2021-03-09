package main

import (
	"Jin"
)

func main() {
	j := Jin.New()
	j.Use(Jin.Logger())

	j.Static("/assets", "./static")

	v1 := j.Group("/v1")
	v1.GET("/", func(c *Jin.Context) {
		c.String(200, "hello jin")
	})
	v1.GET("/hello", func(c *Jin.Context) {
		c.String(200, "hello %s\n", c.Query("name"))
	})

	v2 := j.Group("/v2")
	v2.GET("/hello/:name", func(c *Jin.Context) {
		c.String(200, "hello %s\n", c.Param("name"))
	})
	v2.POST("/login", func(c *Jin.Context) {
		c.JSON(200, Jin.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	_ = j.Run(":8080")
}
