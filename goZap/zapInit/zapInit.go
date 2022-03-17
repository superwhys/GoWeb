package zapInit

import "go.uber.org/zap"

var (
	Logger      *zap.Logger
	SugarLogger *zap.SugaredLogger
)

func InitLogger() {
	// zap.NewProduction()/zap.NewDevelopment()或者zap.Example()
	Logger, _ = zap.NewProduction()
}

func InitSugaredLogger() {
	baseLogger, _ := zap.NewProduction()
	SugarLogger = baseLogger.Sugar()
}
