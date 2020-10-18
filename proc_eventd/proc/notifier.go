package proc

import (
	"errors"

	log "github.com/golang/glog"
)

type ProcEventFork struct {
	ParentPid uint64 // Pid of the process that called fork()
	ChildPid  uint64 // Child process pid created by fork()
}

type ProcEventExec struct {
	Pid uint64 // Pid of the process that called exec()
}

type ProcEventExit struct {
	Pid uint64 // Pid of the process that called exit()
}

type watch struct {
	flags uint32 // Saved value of Watch() flags param
}

type eventListener interface {
	close() error // Watch.Close() closes the OS specific listener
}

type Watcher struct {
	listener eventListener       // Listener interface, netlink or others
	watches  map[uint64]*watch   // Map of watched process ids
	Error    chan error          // Errors are sent on this channel
	Fork     chan *ProcEventFork // Fork events are sent on this channel
	Exec     chan *ProcEventExec // Exec events are sent on this channel
	Exit     chan *ProcEventExit // Exit events are sent on this channel
	done     chan bool           // Used to stop the readEvents() goroutine
	isClosed bool                // Set to true when Close() is first called
}

// Initialize event listener and channels
func NewWatcher() (*Watcher, error) {
	listener, err := CreateListener()

	if err != nil {
		return nil, err
	}

	w := &Watcher{
		listener: listener,
		watches:  make(map[uint64]*watch),
		Fork:     make(chan *ProcEventFork),
		Exec:     make(chan *ProcEventExec),
		Exit:     make(chan *ProcEventExit),
		Error:    make(chan error),
		done:     make(chan bool, 1),
	}

	go w.readEvents()
	return w, nil
}

// Close event channels when done message is received
func (w *Watcher) finish() {
	close(w.Fork)
	close(w.Exec)
	close(w.Exit)
	close(w.Error)
}

// Closes the OS specific event listener,
// removes all watches and closes all event channels.
func (w *Watcher) Close() error {
	if w.isClosed {
		return nil
	}
	w.isClosed = true

	for pid := range w.watches {
		err := w.RemoveWatch(pid)
		if err != nil {
			log.Error("Error in removing watch for pid: ", pid)
		}
	}
	w.done <- true

	return w.listener.close()

}

// Add pid to the watched process set.
// The flags param is a bitmask of process events to capture,
// must be one or more of: PROC_EVENT_FORK, PROC_EVENT_EXEC, PROC_EVENT_EXIT
func (w *Watcher) Watch(pid uint64, flags uint32) error {
	if w.isClosed {
		return errors.New("watcher is closed")
	}

	watchEntry, found := w.watches[pid]

	if found {
		watchEntry.flags |= flags
	} else {
		w.watches[pid] = &watch{flags: flags}
	}

	return nil
}

// Remove pid from the watched process set.
func (w *Watcher) RemoveWatch(pid uint64) error {
	_, ok := w.watches[pid]
	if !ok {
		log.Error("Watch does not exist for %v", pid)
		return errors.New("watch for pid does not exist")
	}
	delete(w.watches, pid)

	return nil
}

// Internal helper to check if there is a message on the "done" channel.
// The "done" message is sent by the Close() method; when received here,
// the Watcher.finish method is called to close all channels and return
// true - in which case the caller should break from the readEvents loop.
func (w *Watcher) isDone() bool {
	var done bool
	select {
	case done = <-w.done:
		w.finish()
	default:
	}
	return done
}
