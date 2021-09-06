//
// Command line ftp client which supports following options
//

package main

import (
	ftpcli "github.com/prashant-sb/go-utils/ftp_client/cli"
	"go.uber.org/zap"
)

var opt *ftpcli.CommandArgs = nil

// TODO: Change logger
var Log *zap.SugaredLogger = nil

func init() {
	logger, _ := zap.NewProduction()
	Log = logger.Sugar()
	opt = ftpcli.NewCmdArgs()

	defer logger.Sync()
}

func main() {
	if err := opt.Sanitize(); err != nil {
		Log.Error(err, "Error in running ftp client operation")
		return
	}

	if err := opt.Run(); err != nil {
		Log.Error(err, "Error in running ftp client operation")
	}
}
