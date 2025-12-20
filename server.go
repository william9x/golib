package golib_core

import (
	"context"
	"fmt"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/william9x/golib-core/log"
	"github.com/william9x/golib-core/web/properties"
	"go.uber.org/fx"
)

func FiberServerOpt() fx.Option {
	return fx.Options(
		ProvideProps(properties.NewFiberProperties),
		ProvideProps(properties.NewHttpRequestLogProperties),

		fx.Provide(NewFiberServer),
		fx.Invoke(RegisterHandlers),
		fx.Invoke(OnStartHttpServerHook),
	)
}

func OnStopHttpServerOpt() fx.Option {
	return fx.Invoke(OnStopHttpServerHook)
}

func NewFiberServer(config *properties.FiberProperties, app *App) *fiber.App {
	return fiber.New(fiber.Config{
		AppName:                      app.Name(),
		Prefork:                      config.Prefork,
		StrictRouting:                config.StrictRouting,
		CaseSensitive:                config.CaseSensitive,
		Immutable:                    config.Immutable,
		UnescapePath:                 config.UnescapePath,
		BodyLimit:                    config.BodyLimit,
		Concurrency:                  config.Concurrency,
		ReadTimeout:                  config.ReadTimeout,
		WriteTimeout:                 config.WriteTimeout,
		IdleTimeout:                  config.IdleTimeout,
		ReadBufferSize:               config.ReadBufferSize,
		WriteBufferSize:              config.WriteBufferSize,
		CompressedFileSuffix:         config.CompressedFileSuffix,
		ProxyHeader:                  config.ProxyHeader,
		GETOnly:                      config.GETOnly,
		DisableKeepalive:             config.DisableKeepalive,
		DisableDefaultDate:           config.DisableDefaultDate,
		DisableDefaultContentType:    config.DisableDefaultContentType,
		DisableHeaderNormalizing:     config.DisableHeaderNormalizing,
		DisableStartupMessage:        config.DisableStartupMessage,
		StreamRequestBody:            config.StreamRequestBody,
		DisablePreParseMultipartForm: config.DisablePreParseMultipartForm,
		Network:                      config.Network,
		EnableTrustedProxyCheck:      config.EnableTrustedProxyCheck,
		TrustedProxies:               config.TrustedProxies,
		EnableIPValidation:           config.EnableIPValidation,
		EnablePrintRoutes:            config.EnablePrintRoutes,
		RequestMethods:               config.RequestMethods,
		EnableSplittingOnParsers:     config.EnableSplittingOnParsers,

		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})
}

func RegisterHandlers(app *App, engine *fiber.App) {
	for _, handler := range app.Handlers() {
		engine.Use(handler)
	}
}

func OnStartHttpServerHook(lc fx.Lifecycle, app *App, fiberServer *fiber.App, fiberConfig *properties.FiberProperties) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Infof("Application will be served at port %d. Service name: %s, service path: %s",
				app.Port(), app.Name(), app.Path())
			go func() {
				if fiberConfig.Tls.Enabled {
					if err := fiberServer.ListenTLS(fmt.Sprintf(":%d", app.Port()), fiberConfig.Tls.CertFile, fiberConfig.Tls.KeyFile); err != nil {
						log.Errorf("Could not serve HTTP request, error [%v]", err)
					}
				} else {
					if err := fiberServer.Listen(fmt.Sprintf(":%d", app.Port())); err != nil {
						log.Errorf("Could not serve HTTPS request, error [%v]", err)
					}
				}
				log.Infof("Stopped HTTP Server")
			}()
			return nil
		},
	})
}

func OnStopHttpServerHook(lc fx.Lifecycle, server *fiber.App) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			log.Infof("Stopping HTTP Server")
			if err := server.ShutdownWithContext(ctx); err != nil {
				log.Errorf("Could not stop HTTP server, error [%v]", err)
			}
			return nil
		},
	})
}
