package response

import (
	"errors"

	"github.com/valyala/fasthttp"
	"github.com/william9x/golib/exception"
)

type Response struct {
	Meta Meta        `json:"meta,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func New(code int, message string, data interface{}) Response {
	return Response{
		Meta: Meta{
			Code:    code,
			Message: message,
		},
		Data: data,
	}
}

func Ok(data interface{}) Response {
	return New(fasthttp.StatusOK, "Successful", data)
}

func Created(data interface{}) Response {
	return New(fasthttp.StatusCreated, "Resource created", data)
}

func Error(err error) Response {
	code := fasthttp.StatusInternalServerError
	message := "Internal Server Error"
	if e, ok := errors.AsType[exception.Exception](err); ok {
		code = int(e.Code())
		message = e.Message()
	}
	return Response{
		Meta: Meta{
			Code:    code,
			Message: message,
		},
	}
}
