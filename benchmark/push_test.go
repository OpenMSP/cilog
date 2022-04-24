package benchmark

import (
	"github.com/openmsp/cilog"
	"github.com/openmsp/cilog/redis_hook"
	v2 "github.com/openmsp/cilog/v2"
	"go.uber.org/zap/zapcore"
	"strconv"
	"strings"
	"testing"
)

var (
	logProxyHost string
	logProxyPort int
)

func init() {
	spit := strings.Split("127.0.0.1:6381", ":")
	port, err := strconv.Atoi(spit[1])
	if err != nil {
		panic(err)
	}
	logProxyHost = spit[0]
	logProxyPort = port
}

func BenchmarkV2(b *testing.B) {
	v2.GlobalConfig(v2.Conf{
		Caller: false,
		Debug:  false,
		Level:  zapcore.InfoLevel,
		AppInfo: &cilog.ConfigAppData{
			AppVersion: "1.0",
			Language:   "zh-cn",
		},
		HookConfig: &redis_hook.HookConfig{
			Key:  "gw_log",
			Host: logProxyHost,
			Port: logProxyPort,
		},
	})
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		v2.Info("msg")
	}
}

func BenchmarkV2Desuger(b *testing.B) {
	v2.GlobalConfig(v2.Conf{
		Caller: false,
		Debug:  false,
		Level:  zapcore.InfoLevel,
		AppInfo: &cilog.ConfigAppData{
			AppVersion: "1.0",
			Language:   "zh-cn",
		},
		HookConfig: &redis_hook.HookConfig{
			Key:  "gw_log",
			Host: logProxyHost,
			Port: logProxyPort,
		},
	})
	desuger := v2.S().Desugar()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		desuger.Info("msg")
	}
}