package logger

import (
	"os"
	"server/internal/apperrors"
	"server/internal/config"
	"strings"

	"github.com/sirupsen/logrus"
)

type ILogger interface {
	DebugFmt(message string, function string, routeNode string)
	Debug(message string)
	Info(message string)
	Error(message string)
	Fatal(message string)
	Printf(message string, args ...interface{})
	Level() int
}

type LogrusLogger struct {
	level  int
	logger *logrus.Logger
}

func NewLogrusLogger(lc *config.LoggingConfig) (LogrusLogger, error) {
	logLevel, err := logrus.ParseLevel(lc.Level)
	if err != nil {
		return LogrusLogger{}, apperrors.ErrInvalidLoggingLevel
	}

	logger := logrus.Logger{
		Out:   os.Stderr,
		Level: logLevel,
		ReportCaller: (lc.LevelBasedReport && logLevel == logrus.TraceLevel) ||
			(!lc.LevelBasedReport && lc.ReportCaller),
		Formatter: &logrus.TextFormatter{
			DisableTimestamp:       lc.DisableTimestamp,
			FullTimestamp:          lc.FullTimestamp,
			DisableLevelTruncation: lc.DisableLevelTruncation,
		},
	}

	return LogrusLogger{
		logger: &logger,
		level:  int(logger.Level),
	}, nil
}

func (l *LogrusLogger) DebugFmt(message string, function string, routeNode string) {
	tabsNum := 0
	switch routeNode {
	case "service":
		tabsNum = 1
	case "storage":
		tabsNum = 2
	}
	l.logger.
		WithFields(logrus.Fields{
			"route_node": routeNode,
			"function":   function,
		}).
		Debug(strings.Repeat("\t", tabsNum) + message)
}

func (l *LogrusLogger) Debug(message string) {
	l.logger.Debug(message)
}

func (l *LogrusLogger) Info(message string) {
	l.logger.Info(message)
}

func (l *LogrusLogger) Error(message string) {
	l.logger.Error(message)
}

func (l *LogrusLogger) Fatal(message string) {
	l.logger.Fatal(message)
}

func (l *LogrusLogger) Printf(message string, args ...interface{}) {
	l.logger.Printf(message)
}

func (l *LogrusLogger) Level() int {
	return l.level
}
