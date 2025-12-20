package response

import (
	"github.com/gofiber/fiber/v2"
	"github.com/william9x/golib-core/web/render"
)

func Write(ctx *fiber.Ctx, res Response) error {
	return render.Render(ctx, res.Meta.HttpStatus(), render.JSON{Data: res})
}

func WriteError(ctx *fiber.Ctx, err error) error {
	return Write(ctx, Error(err))
}
