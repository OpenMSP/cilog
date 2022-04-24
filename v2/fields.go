package v2

import (
	"bytes"
	"github.com/openmsp/cilog"
	"net"
	"os"
	"runtime"
	"strconv"
)

const defaultTimestampFormat = "2006-01-02T15:04:05.000-0700"

const (
	fieldAppName       = "appName"       //微服务appName
	fieldAppID         = "appId"         //服务appId
	fieldAppVersion    = "appVersion"    //微服务app版本号
	fieldAppKey        = "appKey"        //appkey
	fieldChannel       = "channel"       //channel
	fieldSubOrgKey     = "subOrgKey"     //机构唯一码
	fieldTime          = "timestamp"     //日志时间字符串
	fieldLevel         = "level"         //日志等级 : DEBUG、INFO 、NOTICE、WARNING、ERR、CRIT、ALERT、 EMERG(系统不可用)
	fieldHostName      = "hostname"      //主机名
	fieldIP            = "ip"            //ip地址
	fieldPodName       = "podName"       //pod名
	fieldPodIp         = "podIp"         //pod IP
	fieldNodeName      = "nodeName"      //pod内部的node名
	fieldNodeIp        = "nodeIp"        //k8s注入的node节点IP
	fieldContainerName = "containerName" //k8s容器name ，主要进行容器环境区分
	fieldClusterUid    = "clusterUid"    //集群ID
	fieldImageUrl      = "imageUrl"      //应用镜像URL地址
	fieldUniqueId      = "uniqueId"      //部署的服务唯一ID
	fieldSiteUID       = "siteUid"       //可用区唯一标识符
	fieldRunEnvType    = "runEnvType"    //区分开发环境(development)、测试环境(test)、预发布环境 (pre_release)、生产环境(production) 从环境变量中获取
	fieldMessage       = "message"       //日志内容
	fieldLogger        = "logger"        //日志来源函数名
	fieldType          = "type"          //当前日志的所处动作环境，ACCESS|EVENT|RPC|OTHER
	fieldTitle         = "title"         //日志标题，不传就是message前100个字符
	fieldPID           = "pid"           //进程id
	fieldThreadId      = "threadId"      //线程id
	fieldLanguage      = "language"      //语言标识
	fieldURL           = "url"           //⻚面/接口URL
	fieldClientIP      = "clientIp"      //调用者IP
	fieldErrCode       = "errCode"       //异常码
	fieldTraceID       = "traceID"       //全链路TraceId
	fieldSpanID        = "spanID"        //全链路SpanId :在非span产生的上下文环境中，可以留空
	fieldParentID      = "parentID"      //全链路 上级SpanId :在非span产生的上下文环境中，可以留空
	fieldCustomLog1    = "customLog1"    //自定义log1
	fieldCustomLog2    = "customLog2"    //自定义log2
	fieldCustomLog3    = "customLog3"    //自定义log3
)

const (
	DynFieldType       = fieldType
	DynFieldURL        = fieldURL
	DynFieldClientIP   = fieldClientIP
	DynFieldErrCode    = fieldErrCode
	DynFieldTraceID    = fieldTraceID
	DynFieldSpanID     = fieldSpanID
	DynFieldParentID   = fieldParentID
	DynFieldCustomLog1 = fieldCustomLog1
	DynFieldCustomLog2 = fieldCustomLog2
	DynFieldCustomLog3 = fieldCustomLog3
)

var allowedFields = map[string]struct{}{
	fieldAppName:       {},
	fieldAppID:         {},
	fieldAppVersion:    {},
	fieldAppKey:        {},
	fieldChannel:       {},
	fieldSubOrgKey:     {},
	fieldTime:          {},
	fieldLevel:         {},
	fieldHostName:      {},
	fieldIP:            {},
	fieldPodName:       {},
	fieldPodIp:         {},
	fieldNodeName:      {},
	fieldNodeIp:        {},
	fieldContainerName: {},
	fieldClusterUid:    {},
	fieldImageUrl:      {},
	fieldUniqueId:      {},
	fieldSiteUID:       {},
	fieldRunEnvType:    {},
	fieldMessage:       {},
	fieldLogger:        {},
	fieldType:          {},
	fieldTitle:         {},
	fieldPID:           {},
	fieldThreadId:      {},
	fieldLanguage:      {},
	fieldURL:           {},
	fieldClientIP:      {},
	fieldErrCode:       {},
	fieldTraceID:       {},
	fieldSpanID:        {},
	fieldParentID:      {},
	fieldCustomLog1:    {},
	fieldCustomLog2:    {},
	fieldCustomLog3:    {},
}

func getFixedFields(app *cilog.ConfigAppData) map[string]string {
	fields := map[string]string{
		fieldAppName:       app.AppName,
		fieldAppID:         app.AppID,
		fieldAppVersion:    app.AppVersion,
		fieldAppKey:        app.AppKey,
		fieldChannel:       app.Channel,
		fieldSubOrgKey:     app.SubOrgKey,
		fieldTime:          "",
		fieldLevel:         "",
		fieldHostName:      getHostname(),
		fieldIP:            getInternetIP(),
		fieldPodName:       os.Getenv("PODNAME"),
		fieldPodIp:         os.Getenv("PODIP"),
		fieldNodeName:      os.Getenv("NODENAME"),
		fieldNodeIp:        os.Getenv("NODEIP"),
		fieldContainerName: os.Getenv("CONTAINERNAME"),
		fieldClusterUid:    os.Getenv("IDG_CLUSTERUID"),
		fieldImageUrl:      os.Getenv("IDG_IMAGEURL"),
		fieldUniqueId:      os.Getenv("IDG_UNIQUEID"),
		fieldSiteUID:       os.Getenv("IDG_SITEUID"),
		fieldRunEnvType:    os.Getenv("IDG_RUNTIME"),
		fieldMessage:       "",
		fieldLogger:        "",
		fieldType:          "ACCESS",
		fieldTitle:         "",
		fieldPID:           strconv.Itoa(os.Getpid()),
		fieldLanguage:      app.Language,
		fieldURL:           "",
		fieldClientIP:      "",
		fieldErrCode:       "",
		fieldTraceID:       "",
		fieldSpanID:        "",
		fieldParentID:      "",
		fieldCustomLog1:    "",
		fieldCustomLog2:    "",
		fieldCustomLog3:    "",
	}
	if fields[fieldAppName] == "" {
		fields[fieldAppName] = os.Getenv("IDG_SERVICE_NAME")
	}
	if fields[fieldAppID] == "" {
		fields[fieldAppID] = os.Getenv("IDG_APPID")
	}
	if fields[fieldAppVersion] == "" {
		fields[fieldAppVersion] = os.Getenv("IDG_VERSION")
	}
	if fields[fieldAppKey] == "" {
		fields[fieldAppKey] = os.Getenv("IDG_APPKEY")
	}
	return fields
}

func getHostname() (Hostname string) {
	// 查找本机hostname
	hostname, err := os.Hostname()
	if err != nil {
		return ""
	}
	Hostname = hostname
	return
}

func getInternetIP() (IP string) {
	// 查找本机IP
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ip4 := ipnet.IP.To4(); ip4 != nil {
				if ip4[0] == 10 {
					// 赋值新的IP
					IP = ip4.String()
				}
			}
		}
	}
	return
}

// getGID 获取协程ID
func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
