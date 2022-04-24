package hook

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

var _ logrus.Hook = (*LogrusToZapHook)(nil)

type LogrusToZapHook struct {
	callerSkip int
}

func NewLogToZapHook(callerSkip int) logrus.Hook {
	return &LogrusToZapHook{
		callerSkip: 8 + callerSkip,
	}
}

func (l LogrusToZapHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (l LogrusToZapHook) Fire(entry *logrus.Entry) error {
	var fields []zap.Field
	if entry.Data != nil {
		for k, v := range entry.Data {
			fields = append(fields, zap.Any(k, v))
		}
	}
	logger := zap.L().WithOptions(zap.AddCallerSkip(l.callerSkip))
	switch entry.Level {
	case logrus.TraceLevel, logrus.DebugLevel:
		logger.With(fields...).Debug(entry.Message)
	case logrus.InfoLevel:
		logger.With(fields...).Info(entry.Message)
	case logrus.WarnLevel:
		logger.With(fields...).Warn(entry.Message)
	case logrus.ErrorLevel:
		logger.With(fields...).Error(entry.Message)
	case logrus.FatalLevel:
		logger.With(fields...).Fatal(entry.Message)
	case logrus.PanicLevel:
		logger.With(fields...).Panic(entry.Message)
	default:
	}
	return nil
}
