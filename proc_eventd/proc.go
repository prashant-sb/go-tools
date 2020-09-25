package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	log "github.com/golang/glog"
)

const (
	PROCDIR    = "/proc/"
	PROCSTATUS = PROCDIR + "*/status"
)

type ProcIter interface {
	List() error
	Watch(pid uint64) error
}

type ProcMeta struct {
	procAttrs map[string]string
}

type ProcEntry struct {
	procMap map[string]*ProcMeta
}

func NewProcIterator() ProcIter {

	return &ProcEntry{
		procMap: initProcMap(),
	}
}

func getProcMeta() ProcMeta {
	pmeta := make(map[string]string)

	for _, attr := range []string{
		"Name",
		"Pid",
		"PPid",
		"Tgid",
		"State",
		"Umask",
		"Threads",
	} {
		pmeta[attr] = ""
	}

	return ProcMeta{
		procAttrs: pmeta,
	}
}

func initProcMap() map[string]*ProcMeta {
	pentry := make(map[string]*ProcMeta)

	return pentry
}

func (pi *ProcEntry) constructProcMap() error {

	pmap := make(map[string]*ProcMeta)

	files, err := filepath.Glob(PROCSTATUS)
	if err != nil {
		log.Error("Error in reading dir ", err.Error())
		return err
	}

	for _, file := range files {
		pmeta := getProcMeta()

		content, err := ioutil.ReadFile(file)
		if err != nil {
			log.Error("Error while reading ", err.Error())
			return err
		}

		entry := strings.Split(string(content), "\n")
		for field, _ := range pmeta.procAttrs {
			for _, line := range entry {
				attrs := strings.Split(line, ":")
				if attrs[0] == field {
					pmeta.procAttrs[field] = attrs[1]
				}
			}
		}
		pmap[pmeta.procAttrs["Pid"]] = &pmeta
	}
	pi.procMap = pmap

	return nil
}

func (pi *ProcEntry) getProcMap() (map[string]*ProcMeta, error) {

	if err := pi.constructProcMap(); err != nil {
		return nil, err
	}
	return pi.procMap, nil
}

func (pi *ProcEntry) Watch(pid uint64) error {

	eh, err := NewEventHandler()
	if err != nil {
		return err
	}

	if err = eh.Notify(pid); err != nil {
		return err
	}

	return nil
}

func (pi *ProcEntry) List() error {
	pmap, err := pi.getProcMap()
	if err != nil {
		return err
	}

	for _, pr := range pmap {
		for f, val := range pr.procAttrs {
			fmt.Printf("%s\t%s\n", f, val)
		}
		fmt.Println()
	}

	return nil
}
