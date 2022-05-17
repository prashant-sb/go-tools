package simple

import (
	"os"
	"os/signal"
	"os/user"
	"sync"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
)

const defaultSig = syscall.SIGHUP

type SigAction struct {
	sigChannel    chan os.Signal
	sigCompletion chan int
	sigHandler    func(*RoutineArgs)
}

type RoutineArgs struct {
	start       time.Time
	end         time.Time
	currentUser user.User
}

type Runner struct {
	args   *RoutineArgs
	lock   *sync.Mutex
	sigact *SigAction
}

func NewHandler(sig syscall.Signal, actFunc func(*RoutineArgs)) *SigAction {
	sigact := &SigAction{
		sigChannel:    make(chan os.Signal, 1),
		sigCompletion: make(chan int),
		sigHandler:    actFunc,
	}
	signal.Notify(sigact.sigChannel, sig)

	return sigact
}

func SigHandler(args *RoutineArgs) {
	log.Info("Signal handler called with args:", *args)
}

func NewRunner() *Runner {

	return &Runner{
		args: &RoutineArgs{
			start: time.Now(),
		},
		sigact: NewHandler(defaultSig, SigHandler),
	}
}

func (r *Runner) run() {
	log.Info("start time")
	log.Info("current time")
}

func (r *Runner) Start() {
	log.Info("Start")
	go r.run()
}

func (r *Runner) Restart() {
	log.Info("Restart called on runner")
}

func (r *Runner) Stop() {
	log.Info("Stopping routine now")
}
