package common

import (
	"context"
	"encoding/json"
	"io"

	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

// EchoLogrusLogger extend logrus.Logger
type EchoLogrusLogger struct {
	*logrus.Logger
	Ctx    context.Context
	Fields logrus.Fields
}

var commonLogger = &EchoLogrusLogger{
	Logger: logrus.StandardLogger(),
	Ctx:    context.Background(),
	Fields: logrus.Fields{},
}

func Logger() *EchoLogrusLogger {
	return commonLogger
}

func toEchoLevel(level logrus.Level) log.Lvl {
	switch level {
	case logrus.DebugLevel:
		return log.DEBUG
	case logrus.InfoLevel:
		return log.INFO
	case logrus.WarnLevel:
		return log.WARN
	case logrus.ErrorLevel:
		return log.ERROR
	}

	return log.OFF
}

func (l *EchoLogrusLogger) Output() io.Writer {
	return l.Out
}

func (l *EchoLogrusLogger) SetOutput(w io.Writer) {
	// disable operations that would change behavior of global logrus logger.
}

func (l *EchoLogrusLogger) Level() log.Lvl {
	return toEchoLevel(l.Logger.Level)
}

func (l *EchoLogrusLogger) SetLevel(v log.Lvl) {
	// disable operations that would change behavior of global logrus logger.
}

func (l *EchoLogrusLogger) SetHeader(h string) {
}

func (l *EchoLogrusLogger) Prefix() string {
	return ""
}

func (l *EchoLogrusLogger) SetPrefix(p string) {
}

func (l *EchoLogrusLogger) Print(i ...interface{}) {
	l.Logger.WithContext(l.Ctx).WithFields(l.Fields).Print(i...)
}

func (l *EchoLogrusLogger) Printf(format string, args ...interface{}) {
	l.Logger.WithContext(l.Ctx).WithFields(l.Fields).Printf(format, args...)
}

func (l *EchoLogrusLogger) Printj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.WithContext(l.Ctx).WithFields(l.Fields).Println(string(b))
}

func (l *EchoLogrusLogger) Debug(i ...interface{}) {
	l.Logger.WithContext(l.Ctx).WithFields(l.Fields).Debug(i...)
}

func (l *EchoLogrusLogger) Debugf(format string, args ...interface{}) {
	l.Logger.WithContext(l.Ctx).WithFields(l.Fields).Debugf(format, args...)
}

func (l *EchoLogrusLogger) Debugj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.WithContext(l.Ctx).WithFields(l.Fields).Debugln(string(b))
}

func (l *EchoLogrusLogger) Info(i ...interface{}) {
	l.Logger.WithContext(l.Ctx).WithFields(l.Fields).Info(i...)
}

func (l *EchoLogrusLogger) Infof(format string, args ...interface{}) {
	l.Logger.WithContext(l.Ctx).WithFields(l.Fields).Infof(format, args...)
}

func (l *EchoLogrusLogger) Infoj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.WithContext(l.Ctx).WithFields(l.Fields).Infoln(string(b))
}

func (l *EchoLogrusLogger) Warn(i ...interface{}) {
	l.Logger.WithContext(l.Ctx).WithFields(l.Fields).Warn(i...)
}

func (l *EchoLogrusLogger) Warnf(format string, args ...interface{}) {
	l.Logger.WithContext(l.Ctx).WithFields(l.Fields).Warnf(format, args...)
}

func (l *EchoLogrusLogger) Warnj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.WithContext(l.Ctx).WithFields(l.Fields).Warnln(string(b))
}

func (l *EchoLogrusLogger) Error(i ...interface{}) {
	l.Logger.WithContext(l.Ctx).WithFields(l.Fields).Error(i...)
}

func (l *EchoLogrusLogger) Errorf(format string, args ...interface{}) {
	l.Logger.WithContext(l.Ctx).WithFields(l.Fields).Errorf(format, args...)
}

func (l *EchoLogrusLogger) Errorj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.WithContext(l.Ctx).WithFields(l.Fields).Errorln(string(b))
}

func (l *EchoLogrusLogger) Fatal(i ...interface{}) {
	l.Logger.WithContext(l.Ctx).WithFields(l.Fields).Fatal(i...)
}

func (l *EchoLogrusLogger) Fatalf(format string, args ...interface{}) {
	l.Logger.WithContext(l.Ctx).WithFields(l.Fields).Fatalf(format, args...)
}

func (l *EchoLogrusLogger) Fatalj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.WithContext(l.Ctx).WithFields(l.Fields).Fatalln(string(b))
}

func (l *EchoLogrusLogger) Panic(i ...interface{}) {
	l.Logger.WithContext(l.Ctx).WithFields(l.Fields).Panic(i...)
}

func (l *EchoLogrusLogger) Panicf(format string, args ...interface{}) {
	l.Logger.WithContext(l.Ctx).WithFields(l.Fields).Panicf(format, args...)
}

func (l *EchoLogrusLogger) Panicj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.WithContext(l.Ctx).WithFields(l.Fields).Panicln(string(b))
}
