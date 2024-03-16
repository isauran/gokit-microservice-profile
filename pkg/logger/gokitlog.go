package logger

import (
	"io"

	"github.com/go-kit/log"
)

// logger := GoKitJSONLogger(os.Stdout)
func GoKitJSONLogger(w io.Writer) log.Logger {
	l := log.NewJSONLogger(w)
	l = log.With(l, "time", log.DefaultTimestamp, "caller", log.DefaultCaller)
	return l
}
