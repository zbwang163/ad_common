package logs

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/zbwang163/ad_common/env"
	"github.com/zbwang163/ad_common/time_utils"
	"os"
	"time"
)

var defaultLogger *logrus.Entry

func InitLogger() {
	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&logrus.JSONFormatter{}) //log日志序列化为json
	log.SetReportCaller(true)                 // 打印日志位置

	file, err := os.OpenFile(fmt.Sprintf("/var/log/ad_platform_info/%v.log", time_utils.Time20060102_15(time.Now())),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	ctxLog := log.WithFields(logrus.Fields{
		"ip":       env.GetPodIp(),
		"pod_name": env.GetPodName(),
	})
	defaultLogger = ctxLog
}

func CtxInfo(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.WithField("log_id", ctx.Value("log_id")).Infof(format, args)
}

func CtxError(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.WithField("log_id", ctx.Value("log_id")).Errorf(format, args)
}
