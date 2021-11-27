package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
)

func homePage(c *fiber.Ctx) error {
	return c.Render("pages/home", fiber.Map{
		"Title": "Home",
		"User":  c.Locals("SESSION_KEY_USER_ID"),
	})
}

func place(c *fiber.Ctx) error {
	return c.Render("pages/place", fiber.Map{
		"Title": "Antalya",
		"User":  c.Locals("SESSION_KEY_USER_ID"),
	})
}

func main() {
	engine := html.New("./views", ".html")

	// TODO: delete later
	engine.Reload(true)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(compress.New())
	app.Use(recover.New())

	app.Get("/", homePage)
	app.Get("/place", place)

	app.Listen(":3000")
}
