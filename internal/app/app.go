package app

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New() *fx.App {
	return fx.New(
		fx.Provide(
			newDevelopmentLogger,
			newSugarLogger,
		),
		fx.WithLogger(func(l *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: l}
		}),
		fx.Invoke(do),
	)
}

func newSugarLogger(l *zap.Logger) *zap.SugaredLogger {
	return l.Sugar()
}

func newDevelopmentLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := config.Build()

	logger.Info("Now logs should be colored")

	return logger
}

func do(l *zap.SugaredLogger) {
	l.Infoln("jopa")
	l.Infoln("jopa")
	l.Infoln("jopa")
}
