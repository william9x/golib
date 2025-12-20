package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/william9x/golib-core/log"
)

// Recover middleware for Fiber that recovers from panics anywhere in the stack chain
// and handles the control to the centralized ErrorHandler.
func Recover() func(ctx *fiber.Ctx) error {
	log.Infof("[Middleware] Registering Recover middleware")
	return recover.New()
}
