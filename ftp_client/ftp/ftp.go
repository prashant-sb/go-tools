package ftp

import (
	"bufio"
	"context"
	"io"
	"net"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

type TargetType int
type RequestType int

const (
	ListRequest RequestType = iota
	UploadRequest
	DownloadRequest
	UnknownRequest
)

type Cred struct {
	User    string
	Server  string
	Port    string
	Timeout time.Duration
}

type Request struct {
	Target    string
	Recurse   bool
	Type      RequestType
	EntryType TargetType
}

type dialOptions struct {
	dialer net.Dialer
	conn   net.Conn
	cred   *Cred
	host   string
}

type FtpConnect struct {
	options *dialOptions
	writer  io.Writer
}

type DialOption struct {
	setup func(do *dialOptions)
}

func DefaultOptions(cr *Cred) DialOption {
	return DialOption{func(do *dialOptions) {
		do.conn = nil
		do.host = "localhost"
		do.cred = cr
	}}
}

func NewConnection(options ...DialOption) (*FtpConnect, error) {
	dopt := &dialOptions{}
	for _, opt := range options {
		opt.setup(dopt)
	}

	ctx, cancel := context.WithTimeout(context.Background(), dopt.cred.Timeout)
	defer cancel()

	tconn, err := dopt.dialer.DialContext(ctx, "tcp", dopt.cred.Server+":"+dopt.cred.Port)
	if err != nil {
		return nil, err
	}
	rAddr := tconn.RemoteAddr().(*net.TCPAddr)
	dopt.host = rAddr.IP.String()
	dopt.conn = tconn

	log.Info("FTP connected: ", dopt.cred.Server)
	return &FtpConnect{
		options: dopt,
		writer:  bufio.NewWriter(os.Stdout),
	}, nil
}

func Close(ftpConn *FtpConnect) {
	ftpConn.options.conn.Close()
}

func Download(ftpConn *FtpConnect, req *Request) error {
	return nil
}

func Upload(ftpConn *FtpConnect, req *Request) error {
	return nil
}

func List(ftpConn *FtpConnect, req *Request) error {
	return nil
}
