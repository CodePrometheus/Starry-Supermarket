package initialize

import "go.uber.org/zap"

func Logger() {
	// 定义全局
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}
