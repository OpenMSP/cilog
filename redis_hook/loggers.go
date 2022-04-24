package redis_hook

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/zap/zapcore"
	"strings"
)

func CreateLogrusOriginLogMessage(entry *logrus.Entry) map[string]interface{} {
	fields := make(map[string]interface{})
	for k, v := range entry.Data {
		fields[k] = v
	}
	var level = strings.ToUpper(entry.Level.String())
	if level == "ERROR" {
		level = "ERR"
	}
	fields["level"] = level
	fields["message"] = entry.Message
	return fields
}

// zap need extra data for fields
func CreateZapOriginLogMessage(entry *zapcore.Entry, data map[string]interface{}) map[string]interface{} {
	fields := make(map[string]interface{}, len(data))
	if data != nil {
		for k, v := range data {
			fields[k] = v
		}
	}
	var level = strings.ToUpper(entry.Level.String())
	if level == "ERROR" {
		level = "ERR"
	}
	if level == "WARN" {
		level = "WARNING"
	}
	if level == "FATAL" {
		level = "CRIT"
	}
	fields["level"] = level
	fields["message"] = entry.Message
	return fields
}
