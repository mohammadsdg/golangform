package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./views" , ".html")

	app:= fiber.New(fiber.Config{
		Views: engine,
	})
	app.Get("/" ,func(c *fiber.Ctx) error {
		return c.Render ("index" , fiber.Map{
			"Title" : "hello html ...", 
		})
	})

	app.Get("/index" ,func(c *fiber.Ctx) error {
		return c.Render ("index" ,fiber.Map{
			"Title" : "hello index..." ,
		})
	})

	app.Get("/call" ,func(c *fiber.Ctx) error {
		return c.Render ("index" , fiber.Map{
			"Title" : "calling ..." ,
			"Label" : "user name :" ,
		})
	})
	log.Fatal(app.Listen(":3000"))
}