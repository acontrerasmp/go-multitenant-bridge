package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"gopato.io/handler"
	"gopato.io/middleware"
)

func SetupRoutes(app *fiber.App) {
	corsConfig := cors.New(cors.Config{
		AllowOrigins: "http://122.0.0.1, https://monpasse.com, http://127.0.0.1",
		AllowHeaders: "Origin, Content-Type, Accept",
	})
	// csrfConfig := csrf.New(csrf.Config{})
	api := app.Group(("api/"), logger.New(), corsConfig)
	api.Use(middleware.SetTenant())
	api.Get("monitor/", monitor.New())
	api.Post("createnant", handler.CreateTenant)
	api.Get("tenants/", handler.GetAllTenants)
}
