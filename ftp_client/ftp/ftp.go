package ftp

import (
	"io"
	"net"
	"time"
)

type RequestType int
type TargetType int

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

type DialOptions struct {
	dialer net.Dialer
	conn   net.Conn
	cred   *Cred
	host   string
}

type FtpConnect struct {
	cred    *Cred
	options *DialOptions
	writer  io.Writer
}

func NewConnection(cred *Cred, dialer ...DialOptions) (*FtpConnect, error) {
	return nil, nil
}

func Close(ftpConn *FtpConnect) {
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
