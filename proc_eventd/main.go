package main

import "flag"

var (
	list = flag.String("list", false, "List running processes")
	wpid = flag.String("watch", "", "watch process by pid for events")
)

func main() {
	flag.Parse()

	pitr = NewProcIterator()
	if *list == true {
		pitr.List()
	}

	if *wpid != nil {
		pitr.Watch(pid)
	}

}
