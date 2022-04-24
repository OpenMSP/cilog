package example

import (
	"github.com/openmsp/cilog"
	"github.com/openmsp/cilog/redis_hook"
	logger "github.com/openmsp/cilog/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

var (
	EnvEnableRedisOutput bool // 模拟环境变量
	EnvDebug             bool
)

func init() {
	EnvEnableRedisOutput = true
	EnvDebug = true
	initLogger()
}

func initLogger() {
	conf := &logger.Conf{
		Level:  zapcore.DebugLevel, // 输出日志等级
		Caller: true,               // 是否开启记录调用文件夹+行数+函数名
		Debug:  true,               // 是否开启 Debug
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
		conf.StackLevel = zapcore.DPanicLevel // 指定输出函数栈的日志级别，默认是ERROR级别
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

func testLogger() {
	// 获取 Logger 实例
	instance := logger.S()
	// 严格模式
	strict := instance.Desugar()
	// 自定义 Options
	lo := strict.WithOptions(zap.AddStacktrace(zapcore.DebugLevel)).Sugar()
	lo.Named("custom").Info("自定义 Logger")
	// Hook 全局
	zap.ReplaceGlobals(logger.S().Desugar())
}