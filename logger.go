package golib

import (
	"github.com/pkg/errors"
	"github.com/william9x/golib/log"
	"github.com/william9x/golib/web/middleware"
	"github.com/william9x/golib/web/properties"
	"go.uber.org/fx"
)

func LoggingOpt() fx.Option {
	return fx.Options(
		ProvideProps(log.NewProperties),
		fx.Provide(NewZapLogger),
		fx.Provide(log.NewFiberZapLogger),
		fx.Invoke(RegisterLogger),
		fx.Invoke(RegisterRequestContextMiddleware),
	)
}

type ZapLoggerIn struct {
	fx.In
	Props             *log.Properties
	ContextExtractors []log.ContextExtractor `group:"log_context_extractor"`
}

func NewZapLogger(in ZapLoggerIn) (*log.ZapLogger, error) {
	// Create new logger instance
	logger, err := log.NewZapLogger(&log.Options{
		Development:       in.Props.Development,
		LogLevel:          in.Props.LogLevel,
		JsonOutputMode:    in.Props.JsonOutputMode,
		DisableCaller:     in.Props.DisableCaller,
		DisableStacktrace: in.Props.DisableStacktrace,
		CallerSkip:        in.Props.CallerSkip,
		ContextExtractors: in.ContextExtractors,
	})
	if err != nil {
		return nil, errors.WithMessage(err, "init logger failed")
	}
	return logger, nil
}

func RegisterLogger(fiberZapLogger log.Logger) {
	log.ReplaceGlobal(fiberZapLogger)
}

func RegisterRequestContextMiddleware(
	app *App,
	zapLogger *log.ZapLogger,
	httpRequestLogProperties *properties.HttpRequestLogProperties,
) {
	app.AddHandler(middleware.RequestContext(zapLogger.GetCore(), httpRequestLogProperties))
}
