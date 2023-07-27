package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/spf13/pflag"
)

func main() {
	port := pflag.IntP("port", "p", 7270, "Port of Hachidori")
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
		//ViewsLayout: "main",
	})
	app.Static("/", "./assets")

	initAPI(app)
	exit := make(chan bool)

	fmt.Println(app.Listen(fmt.Sprintf(":%d", *port)))
	fmt.Printf("Listening on port %d\n", *port)
	<-exit
}

type nav struct{
	Name string
	URL string
}

func render(c *fiber.Ctx, tmpl string, fm fiber.Map) error {
	mapP := fm
	mapP["Nav"] = []nav{
		{
			Name: "Home",
			URL: "/",
		},
	}

	return c.Render(tmpl, mapP)
}
