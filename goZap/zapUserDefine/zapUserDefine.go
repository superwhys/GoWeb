package zapUserDefine

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	UdLogger      *zap.Logger
	UdLoggerSugar *zap.SugaredLogger
)

func getEncoder() zapcore.Encoder {
	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	// 更改为普通的encoder
	//return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // 更改时间编码
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 更改level格式
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	//file, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)

	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test2.log",
		MaxSize:    1,  // m
		MaxAge:     5,  // 最大备份数量
		MaxBackups: 30, // 最大备份天数
		LocalTime:  false,
		Compress:   false, // 是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
	//return zapcore.AddSync(os.Stdout)
}

func InitUdLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	// zapcore.Core需要三个配置——Encoder，WriteSyncer，LogLevel。
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	// zap.AddCaller() 添加调用者信息
	UdLogger = zap.New(core, zap.AddCaller())
	UdLoggerSugar = UdLogger.Sugar()
}
