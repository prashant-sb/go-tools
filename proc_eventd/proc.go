package main

type ProcOps interface {
	List() ProcMeta
	Watch(pid uint16) error
}

type ProcIter interface {
	List()
	Watch(pid uint16)
}

type ProcEntry struct {
	proclist map[uint16]ProcMeta
}

type ProcMeta struct {
	pid  uint16
	uid  uint16
	gid  uint16
	ppid uint16
	name string
}

func NewProcIterator() ProcEntry {
	return nil
}
