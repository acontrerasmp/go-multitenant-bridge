package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gopato.io/db"
	"gopato.io/models/public"
)

func GetAllTenants(c *fiber.Ctx) error {
	tenantCtx := c.Locals("tenant").(public.Tenant)
	fmt.Println(tenantCtx)
	db := db.Db
	var tenants []public.Tenant
	db.Limit(10).Find(&tenants)
	return c.Status(200).JSON(fiber.Map{"message": "all tenants", "data": tenants})

}
