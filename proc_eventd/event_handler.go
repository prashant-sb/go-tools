package main
import (
	"golang.org/x/sys/unix"
	log "github.com/golang/glog"
)

type Events struct {
}

type EventHandler interface {
	Notify(uint64) error
}

func NewEvent() (*Events, error) {
	e := Events{}
	return &e, nil
}

func NewEventHandler() (EventHandler, error) {
	return NewEvent()
}

func (e *Events) Notify(pid uint64) error {

	nlSock, err := unix.Socket(unix.AF_NETLINK, unix.SOCK_DGRAM, unix.AF_INET)
	if err != nil {
		log.Error("Error on creating the socket: %v", err)
		return err
	}
	log.Info("Socket opened .. ")

	unix.Close(nlSock)

	return nil
}
