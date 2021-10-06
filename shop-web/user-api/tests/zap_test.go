package main

import (
	"go.uber.org/zap"
	"testing"
)

func TestLogOut(t *testing.T) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./zap.log",
	}
}

func TestLog(t *testing.T) {
	//logger, _ := tests.NewProduction() // 生产环境
	logger, _ := zap.NewDevelopment() // 开发环境
	defer logger.Sync()               // flushes buffer, if any
	//sugar := logger.Sugar()
	//url := "http://localhost:9000"
	logger.Info("Failed to fetch URL: %s")
	//sugar.Infow("failed to fetch URL",
	//	// Structured context as loosely typed key-value pairs.
	//	"url", url,
	//	"attempt", 3,
	//	"backoff", time.Second,
	//)
	//sugar.Infof("Failed to fetch URL: %s", url)
}
