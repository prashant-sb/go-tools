package proc

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	log "github.com/golang/glog"
)

// Process data structures parsed from /proc
const (
	PROCDIR    = "/proc/"
	PROCSTATUS = PROCDIR + "*/status"
)

// Iterator interface to list and watch process
type ProcIter interface {
	List() error
	Watch(pid uint64) error
	GetProcMap() (map[string]*ProcMeta, error)
}

// Process attribute map
// with fields specified as key
type ProcMeta struct {
	procAttrs map[string]string
}

// Process map of pid and Process Attribute values
type ProcEntry struct {
	procMap map[string]*ProcMeta
}

// Init Iterator interface
func NewProcIterator() ProcIter {

	return &ProcEntry{
		procMap: InitProcMap(),
	}
}

// Fields that are listable.
// Initialize Process attribute map with empty values
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

// Stub to init ProcEntry
func InitProcMap() map[string]*ProcMeta {
	pentry := make(map[string]*ProcMeta)

	return pentry
}

// Contructs the map with fields from Process Attributes
// and saves with key as Pid with ProcEntry.
// TODO: Add ProcessAttr processing to diffrent method
// so that, it could be useful for getting details for process
// with given Pid
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
		for field := range pmeta.procAttrs {
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

// contruct and returns the process map at the instance.
func (pi *ProcEntry) GetProcMap() (map[string]*ProcMeta, error) {

	if err := pi.constructProcMap(); err != nil {
		return nil, err
	}
	return pi.procMap, nil
}

// Init and calls the process watcher
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

// List the details of process map
func (pi *ProcEntry) List() error {
	pmap, err := pi.GetProcMap()
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
