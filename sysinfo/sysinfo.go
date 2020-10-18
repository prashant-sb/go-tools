package main

import (
	"errors"
	log "github.com/golang/glog"
	hw "github.com/jaypipes/ghw"
	"os"
)

// Defines the system details
type sysinfo struct {
	CPU     hw.CPUInfo
	Bios    hw.BIOSInfo
	Memory  hw.MemoryInfo
	Storage hw.BlockInfo
	Network hw.NetworkInfo
}

// Informer : interface for sysinfo
type Informer interface {
	String(string) (string, error)
	FmtOption() filefmt
}

// File formatting options
type fileopt struct {
	info     *sysinfo
	fileType string
	filePath string
}

// Interface for fileopt
type filefmt interface {
	get() (string, error)
	To(string) error
	FormatAs(string)
}

// NewInformer returns interface binding to sysinfo
func NewInformer() (Informer, error) {
	sysinfo, err := collectSysinfo()
	if err != nil {
		return nil, err
	}

	return sysinfo, nil
}

// FmtOption binds filefmt interface to
// file formatting options
func (s *sysinfo) FmtOption() filefmt {
	ft := fileopt{
		filePath: "sysinfo.json",
		fileType: "json",
		info:     s,
	}
	return &ft
}

// FormatAs sets the formatting type on option
func (opt *fileopt) FormatAs(ftype string) {
	opt.fileType = ftype
}

// get is filefmt interface method,
// formats sysinfo based on option
func (opt *fileopt) get() (string, error) {
	sinfo := opt.info
	ft := opt.fileType

	ftstr, err := sinfo.String(ft)
	if err != nil {
		return ftstr, err
	}

	return ftstr, nil
}

// To saves the formatted output to
// given file.
func (opt *fileopt) To(file string) error {
	var info string
	var err error

	info, err = opt.get()
	if err != nil {
		return err
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(info)
	if err != nil {
		return err
	}

	return nil
}

// Stringer for sysinfo, supports yaml | json
func (s *sysinfo) String(ft string) (string, error) {
	var tostr string

	switch {
	case ft == "json":
		tostr = s.CPU.JSONString(true) +
			s.Bios.JSONString(true) +
			s.Memory.JSONString(true) +
			s.Storage.JSONString(true) +
			s.Network.JSONString(true)

	case ft == "yaml":
		tostr = s.CPU.YAMLString() +
			s.Bios.YAMLString() +
			s.Memory.YAMLString() +
			s.Storage.YAMLString() +
			s.Network.YAMLString()
	default:
		return tostr, errors.New("Type not supported")
	}

	return tostr, nil
}

// collectSysinfo prepares the sysinfo
func collectSysinfo() (*sysinfo, error) {
	sinfo := sysinfo{}

	Mem, err := hw.Memory()
	if err != nil {
		log.Info("Error in collecting Memory info: ", err.Error())
		return nil, err
	}
	sinfo.Memory = *Mem

	CPU, err := hw.CPU()
	if err != nil {
		log.Info("Error in collecting CPU info: ", err.Error())
		return nil, err
	}
	sinfo.CPU = *CPU

	Bios, err := hw.BIOS()
	if err != nil {
		log.Info("Error in collecting Bios info: ", err.Error())
		return nil, err
	}
	sinfo.Bios = *Bios

	block, err := hw.Block()
	if err != nil {
		log.Info("Error in collecting Storage info: ", err.Error())
		return nil, err
	}
	sinfo.Storage = *block

	net, err := hw.Network()
	if err != nil {
		log.Info("Error in collecting Network info: ", err.Error())
		return nil, err
	}
	sinfo.Network = *net

	return &sinfo, nil
}
