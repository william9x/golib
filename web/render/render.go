package render

import (
	"errors"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/william9x/golib-core/log"
)

// Renderer interface is to be implemented by JSON, XML, HTML, YAML and so on.
type Renderer interface {
	// Render writes data with custom ContentType.
	Render(*fiber.Ctx) error

	// WriteContentType writes custom ContentType.
	WriteContentType(ctx *fiber.Ctx)
}

// Render writes data with custom http status code
func Render(ctx *fiber.Ctx, httpStatus int, r Renderer) error {
	ctx.Status(httpStatus)
	if err := r.Render(ctx); err != nil {
		if !errors.Is(err, syscall.EPIPE) &&
			!errors.Is(err, syscall.ECONNRESET) {
			log.Errorf("Cannot render response with error [%v]", err)
		}
		return err
	}
	return nil
}

// writeContentType writes content type to a writer
func writeContentType(ctx *fiber.Ctx, value []string) {
	for _, v := range value {
		ctx.Response().Header.SetContentType(v)
	}
}
