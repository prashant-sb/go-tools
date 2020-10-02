package main

import (
	"flag"
	log "github.com/golang/glog"
)

const (
	STDOUT = "/dev/stdout"
)

var (
	format = flag.String("format", "json", "Prints the information in json | pretty format")
	saveas = flag.String("saveas", STDOUT, "Saves the information in given file or stdout")
)

func main() {
	flag.Parse()

	sysinfo, err := NewInformer()
	if err != nil {
		log.Error("Error: ", err.Error())
		return
	}
	fopt := sysinfo.FmtOption()

	fopt.FormatAs(*format)
	err = fopt.To(*saveas)
	if err != nil {
		log.Error("Error : ", err.Error())
	}
}