package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"gopato.io/db"
	"gopato.io/router"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("cant load venv")
	}
	prod, _ := strconv.ParseBool(os.Getenv("PROD"))
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	app := fiber.New(
		fiber.Config{
			Prefork: prod,
			AppName: "mp_auth",
		},
	)
	app.Use(cors.New())
	router.SetupRoutes(app)
	db.ConnectDb()
	// config.SetupAws()
	log.Fatal(app.Listen(":8080"))
}
