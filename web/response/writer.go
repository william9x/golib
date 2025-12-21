package response

import (
	"github.com/gofiber/fiber/v2"
	"github.com/william9x/golib/web/render"
)

func Write(ctx *fiber.Ctx, res Response) error {
	return render.Render(ctx, res.Meta.HttpStatus(), render.JSON{Data: res})
}

// Write2 is an alias of Write, without data & meta in response body
func Write2(ctx *fiber.Ctx, res Response) error {
	return render.Render(ctx, res.Meta.HttpStatus(), render.JSON{Data: res.Data})
}

func WriteError(ctx *fiber.Ctx, err error) error {
	return Write(ctx, Error(err))
}
