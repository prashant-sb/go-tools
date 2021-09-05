package ftpcli

import (
	"errors"
	"flag"

	"github.com/prashant-sb/go-utils/ftp_client/ftp"
)

type CommandArgs map[string]Arg

type Arg struct {
	Name    string
	Desc    string
	Short   string
	Default interface{}
	Value   interface{}
}

func NewCmdArgs() *CommandArgs {
	return &CommandArgs{
		"server": {
			Name:    "server",
			Desc:    "Address of FTP Server",
			Short:   "s",
			Default: "localhost",
		},
		"list": {
			Name:    "list",
			Desc:    "List directory from FTP Server",
			Short:   "l",
			Default: "/",
		},
		"upload": {
			Name:    "upload",
			Desc:    "Local file or dir path",
			Short:   "u",
			Default: "",
		},
		"download": {
			Name:    "download",
			Desc:    "Local file or dir path",
			Short:   "d",
			Default: "",
		},
		"user": {
			Name:    "user",
			Desc:    "FTP User name",
			Short:   "f",
			Default: "n",
		},
		"port": {
			Name:    "port",
			Desc:    "FTP port for connection",
			Short:   "p",
			Default: "",
		},
		"recurse": {
			Name:    "recurse",
			Desc:    "Traverse all embeded directories",
			Short:   "r",
			Default: false,
		},
	}
}

func (c *CommandArgs) Sanitize() error {
	for name, arg := range *c {
		// Get correct default type
		switch arg.Default.(type) {

		case bool:
			var val = flag.Bool(name, arg.Default.(bool), arg.Desc)
			flag.BoolVar(val, arg.Short, arg.Default.(bool), arg.Desc)
			arg.Value = val

		case string:
			var val = flag.String(name, arg.Default.(string), arg.Desc)
			flag.StringVar(val, arg.Short, arg.Default.(string), arg.Desc)
			arg.Value = val
		}
	}

	flag.Parse()

	// TODO: Add option validations

	return nil
}

func (c *CommandArgs) getServerParams() (*ftp.Cred, error) {
	cmd := *c
	if _, exists := cmd["server"]; !exists {
		return nil, errors.New("server param not found")
	}

	if _, exists := cmd["port"]; !exists {
		return nil, errors.New("port param not found")
	}

	if _, exists := cmd["user"]; !exists {
		return nil, errors.New("user param not found")
	}

	return &ftp.Cred{
		User:   cmd["user"].Value.(string),
		Port:   cmd["port"].Value.(string),
		Server: cmd["server"].Value.(string),
	}, nil
}

func (c *CommandArgs) Request() *ftp.Request {
	cmd := *c
	if _, exists := cmd["upload"]; exists {
		return &ftp.Request{
			Type:    ftp.UploadRequest,
			Target:  cmd["upload"].Value.(string),
			Recurse: cmd["recurse"].Value.(bool),
		}
	}
	if _, exists := cmd["download"]; exists {
		return &ftp.Request{
			Type:    ftp.DownloadRequest,
			Target:  cmd["download"].Value.(string),
			Recurse: cmd["recurse"].Value.(bool),
		}
	}
	if _, exists := cmd["list"]; exists {
		return &ftp.Request{
			Type:    ftp.ListRequest,
			Target:  cmd["list"].Value.(string),
			Recurse: cmd["recurse"].Value.(bool),
		}
	}

	return &ftp.Request{
		Type: ftp.UnknownRequest,
	}
}

func (c *CommandArgs) Run() error {
	cred, err := c.getServerParams()
	if err != nil {
		return err
	}

	conn, err := ftp.NewConnection(cred)
	if err != nil {
		return err
	}
	defer ftp.Close(conn)

	req := c.Request()
	switch req.Type {
	case ftp.ListRequest:
		if err := ftp.List(conn, req); err != nil {
			return err
		}

	case ftp.UploadRequest:
		if err := ftp.Upload(conn, req); err != nil {
			return err
		}

	case ftp.DownloadRequest:
		if err := ftp.Download(conn, req); err != nil {
			return err
		}

	default:
		return errors.New("Unknown request type")
	}

	return nil
}
