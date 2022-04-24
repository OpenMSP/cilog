package cilog

import "github.com/sirupsen/logrus"

var (
    logConfig *ConfigLog
)

type ConfigLog struct {
    Log *ConfigLogData
    App *ConfigAppData
}

type ConfigLogData struct {
    OutPut string
    Debug  bool
    Key    string
    Level  logrus.Level
    Redis  struct {
        Host string
        Port int
    }
}

type ConfigAppData struct {
    AppName    string
    AppID      string
    AppVersion string
    AppKey     string
    Channel    string
    SubOrgKey  string
    Language   string
}

func ConfigLogInit(configLogData *ConfigLogData, configAppData *ConfigAppData) {
    logConfig = &ConfigLog{
        Log: configLogData,
        App: configAppData,
    }
    if logConfig.App == nil {
        logConfig = configLogGetDefault()
    }
    loggerInit()
}

func configLogGetDefault() *ConfigLog {
    c := new(ConfigLog)
    c.Log = &ConfigLogData{}
    c.App = &ConfigAppData{}
    c.App.AppName = "appName"
    c.Log.Debug = true
    c.Log.OutPut = "stdout"
    return c
}
