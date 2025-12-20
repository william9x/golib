package golib_core

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/william9x/golib-core/config"
	"github.com/william9x/golib-core/log"
	"github.com/william9x/golib-core/web/middleware"
	"go.uber.org/fx"
)

func AppOpt() fx.Option {
	return fx.Options(
		fx.Provide(func() context.Context {
			return context.Background()
		}),
		ProvideProps(config.NewAppProperties),
		fx.WithLogger(log.NewFxLogger),
		fx.Provide(New),
	)
}

func New(context context.Context, props *config.AppProperties) *App {
	app := App{context: context, props: props}
	app.AddHandler(
		middleware.Recover(),
		middleware.RequestId(),
	)
	return &app
}

type App struct {
	props    *config.AppProperties
	context  context.Context
	handlers []fiber.Handler
}

func (a *App) Name() string {
	return a.props.Name
}

func (a *App) Port() int {
	return a.props.Port
}

func (a *App) Path() string {
	return a.props.Path
}

func (a *App) Context() context.Context {
	return a.context
}

func (a *App) Handlers() []fiber.Handler {
	return a.handlers
}

func (a *App) AddHandler(handlers ...fiber.Handler) {
	a.handlers = append(a.handlers, handlers...)
}
