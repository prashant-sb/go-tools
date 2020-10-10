## Display or Create Users in Linux

Display linux user information or create / delete user.

### Usage

```
Usage of ./run:
  -alsologtostderr
    	log to standard error as well as files
  -create
    	Creates the system user
  -delete
    	Deletes the system user
  -from string
    	Json configuration for create user
  -list
    	Lists the system users
  -log_backtrace_at value
    	when logging hits line file:N, emit a stack trace
  -log_dir string
    	If non-empty, write log files in this directory
  -logtostderr
    	log to standard error instead of files
  -stderrthreshold value
    	logs at or above this threshold go to stderr
  -user string
    	List specific system user
  -v value
    	log level for V logs
  -vmodule value
    	comma-separated list of pattern=N settings for file-filtered logging

```
#### Add new user

```
./run -create -from ./usr.json -logtostderr

Enter Password for test: 
User test added

Example usr.json :
{
   "uid": "0",
   "gid": "0",
   "userName": "test",
   "groupName": "syslog",
   "name": "Test User",
   "homeDir": "/home/test"
}
```
#### User information

```
./run -logtostderr -list -user test
{
   "uid": "1002",
   "gid": "1002",
   "userName": "test",
   "groupName": "test",
   "homeDir": "/home/test"
}
```

#### List all users
```
./run -logtostderr -list
{
   "users": [
      {
         "uid": "0",
         "gid": "0",
         "userName": "root",
         "groupName": "root",
         "name": "root",
         "homeDir": "/root"
      },
      {
         "uid": "1",
         "gid": "1",
         "userName": "daemon",
         "groupName": "daemon",
         "name": "daemon",
         "homeDir": "/usr/sbin"
      },
      {
         "uid": "2",
         "gid": "2",
         "userName": "bin",
         "groupName": "bin",
         "name": "bin",
         "homeDir": "/bin"
      },
...
...
```
#### Delete user

```
./run -logtostderr -delete -user test
test user deleted.
```
