package logging

import (
	"go-web-scraper/internal/config"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Info(msg string)
	Error(msg string)
	Debug(msg string)
	WithFields(fields map[string]interface{}) Logger
	WithError(err error) Logger
}

type logrusAdapter struct {
	entry *logrus.Entry
}

func (l *logrusAdapter) Info(msg string) {
	l.entry.Info(msg)
}

func (l *logrusAdapter) Error(msg string) {
	l.entry.Error(msg)
}

func (l *logrusAdapter) Debug(msg string) {
	l.entry.Debug(msg)
}

func (l *logrusAdapter) WithFields(fields map[string]interface{}) Logger {
	return &logrusAdapter{entry: l.entry.WithFields(fields)}
}

func (l *logrusAdapter) WithError(err error) Logger {
	return &logrusAdapter{entry: l.entry.WithField("error", err.Error())}
}

var (
	baseLogger  = logrus.New()
	loggerCache = map[string]Logger{}
	mu          sync.Mutex
)

func init() {
	baseLogger.SetOutput(os.Stdout)
	baseLogger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	baseLogger.SetLevel(logrus.InfoLevel)
}

func Configure(config *config.Log) {
	parsed, err := logrus.ParseLevel(config.Level)
	if err != nil {
		baseLogger.Warnf("Invalid log level '%s', defaulting to 'info'", config.Level)
		parsed = logrus.InfoLevel
	}
	baseLogger.SetLevel(parsed)
}

func GetLogger(tag string) Logger {
	mu.Lock()
	defer mu.Unlock()

	if logger, ok := loggerCache[tag]; ok {
		return logger
	}

	entry := baseLogger.WithField("component", tag)
	wrapped := &logrusAdapter{entry: entry}
	loggerCache[tag] = wrapped
	return wrapped
}
