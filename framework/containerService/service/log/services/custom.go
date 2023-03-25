package services

import (
	"github.com/tianzhaocn/skyscraper/framework"
	"github.com/tianzhaocn/skyscraper/framework/containerService/contract"
	"io"
)

type CustomLog struct {
	Log
}

func NewCustomLog(params ...interface{}) (interface{}, error) {
	c := params[0].(framework.Container)
	level := params[1].(contract.LogLevel)
	ctxFielder := params[2].(contract.CtxFielder)
	formatter := params[3].(contract.Formatter)
	output := params[4].(io.Writer)

	log := &ConsoleLog{}

	log.SetLevel(level)
	log.SetCtxFielder(ctxFielder)
	log.SetFormatter(formatter)

	log.SetOutput(output)
	log.c = c
	return log, nil
}
