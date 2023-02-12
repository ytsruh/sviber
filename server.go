package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sviber/routes"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout: 5 * time.Second,
	})

	// Setup middleware
	app.Use(compress.New())
	app.Use(helmet.New())
	app.Use(recover.New())

	//Define API routes & set static path for client
	routes.SetRoutes((app))
	routes.SetHooks((app))
	app.Static("/", "./public")

	//Start server with graceful shutdown
	// Listen from goroutine
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	<-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	fmt.Println("Fiber was successful shutdown.")
}
