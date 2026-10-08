package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	var keys []string
	app.Post("/x", func(c fiber.Ctx) error {
		for key := range c.Request().Header.All() {
			keys = append(keys, string(key))
		}
		fmt.Println("keys Fiber gives us:", keys)
		return c.SendStatus(200)
	})
	req := httptest.NewRequest("POST", "/x", nil)
	req.Header.Add("x-tapes-agent-name", "agent-a")
	req.Header.Add("x-request-id", "r1")
	req.Header.Add("Cookie", "a=1")
	req.Header.Add("cookie", "b=2")
	resp, _ := app.Test(req)
	resp.Body.Close()
}