package main

import (
	"context"
	"msg_queue/routes"
	"msg_queue/workers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	numWorkers := 5

	for i := 1; i <= numWorkers; i++ {
		go workers.StartWorkers(ctx, i)
	}
	app := fiber.New()
	routes.SetupRoutes(app)
	app.Listen(":3000")
}
