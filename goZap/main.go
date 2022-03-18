package main

import (
	"github.com/superwhys/GoWeb/goZap/zapInit"
	"github.com/superwhys/GoWeb/goZap/zapUserDefine"
	"go.uber.org/zap"
	"net/http"
)

func simpleHttpGetLogger(url string) {
	resp, err := http.Get(url)
	if err != nil {
		zapInit.Logger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		zapInit.Logger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}

func simpleHttpGetSugarLogger(url string) {
	resp, err := http.Get(url)
	if err != nil {
		zapInit.SugarLogger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		zapInit.SugarLogger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}

func simpleHttpGetUdLogger(url string) {
	resp, err := http.Get(url)
	if err != nil {
		zapUserDefine.UdLogger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		zapUserDefine.UdLogger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}

func simpleHttpGetSugarUdLogger(url string) {
	resp, err := http.Get(url)
	if err != nil {
		zapUserDefine.UdLoggerSugar.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		zapUserDefine.UdLoggerSugar.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}

func main() {
	zapInit.InitLogger()
	zapInit.InitSugaredLogger()
	defer zapInit.Logger.Sync()
	defer zapInit.SugarLogger.Sync()

	simpleHttpGetLogger("https://www.baidu.com")
	//simpleHttpGetLogger("http://www.google.com")

	simpleHttpGetSugarLogger("https://www.baidu.com")
	//simpleHttpGetSugarLogger("http://www.google.com")

	zapUserDefine.InitUdLogger()
	defer zapUserDefine.UdLogger.Sync()
	defer zapUserDefine.UdLoggerSugar.Sync()

	simpleHttpGetSugarUdLogger("https://www.baidu.com")
	simpleHttpGetUdLogger("https://www.baidu.com")
	for i := 0; i < 10000; i++ {
		zapUserDefine.UdLoggerSugar.Info("test logger")
	}
}
