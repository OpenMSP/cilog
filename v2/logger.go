package v2

import (
	"github.com/openmsp/cilog"
	"github.com/openmsp/cilog/redis_hook"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ciCore zapcore.Core

var (
	std          *Logger
	stdCallerFix *Logger

	n *zap.Logger
)

// Logger 实例
type Logger struct {
	*zap.SugaredLogger
	conf *Conf
}

// Conf 配置
type Conf struct {
	Caller     bool
	Debug      bool
	StackLevel zapcore.Level
	Level      zapcore.Level
	Encoding   string                 // json, console
	AppInfo    *cilog.ConfigAppData   // fixed fields
	HookConfig *redis_hook.HookConfig // set to nil if disabled
	ZapConfig  *zap.Config            // for custom
}

// Clone ...
func Clone(l *Logger) *Logger {
	c := *l.conf
	return &Logger{
		SugaredLogger: l.SugaredLogger,
		conf:          &c,
	}
}

// S 获取单例
func S() *Logger {
	return std
}

// N 获取 Zap Logger
func N() *zap.Logger {
	return n
}

// GlobalConfig init
func GlobalConfig(conf Conf) error {
	c := conf
	l, err := newLogger(&c)
	if err != nil {
		return err
	}
	std = &Logger{
		SugaredLogger: l.Sugar(),
		conf:          &c,
	}
	stdCallerFix = &Logger{
		SugaredLogger: l.WithOptions(zap.AddCallerSkip(1)).Sugar(),
		conf:          &c,
	}
	n = std.Desugar()
	return nil
}

func init() {
	l, _ := newLogger(&Conf{
		Level:      zapcore.InfoLevel,
		StackLevel: zapcore.ErrorLevel,
	})
	std = &Logger{
		SugaredLogger: l.Sugar(),
		conf:          &Conf{},
	}
	stdCallerFix = &Logger{
		SugaredLogger: l.WithOptions(zap.AddCallerSkip(1)).Sugar(),
		conf:          &Conf{},
	}
	n = std.Desugar()
}

// NewZapLogger 创建自定义 Logger
func NewZapLogger(c *Conf) (l *zap.Logger, err error) {
	return newLogger(c)
}

func newLogger(c *Conf) (l *zap.Logger, err error) {
	var conf zap.Config
	if c.ZapConfig != nil {
		conf = *c.ZapConfig
	} else {
		conf = zap.NewProductionConfig()
		conf.EncoderConfig = zap.NewDevelopmentEncoderConfig()
		if c.Debug {
			conf = zap.NewDevelopmentConfig()
			conf.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		}
		if c.Encoding != "" {
			conf.Encoding = c.Encoding
		} else {
			conf.Encoding = "console"
		}
	}
	conf.Level = zap.NewAtomicLevelAt(c.Level)
	if c.HookConfig != nil {
		hook, _ := redis_hook.NewHook(*c.HookConfig)
		_ciCore = NewCiCore(hook)
		fixedFields := getFixedFields(c.AppInfo)
		for k, v := range fixedFields {
			_ciCore = _ciCore.With([]zapcore.Field{zap.String(k, v)})
		}
		if !c.Debug {
			l, err = conf.Build(zap.WrapCore(func(core zapcore.Core) zapcore.Core {
				return zapcore.NewTee(core, _ciCore)
			}))
		} else {
			l, err = conf.Build()
		}
	} else {
		l, err = conf.Build()
	}
	if err != nil {
		return nil, errors.Wrap(err, "zap core init error")
	}
	l = l.WithOptions(zap.WithCaller(c.Caller), zap.AddStacktrace(c.StackLevel))
	return
}
