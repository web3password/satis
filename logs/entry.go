/*
Copyright (C) 2024 Web3Password PTE. LTD.(Singapore UEN: 202333030C) - All Rights Reserved

Web3Password PTE. LTD.(Singapore UEN: 202333030C) holds the copyright of this file.

Unauthorized copying or redistribution of this file in binary forms via any medium is strictly prohibited.

For more information, please refer to https://www.web3password.com/web3password_license.txt
*/

package logs

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"runtime"
	"strings"
)

const (
	JsonFormatter = "json"
	TextFormatter = "text"

	errorKey = "error"
)

var (
	errUnknowLevel = fmt.Errorf("got unknown logger level")

	levels = map[Level]logrus.Level{
		PanicLevel: logrus.PanicLevel,
		FatalLevel: logrus.FatalLevel,
		ErrorLevel: logrus.ErrorLevel,
		WarnLevel:  logrus.WarnLevel,
		InfoLevel:  logrus.InfoLevel,
		DebugLevel: logrus.DebugLevel,
		TraceLevel: logrus.TraceLevel,
	}
)

type log struct {
	entry        *logrus.Entry
	depth        int
	reportCaller bool
}

type Option struct {
	Output           io.Writer
	Level            Level
	Formatter        string
	EnableHTMLEscape bool
	ReportCaller     bool
}

var defaultOption = &Option{
	Output:           os.Stdout,
	Level:            InfoLevel,
	Formatter:        JsonFormatter,
	EnableHTMLEscape: false,
	ReportCaller:     true,
}

func New() Logger {
	return NewWithOption(defaultOption)
}

func NewWithOption(option *Option) Logger {
	if option == nil {
		option = defaultOption
	}

	logger := logrus.New()

	// set level
	driverLevel, exists := levels[option.Level]
	if exists {
		logger.SetLevel(driverLevel)
	}

	// set formatter
	if option.Formatter == JsonFormatter {
		logger.SetFormatter(&logrus.JSONFormatter{DisableHTMLEscape: !option.EnableHTMLEscape})
	}

	// set output
	if option.Output != nil {
		logger.SetOutput(option.Output)
	} else {
		logger.SetOutput(os.Stdout)
	}

	// set no lock
	logger.SetNoLock()

	return &log{
		entry:        logrus.NewEntry(logger),
		depth:        1,
		reportCaller: option.ReportCaller,
	}
}

func (l *log) log(level logrus.Level, args ...interface{}) {
	entry := l.entry
	// report caller
	if l.reportCaller {
		entry = l.entry.WithField("file", caller(l.depth+3))
	}
	entry.Log(level, args...)

	if metricsFunc != nil {
		levelStr := ""
		if level == logrus.ErrorLevel {
			levelStr = "ERROR"
		} else if level == logrus.FatalLevel {
			levelStr = "FATAL"
		}
		if levelStr != "" {
			metricsFunc(levelStr)
		}
	}
}

func (l *log) Log(level Level, args ...interface{}) {
	driverLevel, exists := levels[level]
	if !exists {
		l.log(logrus.WarnLevel, errUnknowLevel)
		return
	}
	if l.entry.Logger.IsLevelEnabled(driverLevel) {
		l.log(driverLevel, args...)
	}
}

func (l *log) Logf(level Level, format string, args ...interface{}) {
	driverLevel, exists := levels[level]
	if !exists {
		l.log(logrus.WarnLevel, errUnknowLevel)
		return
	}
	if l.entry.Logger.IsLevelEnabled(driverLevel) {
		l.log(driverLevel, fmt.Sprintf(format, args...))
	}
}

func (l *log) Trace(args ...interface{}) {
	l.Log(TraceLevel, args...)
}

func (l *log) Tracef(format string, args ...interface{}) {
	l.Logf(TraceLevel, format, args...)
}

func (l *log) Debug(args ...interface{}) {
	l.Log(DebugLevel, args...)
}

func (l *log) Debugf(format string, args ...interface{}) {
	l.Logf(DebugLevel, format, args...)
}

func (l *log) Info(args ...interface{}) {
	l.Log(InfoLevel, args...)
}

func (l *log) Errorf(format string, args ...interface{}) {
	l.Logf(ErrorLevel, format, args...)
}

func (l *log) Fatal(args ...interface{}) {
	l.Log(FatalLevel, args...)
	l.entry.Logger.Exit(1)
}
func (l *log) Infof(format string, args ...interface{}) {
	l.Logf(InfoLevel, format, args...)
}

func (l *log) Warnf(format string, args ...interface{}) {
	l.Logf(WarnLevel, format, args...)
}
func (l *log) Warn(args ...interface{}) {
	l.Log(WarnLevel, args...)
}

func (l *log) Error(args ...interface{}) {
	l.Log(ErrorLevel, args...)
}

func (l *log) Fatalf(format string, args ...interface{}) {
	l.Logf(FatalLevel, format, args...)
	l.entry.Logger.Exit(1)
}

func (l *log) Panic(args ...interface{}) {
	l.Log(PanicLevel, args...)
	panic(fmt.Sprint(args...))
}

func (l *log) Panicf(format string, args ...interface{}) {
	l.Logf(PanicLevel, format, args...)
	panic(fmt.Sprintf(format, args...))
}

func (l *log) Print(args ...interface{}) {
	l.Log(InfoLevel, args...)
}

func (l *log) Printf(format string, args ...interface{}) {
	l.Logf(InfoLevel, format, args...)
}

func (l *log) WithFields(fields Fields) Logger {
	if l.reportCaller {
		if err, ok := fields[errorKey].(interface {
			Stack() []string
		}); ok {
			fields["err.stack"] = strings.Join(err.Stack(), ";")
		}
	}
	return &log{entry: l.entry.WithFields(logrus.Fields(fields)), reportCaller: l.reportCaller}
}
func (l *log) WithField(key string, value interface{}) Logger {
	return l.WithFields(Fields{key: value})
}

func (l *log) WithError(err error) Logger {
	return l.WithFields(Fields{errorKey: err})
}

func (l *log) SetOutput(output io.Writer) {
	l.entry.Logger.SetOutput(output)
}

func (l *log) GetOutput() io.Writer {
	return l.entry.Logger.Out
}
func (l *log) SetLevel(level Level) {
	driverLevel, exists := levels[level]
	if !exists {
		l.log(logrus.WarnLevel, errUnknowLevel)
	}
	l.entry.Logger.SetLevel(driverLevel)
}

func (l *log) GetLevel() Level {
	return Level(l.entry.Logger.GetLevel())
}

func caller(depth int) string {
	_, f, n, ok := runtime.Caller(1 + depth)
	if !ok {
		return ""
	}
	if ok {
		idx := strings.LastIndex(f, "git.web3password.com")
		if idx >= 0 {
			f = f[idx+18:]
		}
	}
	return fmt.Sprintf("%s:%d", f, n)
}
