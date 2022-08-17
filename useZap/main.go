package main

import (
	"fmt"

	"useZap/zapdemo"
	"go.uber.org/zap"
)

var logger *zap.Logger

func ProductionLogger() {
	logger, _ = zap.NewProduction()
	logger.Info("NewProduction log info")
}

func DevelopmentLogger() {
	logger, _ = zap.NewDevelopment()
	logger.Info("NewDevelopment log info")
}

func ExampleLogger() {
	logger = zap.NewExample()
	logger.Info("NewExample log info")
}

func CreateLogger() {
	// 初始化logger
	logger := zap.NewExample()
	// 使用defer logger.Sync()将缓存同步到文件中。
	defer logger.Sync()
	// 记录日志
	logger.Info("NewExample",
		zap.String("name", "张三"),
		zap.Int("age", 18),
	)
	productionLogger, _ := zap.NewProduction()
	defer productionLogger.Sync()
	productionLogger.Info("NewProduction",
		zap.String("name", "张三"),
		zap.Int("age", 18),
	)
	devLogger, _ := zap.NewDevelopment()
	defer devLogger.Sync()
	devLogger.Info("NewDevelopment",
		zap.String("name", "张三"),
		zap.Int("age", 18),
	)
}

func main() {
	ProductionLogger()
	DevelopmentLogger()
	ExampleLogger()
	fmt.Println()
	CreateLogger()

	fmt.Println()

	err := zapdemo.Init()
	if err != nil {
		panic(err)
	}

	zapdemo.SugarLogger.Debugf("This is err: %s", "ayunwSky")
}
