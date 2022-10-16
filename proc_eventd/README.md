### List and Watch processes for events

Following are the events that are watchable - 

- EXEC
- FORK
- EXIT
- ALL

```
Usage of ./run:
  -alsologtostderr
    	log to standard error as well as files
  -list
    	List running processes
  -log_backtrace_at value
    	when logging hits line file:N, emit a stack trace
  -log_dir string
    	If non-empty, write log files in this directory
  -logtostderr
    	log to standard error instead of files
  -stderrthreshold value
    	logs at or above this threshold go to stderr
  -v value
    	log level for V logs
  -vmodule value
    	comma-separated list of pattern=N settings for file-filtered logging
  -watch uint
    	watch process by pid for events
```

##### Listing processes

```
./run -list

Pid             1520
PPid            1
Tgid            1520
State           S (sleeping)
Umask           0002
Threads         1
Name            systemd

Pid             670
PPid            1
Tgid            670
State           S (sleeping)
Umask           0022
Threads         1
Name            networkd-dispat

...
```
##### Watching process events

Run with -watch <pid> flag

```
./run -watch 2086 -logtostderr
I1009 19:33:11.777662    5866 event_handler.go:57] Watching pid: 2086
I1009 19:33:15.846381    5866 event_handler.go:62] Fork event:{2086 5872}
I1009 19:33:15.847285    5866 event_handler.go:64] Exec event:{5872}
I1009 19:33:15.849077    5866 event_handler.go:62] Fork event:{5872 5873}
I1009 19:33:15.849360    5866 event_handler.go:64] Exec event:{5873}
I1009 19:33:15.849720    5866 event_handler.go:62] Fork event:{5873 5875}
I1009 19:33:15.849833    5866 event_handler.go:64] Exec event:{5875}
I1009 19:33:15.850188    5866 event_handler.go:66] Exit event:{5875}
I1009 19:33:15.850406    5866 event_handler.go:62] Fork event:{5873 5876}
I1009 19:33:15.850469    5866 event_handler.go:62] Fork event:{5876 5877}
I1009 19:33:15.850670    5866 event_handler.go:64] Exec event:{5877}
I1009 19:33:15.851043    5866 event_handler.go:66] Exit event:{5877}
I1009 19:33:15.851116    5866 event_handler.go:66] Exit event:{5876}
I1009 19:33:15.851209    5866 event_handler.go:66] Exit event:{5873}
I1009 19:33:15.851570    5866 event_handler.go:62] Fork event:{5872 5878}
I1009 19:33:15.851886    5866 event_handler.go:64] Exec event:{5878}
I1009 19:33:15.852161    5866 event_handler.go:66] Exit event:{5878}
I1009 19:33:23.506551    5866 event_handler.go:66] Exit event:{5872}

```
