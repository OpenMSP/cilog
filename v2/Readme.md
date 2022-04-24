# Cilog V2

**V2 版本目前仅支持输出到 Log Proxy**

### Features
1. level 改造符合公司日志标准
2. 兼容 Zap 所有方法
3. 额外添加 `package.Info()` 系列方法
4. 支持动态自定义 Field
5. 支持启用 Caller 记录调用函数
6. Error 级别记录调用 Stack

默认情况下日志会同时输出到 stdout, 以及 Log Proxy.    

### Usage

在目录 [v2/example](./v2/example) 查看使用示例.   

```go
import (
	"gitlab.oneitfarm.com/bifrost/cilog"
	"gitlab.oneitfarm.com/bifrost/cilog/redis_hook"
	logger "gitlab.oneitfarm.com/bifrost/cilog/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

func initLogger() {
	conf := &logger.Conf{
		Level:  zapcore.DebugLevel, // 输出日志等级
		Caller: true, // 是否开启记录调用文件夹+行数+函数名
		Debug:  true, // 是否开启 Debug
		// 输出到 redis 的日志全部都是 info 级别以上
		// 不用填写 AppName, AppID 默认会从环境变量获取
		AppInfo: &cilog.ConfigAppData{
			AppVersion: "1.0",
			Language:   "zh-cn",
		},
	}
	if !EnvDebug || EnvEnableRedisOutput {
		// 如果是生产环境
		conf.Level = zapcore.InfoLevel
		conf.HookConfig = &redis_hook.HookConfig{
			Key:  "gw_log",                      // 填写日志 key
			Host: "redis-cluster-proxy-log.msp", // 填写 log proxy host
			// k8s 集群内填写 redis-cluster-proxy-log.msp
			Port: 6380, // 填写 log proxy port
			// 默认填写 6380
		}
	}
	err := logger.GlobalConfig(*conf)
	if err != nil {
		// 处理 logger 初始化错误
		// log-proxy 连接失败会报错
		// 若不影响程序执行，可忽视
		log.Print("[ERR] Logger init error: ", err)
	}
	logger.With(logger.DynFieldErrCode, 400).Debug("测试附加字段")
	logger.Infof("info test: %v", "data")
}
```

#### Notice
HookConfig 设置为 `nil`, 则不输出到 Redis.       
本地测试环境下请设置为 `nil`.       