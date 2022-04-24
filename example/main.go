package main

import (
	"errors"
	"github.com/openmsp/cilog"
	"github.com/sirupsen/logrus"
)

func main() {
	configLogData := &cilog.ConfigLogData{
		OutPut: "redis",           //redis 输出到日志中心redis ，stdout 输出到终端,both 全部输出
		Debug:  true,              //目前无用，后期增加输出stack功能
		Key:    "my_redis_key",    //发送到redis的key名
		Level:  logrus.TraceLevel, //设置打印级别
		Redis: struct {
			Host string
			Port int
		}{Host: "127.0.0.1", Port: 6379},
	}
	configAppData := &cilog.ConfigAppData{
		AppName:    "AppName",
		AppID:      "AppID",
		AppVersion: "AppVersion",
		AppKey:     "AppKey",
		Channel:    "Channel",
		SubOrgKey:  "SubOrgKey",
		Language:   "ch",
	}
	cilog.ConfigLogInit(configLogData, configAppData)
	cilog.LogErrordw(cilog.LogNameMysql, cilog.DynamicFields{
		TraceID:   "TraceId",
		AppKey:    "AppKey",
		Channel:   "Channel",
		SubOrgKey: "SubOrgKey",
	}, "mysql err", errors.New("connect error"))
	cilog.LogErrorw(cilog.LogNameMysql, "mysql err", errors.New("connect error"))
}
