package render

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
)

var jsonContentType = []string{"application/json; charset=utf-8"}

// JSON contains the given interface object.
type JSON struct {
	Data interface{}
}

// Render (JSON) writes data with custom ContentType.
func (r JSON) Render(ctx *fiber.Ctx) (err error) {
	if err = WriteJSON(ctx, r.Data); err != nil {
		return err
	}
	return
}

// WriteContentType (JSON) writes JSON ContentType.
func (r JSON) WriteContentType(ctx *fiber.Ctx) {
	writeContentType(ctx, jsonContentType)
}

// WriteJSON marshals the given interface object and writes it with custom ContentType.
func WriteJSON(ctx *fiber.Ctx, obj interface{}) error {
	writeContentType(ctx, jsonContentType)
	jsonBytes, err := sonic.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = ctx.Write(jsonBytes)
	return err
}
