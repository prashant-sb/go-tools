package proc

import (
	"os/exec"
	"strconv"
	"strings"
	"testing"

	prc "github.com/prashant-sb/go-tools/proc_eventd/proc"
)

const (
	listProc string = "ps -e --no-header | wc -l"
)

func TestList(t *testing.T) {
	procIter := prc.NewProcIterator()

	pmap, err := procIter.GetProcMap()
	if err != nil {
		t.Errorf("List() FAILED with %v", err.Error())
	}
	procCmd := exec.Command("/bin/bash", "-c", listProc)
	stdout, err := procCmd.Output()
	if err != nil {
		t.Errorf("List() FAILED with %v", err.Error())
		return
	}

	pout := strings.TrimSpace(string(stdout))
	tprocs, err := strconv.Atoi(pout)
	if err != nil {
		t.Errorf("List() FAILED: error %v", err.Error())
		return
	}

	if len(pmap)+2 != tprocs {
		t.Errorf("List() FAILED: Process count error %v:%v", tprocs, len(pmap)+2)
		return
	}

	t.Logf("List() PASSED")
}

func TestWatch(t *testing.T) {
	t.Logf("Watch() PASSED")
}
