package main

import (
	"nanokv/kvstore"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	kv := kvstore.NewKeyValueStore()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("This is a simple KV store like Redis in Go.")
	})

	app.Get("/get/:key", func(c *fiber.Ctx) error {
		key := c.Params("key")
		value, ok := kv.Get(key)
		if !ok {
			return c.SendString("The key " + key + " doesn't exist!")
		}
		return c.SendString("The key " + key + " has value " + value + ".")
	})

	app.Post("/set/:key/:value", func(c *fiber.Ctx) error {
		key := c.Params("key")
		value := c.Params("value")
		kv.Set(key, value, 10*time.Minute)
		return c.SendString("Key " + key + " value " + value + ".")
	})

	app.Delete("/delete/:key", func(c *fiber.Ctx) error {
		key := c.Params("key")
		ok := kv.Delete(key)
		if !ok {
			return c.SendString("The key " + key + " doesn't exist!")
		}

		return c.SendString("Successfully deleted!!!")
	})

	app.Listen(":3456")
}
