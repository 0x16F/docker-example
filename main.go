package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ilyakaznacheev/cleanenv"
)

type ConfigDatabase struct {
	Port string `env:"PORT" env-default:"8080"`
}

func doRequest(c *fiber.Ctx) error {
	log.Println("do request")

	t := time.Now()

	resp, err := http.Get("https://google.com")
	if err != nil {
		return err
	}

	sinceResult := time.Since(t)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"size": len(body),
		"ping": sinceResult.Milliseconds(),
	})
}

func main() {
	cfg := ConfigDatabase{}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/", doRequest)

	app.Listen(":" + cfg.Port)
}
