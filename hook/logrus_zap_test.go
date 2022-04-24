package hook

import (
	v2 "github.com/openmsp/cilog/v2"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"io/ioutil"
	"testing"
)

func TestNewLogToZapHook(t *testing.T) {
	hook := NewLogToZapHook(0)
	logrus.AddHook(hook)

	v2.GlobalConfig(v2.Conf{
		Caller: true,
	})

	zap.ReplaceGlobals(v2.N())

	logrus.SetOutput(ioutil.Discard)
	logrus.Error("test")
}
