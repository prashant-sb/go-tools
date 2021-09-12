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

type FtpConnector struct {
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

func NewConnection(options ...DialOption) (*FtpConnector, error) {
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
	return &FtpConnector{
		options: dopt,
		writer:  bufio.NewWriter(os.Stdout),
	}, nil
}

func Close(ftpConn *FtpConnector) {
	ftpConn.options.conn.Close()
}

func Download(ftpConn *FtpConnector, req *Request) error {
	return nil
}

func Upload(ftpConn *FtpConnector, req *Request) error {
	return nil
}

func List(ftpConn *FtpConnector, req *Request) error {
	return nil
}

func Delete(ftpConn *FtpConnector, req *Request) error {
	return nil
}

// -----
func (f *FtpConnector) login(user, password string) error {
	return nil
}

func (f *FtpConnector) logout() error {
	return nil
}

func (f *FtpConnector) changeDir(to string) (string, error) {
	return "", nil
}

func (f *FtpConnector) currentDir() (string, error) {
	return "", nil
}

func (f *FtpConnector) createDir(path string) error {
	return nil
}

func (f *FtpConnector) removeDir(path string) error {
	return nil
}

func (f *FtpConnector) rename(from, to string) error {
	return nil
}

func (f *FtpConnector) retr(path string) error {
	return nil
}

func (f *FtpConnector) stor(path string, r io.Reader) error {
	return nil
}

func (f *FtpConnector) quit() error {
	return nil
}
