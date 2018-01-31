package main

import (
	"os"

	"github.com/cihub/seelog"
)

var (
	Logger   seelog.LoggerInterface
	err      error
	basePath = os.Getenv("GOPATH")
)

func init() {
	Logger, err = seelog.LoggerFromConfigAsFile(basePath + "/src/go-cms/log.seelog.xml")
	if err != nil {
		seelog.Critical("err parsing config log file", err)
	}
}
