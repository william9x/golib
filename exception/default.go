package exception

import "github.com/valyala/fasthttp"

var (
	BadRequest   = New(fasthttp.StatusBadRequest, "Bad Request")
	Unauthorized = New(fasthttp.StatusUnauthorized, "Unauthorized")
	Forbidden    = New(fasthttp.StatusForbidden, "Forbidden")
	NotFound     = New(fasthttp.StatusNotFound, "Resource Not Found")
	SystemError  = New(fasthttp.StatusInternalServerError, "System Error")
)
