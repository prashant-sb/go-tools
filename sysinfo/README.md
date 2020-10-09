##### Collect system information #####
----
Utility collects system information about following components

- CPU
- BIOS
- Memory
- Block storage
- Network Interface

----
##### Options #####
```
Usage of ./run:
  -alsologtostderr
    	log to standard error as well as files
  -format string
    	Prints the information in json | pretty format (default "json")
  -log_backtrace_at value
    	when logging hits line file:N, emit a stack trace
  -log_dir string
    	If non-empty, write log files in this directory
  -logtostderr
    	log to standard error instead of files
  -saveas string
    	Saves the information in given file or stdout (default "/dev/stdout")
  -stderrthreshold value
    	logs at or above this threshold go to stderr
  -v value
    	log level for V logs
  -vmodule value
    	comma-separated list of pattern=N settings for file-filtered logging
```
----

- Collecting info in json
    ```
    ./run -format json -saveas info.json

    ```
- Collecting info in yaml
    ```
    ./run -format yaml -saveas info.yaml

    ```

- References

    Utility based upon implementation github.com/jaypipes/ghw

----
