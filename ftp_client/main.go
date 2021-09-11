//
// Command line ftp client which supports following options
//

package main

import (
	ftpcli "github.com/prashant-sb/go-utils/ftp_client/cli"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

var opt ftpcli.CommandArgs = nil

// TODO: Change logger
var Log *zap.Logger = nil

func init() {
	opt = ftpcli.NewCmdArgs()

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
}

func main() {
	if err := opt.Sanitize(); err != nil {
		log.Errorf("Error in running ftp client operations: %s", err)
		return
	}

	if err := opt.Run(); err != nil {
		log.Error(err, "Error in running ftp client operations: %s", err)
	}
}
