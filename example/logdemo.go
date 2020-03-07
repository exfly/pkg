package main

import (
	"github.com/exfly/pkg/log"
)

type empty struct{}

var LOG = log.GetLogger(log.PkgPath(empty{}))

func main() {
	LOG.SetReportCaller(true)
	LOG.SetFormatter("", &log.JSONFormatter{})
	LOG.Infoln("aaaa")

	LOG.Info("aaa")
}
