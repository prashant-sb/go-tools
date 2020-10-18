package proc

import (
	log "github.com/golang/glog"
)

// Watchable events
type Events struct {
	all  uint32
	exec uint32
	fork uint32
	exit uint32
}

// User level interface for events
type EventHandler interface {
	Notify(uint64) error
}

// Initialize event
func NewEvent() (*Events, error) {
	e := Events{
		all:  PROC_EVENT_ALL,
		exec: PROC_EVENT_EXEC,
		fork: PROC_EVENT_FORK,
		exit: PROC_EVENT_EXIT,
	}

	return &e, nil
}

// Returns interface binded with Event type
func NewEventHandler() (EventHandler, error) {
	return NewEvent()
}

// Create and watch for given pid.
// Waits till exit.
// TODO: Add methods for subscriptions for Events
func (e *Events) Notify(pid uint64) error {

	notif, err := NewWatcher()
	if err != nil {
		log.Error("Error occured in creating process watcher: ", err.Error())
		return err
	}

	defer notif.Close()
	done := make(chan bool)

	err = notif.Watch(pid, e.all)
	if err != nil {
		return err
	}

	// Process events
	log.Info("Watching pid: ", pid)
	go func() {
		for {
			select {
			case ev := <-notif.Fork:
				log.Info("Fork event:", *ev)
			case ev := <-notif.Exec:
				log.Info("Exec event:", *ev)
			case ev := <-notif.Exit:
				log.Info("Exit event:", *ev)
			case err := <-notif.Error:
				if err != nil {
					log.Error("Error:", err)
				}
			}
		}
	}()
	<-done

	return nil
}
