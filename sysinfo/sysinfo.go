package main

import (
	"os"
	"errors"
	hw "github.com/jaypipes/ghw"
	log "github.com/golang/glog"
)

type sysinfo struct{
	Cpu		hw.CPUInfo 
	Bios	hw.BIOSInfo
	Memory	hw.MemoryInfo
	Storage hw.BlockInfo
	Network hw.NetworkInfo
}

type fileopt struct {
	info	 *sysinfo
	fileType string
	filePath string
}

type filefmt interface{
	get() (string, error)
	To(string) error
	FormatAs(string)
}

type informer interface{
	String(string) (string, error)
	FmtOption() filefmt
}

func NewInformer() (informer, error) {
	sysinfo, err := collectSysinfo()
	if err != nil {
		return nil, err
	}
	
	return sysinfo, nil
}

func (s *sysinfo) String(ft string) (string, error){
	var tostr string

	switch	{
	case ft == "json":
		tostr = s.Cpu.JSONString(true) + 
			s.Bios.JSONString(true)	+
			s.Memory.JSONString(true) +	
			s.Storage.JSONString(true) +
			s.Network.JSONString(true)
	
	case ft == "yaml":
		tostr = s.Cpu.YAMLString() + 
			s.Bios.YAMLString()	+
			s.Memory.YAMLString() +	
			s.Storage.YAMLString() +
			s.Network.YAMLString()
	default:
		return tostr, errors.New("Type not supported")
	}

	return tostr, nil
}

func (s *sysinfo) FmtOption() filefmt {
	ft := fileopt{
		filePath: "sysinfo.json",
		fileType: "json",
		info: s,
	}
	return &ft
}

func (opt *fileopt) FormatAs(ftype string) {
	opt.fileType = ftype
}

func (opt *fileopt) get() (string, error) {
	sinfo := opt.info
	ft := opt.fileType

	ftstr, err := sinfo.String(ft)
	if err != nil {
		return ftstr, err
	}
	
	return ftstr, nil
}

func (opt *fileopt) To(file string) error {
	var info string
	var err error
		
	info, err = opt.get()
	if err != nil {
		return err
	}

	f, err := os.Create(file)
	defer f.Close()
	if err != nil {
		return err
	}
	_, err = f.WriteString(info)
	if err != nil {
		return err
	}

	return nil
}

func collectSysinfo() (*sysinfo, error) {
	sinfo := sysinfo{}	

	Mem, err := hw.Memory()
	if err != nil {
		log.Info("Error in collecting Memory info: ", err.Error())
		return nil, err
	}
	sinfo.Memory = *Mem

	Cpu, err := hw.CPU()
	if err != nil {
		log.Info("Error in collecting Cpu info: ", err.Error())
		return nil, err
	}
	sinfo.Cpu = *Cpu

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