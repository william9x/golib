package middleware

import (
	"fmt"

	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/william9x/golib/log"
	"github.com/william9x/golib/web/constant"
	"github.com/william9x/golib/web/properties"
	"go.uber.org/zap"
)

// RequestContext middleware for Fiber that logs HTTP request/response details.
func RequestContext(
	zapLogger *zap.Logger,
	httpRequestLogProperties *properties.HttpRequestLogProperties,
) func(ctx *fiber.Ctx) error {
	log.Infof("[Middleware] Registering RequestContext middleware")
	if httpRequestLogProperties.Disabled {
		log.Infof("[Middleware] RequestContext middleware is disabled")
		return func(ctx *fiber.Ctx) error {
			return ctx.Next()
		}
	}

	allDisabledUrls := httpRequestLogProperties.AllDisabledUrls()
	allDisabledUrlsStr := make([]string, len(allDisabledUrls))
	for i := 0; i <= len(allDisabledUrls)-1; i++ {
		allDisabledUrlsStr[i] = allDisabledUrls[i].UrlPattern
	}
	log.Infof("[Middleware] RequestContext middleware is enabled, disabled URLs: %v", allDisabledUrlsStr)

	return fiberzap.New(fiberzap.Config{
		SkipURIs: allDisabledUrlsStr,
		Logger:   zapLogger,
		Fields: []string{
			"status",
			"latency",
			"ip",
			"ua",
			"path",
			"method",
			"queryParams",
		},
		// TODO: provide a way to customize fields without modifying this library
		FieldsFunc: func(c *fiber.Ctx) []zap.Field {
			return []zap.Field{
				zap.String("request_id", fmt.Sprintf("%s", c.Request().UserValue(constant.UserValueRequestId))),
				zap.String("platform", c.Get(constant.HeaderPlatform)),
				zap.String("on_mobile", c.Get(constant.HeaderIsOnMobile)),
			}
		},
	})
}
