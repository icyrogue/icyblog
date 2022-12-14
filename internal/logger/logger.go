package logger

import (
	"go.uber.org/zap"
)

type Options struct {
	mode string
}

func New(options *Options) *zap.SugaredLogger {
	config := zap.NewProductionConfig()
	//config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stdout"}
	logger, _ := config.Build()
	return logger.Sugar()
}
