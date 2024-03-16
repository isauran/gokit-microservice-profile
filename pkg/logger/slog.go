package logger

import (
	"fmt"
	"io"
	"log/slog"
	"path/filepath"
)

// logger := SlogJSONLogger(os.Stdout, slog.LevelInfo)
func SlogJSONLogger(w io.Writer, level slog.Level) *slog.Logger {
	opts := slog.HandlerOptions{
		AddSource: true,
		Level:     &level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.SourceKey {
				s, _ := a.Value.Any().(*slog.Source)
				if s != nil {
					return slog.String("caller", fmt.Sprintf("%s:%d", filepath.Base(s.File), s.Line))
				}
			}
			return a
		},
	}
	h := slog.NewJSONHandler(w, &opts)
	l := slog.New(h)
	slog.SetDefault(l)
	return l
}
