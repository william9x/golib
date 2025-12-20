package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"
	"github.com/william9x/golib-core/log"
	"github.com/william9x/golib-core/web/constant"
)

// RequestId middleware responsible to inject RequestId to request attributes
// RequestId is usually sent in the request header by the client (see constant.HeaderRequestId),
// but sometimes it doesn't exist, we will generate it automatically
func RequestId() func(ctx *fiber.Ctx) error {
	log.Infof("[Middleware] Registering RequestId middleware")
	return requestid.New(requestid.Config{
		Next:       nil,
		Header:     constant.HeaderRequestId,
		ContextKey: constant.UserValueRequestId,
		Generator: func() string {
			newRequestId, _ := uuid.NewV7()
			return newRequestId.String()
		},
	})
}
