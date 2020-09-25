package main

import (
	"flag"

	log "github.com/golang/glog"
)

var (
	list = flag.Bool("list", false, "List running processes")
	wpid = flag.Uint64("watch", 0, "watch process by pid for events")
)

func main() {
	flag.Parse()

	procIter := NewProcIterator()

	if *list == true {
		log.Info("Listing processes ")
		if err := procIter.List(); err != nil {
			log.Error("Error : ", err.Error())
		}
	}

	if *wpid != 0 {
		procIter.Watch(*wpid)
	}
}
