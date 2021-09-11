package cli

import (
	"errors"
	"flag"
	"os"
	"time"

	"github.com/prashant-sb/go-utils/ftp_client/ftp"
)

type CommandArgs map[string]*Arg

type Arg struct {
	Name    string
	Desc    string
	Short   string
	Default interface{}
	Value   interface{}
}

func NewCmdArgs() CommandArgs {
	return CommandArgs{
		"server": &Arg{
			Name: "server",
			Desc: "Address of FTP Server",
			//Short:   "s",
			Default: "",
		},
		"list": &Arg{
			Name: "list",
			Desc: "List directory from FTP Server",
			//Short:   "l",
			Default: "",
		},
		"upload": &Arg{
			Name: "upload",
			Desc: "Local file or dir path",
			//Short:   "u",
			Default: "",
		},
		"download": &Arg{
			Name: "download",
			Desc: "Local file or dir path",
			//Short:   "d",
			Default: "",
		},
		"user": &Arg{
			Name: "user",
			Desc: "FTP User name",
			//Short:   "n",
			Default: "",
		},
		"password": &Arg{
			Name: "password",
			Desc: "Password for FTP User",
			//Short:   "p",
			Default: "",
		},
		"port": &Arg{
			Name: "port",
			Desc: "FTP port for connection",
			//Short:   "p",
			Default: "",
		},
		"recurse": &Arg{
			Name: "recurse",
			Desc: "Traverse all embeded directories",
			//Short:   "r",
			Default: false,
		},
	}
}

func (c CommandArgs) Sanitize() error {
	for name, arg := range c {
		// Get correct default type
		switch arg.Default.(type) {

		case bool:
			arg.Value = flag.Bool(name, arg.Default.(bool), arg.Desc)
			if len(arg.Short) > 0 {
				flag.BoolVar(arg.Value.(*bool), arg.Short, arg.Default.(bool), arg.Desc)
			}

		case string:
			arg.Value = flag.String(name, arg.Default.(string), arg.Desc)
			if len(arg.Short) > 0 {
				flag.StringVar(arg.Value.(*string), arg.Short, arg.Default.(string), arg.Desc)
			}
		}
	}
	if len(os.Args) == 1 {
		flag.Usage()
		return errors.New("Arguments required")
	}
	flag.Parse()

	// TODO: Add option validations
	return nil
}

func (cmd CommandArgs) getStringVal(key string) (string, error) {
	einval := errors.New("Invalid parameters")
	if _, exists := cmd[key]; !exists {
		return "", einval
	}

	if val, ok := cmd[key].Value.(*string); ok {
		return *val, nil
	}

	return "", einval
}

func (cmd CommandArgs) getServerParams() (*ftp.Cred, error) {
	var err error = nil
	var val string

	cred := &ftp.Cred{
		Port:    "21",
		Timeout: time.Second * 30,
	}

	if val, err = cmd.getStringVal("server"); err == nil {
		cred.Server = val
	}

	if val, err = cmd.getStringVal("user"); err == nil {
		cred.User = val
	}

	if err != nil {
		return nil, err
	}

	if val, err = cmd.getStringVal("port"); err == nil {
		cred.Port = val
	}

	return cred, nil
}

func (cmd CommandArgs) Request() *ftp.Request {

	if _, exists := cmd["upload"]; exists {
		return &ftp.Request{
			Type:    ftp.UploadRequest,
			Target:  *cmd["upload"].Value.(*string),
			Recurse: *cmd["recurse"].Value.(*bool),
		}
	}
	if _, exists := cmd["download"]; exists {
		return &ftp.Request{
			Type:    ftp.DownloadRequest,
			Target:  *cmd["download"].Value.(*string),
			Recurse: *cmd["recurse"].Value.(*bool),
		}
	}
	if _, exists := cmd["list"]; exists {
		return &ftp.Request{
			Type:    ftp.ListRequest,
			Target:  *cmd["list"].Value.(*string),
			Recurse: *cmd["recurse"].Value.(*bool),
		}
	}

	return &ftp.Request{
		Type: ftp.UnknownRequest,
	}
}

func (c CommandArgs) Run() error {
	var err error = nil
	cred, err := c.getServerParams()
	if err != nil {
		return err
	}

	conn, err := ftp.NewConnection(ftp.DefaultOptions(cred))
	if err != nil {
		return err
	}
	defer ftp.Close(conn)

	req := c.Request()
	switch req.Type {
	case ftp.ListRequest:
		err = ftp.List(conn, req)

	case ftp.UploadRequest:
		err = ftp.Upload(conn, req)

	case ftp.DownloadRequest:
		err = ftp.Download(conn, req)

	default:
		err = errors.New("Unknown request type")
	}

	return err
}
