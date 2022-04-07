package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gopato.io/db"
	"gopato.io/models/public"
)

func SetTenant() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var tenant public.Tenant

		if c.IsFromLocal() {
			c.Locals("tenant", tenant)
			return c.Next()
		}
		db := db.Db
		// userdb := c.Locals("userdb").(models.Account)
		tenantDomain := c.Hostname()
		db.Where("domain = ?", tenantDomain).Find(&tenant)
		c.Locals("tenant", tenant)
		fmt.Println(tenant)
		fmt.Println(tenantDomain)
		// context.WithValue()
		return c.Next()
	}
}
