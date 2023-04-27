package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/pusher/pusher-http-go/v5"
)

type Message struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

func main() {

	r := fiber.New()
	r.Use(cors.New())

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ocurrió un error al cargar la variables de entorno.")
	}
	var (
		APP_ID  = os.Getenv("PUSHER_APP_ID")
		KEY     = os.Getenv("PUSHER_KEY")
		SECRET  = os.Getenv("PUSHER_SECRET")
		CLUSTER = os.Getenv("PUSHER_CLUSTER")
		SECURE  = true
	)

	pusherClient := pusher.Client{
		AppID:   APP_ID,
		Key:     KEY,
		Secret:  SECRET,
		Cluster: CLUSTER,
		Secure:  SECURE,
	}

	r.Post("/v1/messages", func(c *fiber.Ctx) error {
		data := map[string]Message{}
		if err := c.BodyParser(&data); err != nil {
			return err
		}
		pusherClient.Trigger("noxy-sysl", "message", data)
		return c.JSON(data)
	})
	r.Listen(":3000")
}
