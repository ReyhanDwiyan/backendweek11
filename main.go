package main

import (
	"fmt"
	"inibackend/config"
	"inibackend/router"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading environtment variables:")
	}
}

func main() {
	app := fiber.New()

	//login req terminal
	app.Use(logger.New())

	// basic cors
	app.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Join(config.GetAllowedOrigin(), ","),
		AllowCredentials: true,
		AllowMethods:     "GET,POST,PUT,DELETE",
	}))

	//route mahasiswa
	router.SetupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "Endpoint tidak ditemukan",
		})
	})

	// Baca Port
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// cek log koneksi di port
	log.Printf("Server running on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error starting server: &v", err)
	}

}
