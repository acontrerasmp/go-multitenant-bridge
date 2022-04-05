package handler

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gopato.io/db"
	"gopato.io/models/public"
	"gopato.io/models/shared"
)

func CreateTenant(c *fiber.Ctx) error {
	db := db.Db
	tenant := new(public.Tenant)

	if err := c.BodyParser(tenant); err != nil {
		// sentry.CaptureException(err)
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{"message": "error parsign the body request", "data": err},
		)
	}
	if err := db.Create(&tenant).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "couldn't create the tenant", "data": err})
	}
	tenant.ID = uuid.New()
	schema_name := strings.Replace(tenant.Domain, ".", "_", -1)
	db.Exec(fmt.Sprintf("CREATE SCHEMA %s", schema_name))
	db.Exec(fmt.Sprintf("SET SEARCH_PATH TO %s,public", schema_name))
	db.AutoMigrate(&shared.Product{})
	return c.Status(200).JSON(fiber.Map{"message": "tenant created", "data": tenant})
}

// func DeleteTenant(c *fiber.Ctx) error {
// 	db := db.Db
// 	tenant := new(public.Tenant)

// }
