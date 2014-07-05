package logging

import (
	stdlog "log"
	"github.com/op/go-logging"
	"os"
)

var Log = logging.MustGetLogger("feedupdater")

func Init() {
	logging.SetFormatter(logging.MustStringFormatter("%{level} %{id} %{message}"))
	logBackend := logging.NewLogBackend(os.Stderr, "", stdlog.LstdFlags|stdlog.Lshortfile)
	logging.SetBackend(logBackend)
}