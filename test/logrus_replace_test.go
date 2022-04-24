package test

import (
	"github.com/openmsp/cilog"
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

var (
	LogProxyHost = "192.168.2.80"
	LogProxyPort = 6381
)

func TestLogrusReplace(t *testing.T) {
	os.Setenv("IDG_UNIQUEID", "test_unique_id")
	configLogData := &cilog.ConfigLogData{
		OutPut: "redis",
		Debug:  false,
		Level:  logrus.InfoLevel,
		Key:    os.Getenv("CILOG_KEY"),
		Redis: struct {
			Host string
			Port int
		}{Host: LogProxyHost, Port: LogProxyPort},
	}
	cilog.ConfigLogInit(configLogData, &cilog.ConfigAppData{
		Language: "en",
		AppID:    "test_app",
	})

	log := cilog.GetLogger().WithFields(cilog.Fields)
	log.Info("test")
}
