package globar

import (
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"
)

var (
	//// zap日志
	Logger *zap.SugaredLogger
	//认证
	Validate *validator.Validate
)
