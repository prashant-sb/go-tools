package main

import (
	"flag"

	log "github.com/golang/glog"
)

var (
	daemonize = flag.Bool("daemonize", true, "stdin | daemonize run")
	config    = flag.String("config", "", "Ldap server configuration file")
)

func main() {

	flag.Parse()
	log.Info("Started ...")

	return
}
