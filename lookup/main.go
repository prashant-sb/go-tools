package main

import (
	"flag"

	log "github.com/golang/glog"
	lkp "github.com/prashant-sb/go-tools/lookup/resolv"
)

var (
	name  = flag.String("name", "", "lookup given hostname")
	ip    = flag.String("ip", "", "Lookup given ipv4 / ipv6")
	debug = flag.Bool("verbose", false, "Debug messages")
)

func main() {
	flag.Parse()
	if *name != "" && lkp.IsValidName(*name) {
		log.Info("Searching for domain name ", *name)
		if err := lkp.ResolveIP(*name); err != nil {
			log.Error(err, "Host resolution failed")
			return
		}
	}

	if *ip != "" && lkp.IsValidIP(*ip) {
		log.Info("Searching for IP ", *ip)
		if err := lkp.ResolveHost(*ip); err != nil {
			log.Error(err, "IP resolution failed")
			return
		}
	}

}
