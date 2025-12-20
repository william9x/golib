package log

import (
	"context"

	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/william9x/golib-core/log/field"
)

type FiberZapLogger struct {
	logger log.AllLogger
}

func NewFiberZapLogger(zapLogger *ZapLogger) (Logger, error) {
	fiberZapLogger := fiberzap.NewLogger(fiberzap.LoggerConfig{
		SetLogger: zapLogger.GetCore(),
	})
	defer fiberZapLogger.Sync()

	return &FiberZapLogger{
		logger: fiberZapLogger,
	}, nil
}

func (l *FiberZapLogger) WithCtx(ctx context.Context, additionalFields ...field.Field) Logger {
	return l
}

func (l *FiberZapLogger) WithField(fields ...field.Field) Logger {
	return l
}

func (l *FiberZapLogger) WithError(err error) Logger {
	return l
}

func (l *FiberZapLogger) WithErrors(errs ...error) Logger {
	return l
}

func (l *FiberZapLogger) WithAny(key string, value interface{}) Logger {
	return l
}

func (l *FiberZapLogger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *FiberZapLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

func (l *FiberZapLogger) Debugc(ctx context.Context, template string, args ...interface{}) {
	l.logger.WithContext(ctx).Debugf(template, args...)
}

func (l *FiberZapLogger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *FiberZapLogger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

func (l *FiberZapLogger) Infoc(ctx context.Context, template string, args ...interface{}) {
	l.logger.WithContext(ctx).Infof(template, args...)
}

func (l *FiberZapLogger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *FiberZapLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

func (l *FiberZapLogger) Warnc(ctx context.Context, template string, args ...interface{}) {
	l.logger.WithContext(ctx).Warnf(template, args...)
}

func (l *FiberZapLogger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *FiberZapLogger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

func (l *FiberZapLogger) Errorc(ctx context.Context, template string, args ...interface{}) {
	l.logger.WithContext(ctx).Errorf(template, args...)
}

func (l *FiberZapLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *FiberZapLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}

func (l *FiberZapLogger) Fatalc(ctx context.Context, template string, args ...interface{}) {
	l.logger.WithContext(ctx).Fatalf(template, args...)
}

func (l *FiberZapLogger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l *FiberZapLogger) Panicf(template string, args ...interface{}) {
	l.logger.Panicf(template, args...)
}

func (l *FiberZapLogger) Panicc(ctx context.Context, template string, args ...interface{}) {
	l.logger.WithContext(ctx).Panicf(template, args...)
}
