package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World")
	})

	r.Run(":3000")
}

// package main

// import (
// 	"github.com/gofiber/fiber/v2"
// )

// func main() {
// 	app := fiber.New()

// 	app.Get("/", func(c *fiber.Ctx) error {
// 		return c.SendString("Hello, World")
// 	})

// 	app.Listen(":3000")
// }
