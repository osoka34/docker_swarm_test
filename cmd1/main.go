package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func main() {
	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		// Непрерывно отправлять "pong" каждую секунду
		//go func() {
		//for {
		//c.Context().Response.Reset()
		time.Sleep(1 * time.Second)
		err := c.SendString("pong 1\n")
		if err != nil {
			return err
		}
		//c.Context().Response.Reset() // Очистить буфер перед отправкой следующего "pong"
		//time.Sleep(1 * time.Second)
		//}
		//}()
		return nil
	})

	port := 30000
	fmt.Printf("Сервер запущен на порту %d\n", port)

	err := app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
}
