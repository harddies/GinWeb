package log

import (
	"GinWeb/config"
	"github.com/harddies/glog"
)

var Logger *glog.Logger

func init() {
	options := glog.LogOptions{
		LogDir:  config.C.Log.LogDir,
		LogName: config.C.Log.LogName,
		Flag:    glog.DefaultFlags,
		Maxsize: 100000, // megabytes
	}
	switch config.C.Log.LogLevel {
	case "debug":
		options.Level = glog.DebugLevel
		break
	case "info":
		options.Level = glog.InfoLevel
		break
	case "warn":
		options.Level = glog.WarnLevel
		break
	case "error":
		options.Level = glog.ErrorLevel
		break
	case "fatal":
		options.Level = glog.FatalLevel
		break
	case "panic":
		options.Level = glog.PanicLevel
		break
	default:
		options.Level = glog.DebugLevel
		break
	}

	switch config.C.Log.LogRotateMode {
	case "none":
		options.Mode = glog.RotateNone
		break
	case "size":
		options.Mode = glog.RotateSize
		options.Maxsize = config.C.Log.Maxsize
		break
	case "hour":
		options.Mode = glog.RotateHour
		break
	case "day":
		options.Mode = glog.RotateDay
		break
	default:
		options.Mode = glog.RotateNone
		break
	}
	var err error
	Logger, err = glog.New(options)
	if err != nil {
		// TODO: 报警处理（短信、邮件）
	}
}
