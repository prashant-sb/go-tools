package main

import (
	"github.com/prashant-sb/go-tools/async/routine"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
}

func main() {
	log.Info("Running simple async singhandler demo")
	rt := routine.NewRunner()
	rt.Start()
}
