package v2

import (
	"testing"

	"github.com/openmsp/cilog"
	"github.com/openmsp/cilog/redis_hook"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestNewLogger(t *testing.T) {
	defer Sync()
	GlobalConfig(Conf{
		Debug:  true,
		Caller: true,
		AppInfo: &cilog.ConfigAppData{
			AppName:    "test",
			AppID:      "test",
			AppVersion: "1.0",
			AppKey:     "test",
			Channel:    "1",
			SubOrgKey:  "key",
			Language:   "zh",
		},
	})
	S().Info("test")
}

func TestColorLogger(t *testing.T) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := config.Build()

	logger.Info("Now logs should be colored")
}

func TestAddtionFieldsLogger(t *testing.T) {
	GlobalConfig(Conf{
		Debug:  true,
		Caller: true,
		AppInfo: &cilog.ConfigAppData{
			AppName:    "test",
			AppID:      "test",
			AppVersion: "1.0",
			AppKey:     "test",
			Channel:    "1",
			SubOrgKey:  "key",
			Language:   "zh",
		},
		HookConfig: &redis_hook.HookConfig{
			Host: "127.0.0.1",
			Port: 9090,
		},
	})

	s := struct {
		AppID string
	}{
		AppID: "test",
	}

	S().With("test", s, "test2", &s, "test3", "test3").Info("test")
}
