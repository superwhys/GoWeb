package main

import (
	"github.com/superwhys/GoWeb/goZap/zapInit"
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

func main() {
	zapInit.InitLogger()
	zapInit.InitSugaredLogger()
	defer zapInit.Logger.Sync()

	simpleHttpGetLogger("https://www.baidu.com")
	simpleHttpGetLogger("http://www.google.com")

	simpleHttpGetSugarLogger("https://www.baidu.com")
	simpleHttpGetSugarLogger("http://www.google.com")
}
