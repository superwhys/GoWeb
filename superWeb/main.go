package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"

	"github.com/superwhys/superWeb/dao/mysql"
	"github.com/superwhys/superWeb/dao/redis"
	"github.com/superwhys/superWeb/logger"
	"github.com/superwhys/superWeb/routes"
	"github.com/superwhys/superWeb/settings"
	"go.uber.org/zap"
)

// main GO Web 开发通用的脚手架
func main() {
	// 1. 加载配置
	if err := settings.InitSetting(); err != nil {
		zap.L().Error("read config failed", zap.Error(err))
		return
	}
	// 2. 初始化日志
	if err := logger.Init(); err != nil {
		zap.L().Error("init logger failed", zap.Error(err))
		return
	}
	// 将缓冲区的日志写入文件
	defer zap.L().Sync()
	// 3. 初始化Mysql连接
	if err := mysql.Init(); err != nil {
		zap.L().Error("connect mysql failed", zap.Error(err))
		return
	}
	defer mysql.Close()
	// 4. 初始化Redis连接
	if err := redis.Init(); err != nil {
		zap.L().Error("connect redis failed", zap.Error(err))
		return
	}
	defer redis.Close()
	// 5. 注册路由
	router := routes.SetUp()
	// 6. 启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: router,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen error: ", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
