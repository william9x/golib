package actuator

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"github.com/william9x/golib/actuator"
	"github.com/william9x/golib/web/response"
)

type Endpoint struct {
	healthService actuator.HealthService
	infoService   actuator.InfoService
}

func NewEndpoint(healthService actuator.HealthService, infoService actuator.InfoService) *Endpoint {
	return &Endpoint{
		healthService: healthService,
		infoService:   infoService,
	}
}

func (c Endpoint) HealthService() actuator.HealthService {
	return c.healthService
}

func (c Endpoint) InfoService() actuator.InfoService {
	return c.infoService
}

func (c Endpoint) Health(ctx *fiber.Ctx) error {
	health := c.healthService.Check(ctx.UserContext())
	var res response.Response
	if health.Status == actuator.StatusDown {
		res = response.New(fasthttp.StatusServiceUnavailable, "Server is down", health)
	} else {
		res = response.New(fasthttp.StatusOK, "Server is up", health)
	}
	return response.Write(ctx, res)
}

func (c Endpoint) Info(ctx *fiber.Ctx) error {
	info := c.infoService.Info()
	return response.Write(ctx, response.Ok(info))
}
