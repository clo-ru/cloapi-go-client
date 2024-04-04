package clo

import (
	"log"
	"os"
)

type Logger interface {
	Trace(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Panic(...interface{})
	Traceln(...interface{})
	Infoln(...interface{})
	Warnln(...interface{})
	Errorln(...interface{})
	Panicln(...interface{})
	Tracef(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	Panicf(string, ...interface{})
}

func NewDefaultLogger(prefix string) Logger {
	return defaultLogger{
		logger: log.New(os.Stdout, prefix, log.LstdFlags),
	}
}

type defaultLogger struct {
	logger *log.Logger
}

func (l defaultLogger) Trace(v ...interface{}) {
	l.logger.Print(v...)
}
func (l defaultLogger) Info(v ...interface{}) {
	l.logger.Print(v...)
}
func (l defaultLogger) Warn(v ...interface{}) {
	l.logger.Print(v...)
}
func (l defaultLogger) Error(v ...interface{}) {
	l.logger.Print(v...)
}
func (l defaultLogger) Panic(v ...interface{}) {
	l.logger.Panic(v...)
}
func (l defaultLogger) Traceln(v ...interface{}) {
	l.logger.Println(v...)
}
func (l defaultLogger) Infoln(v ...interface{}) {
	l.logger.Println(v...)
}
func (l defaultLogger) Warnln(v ...interface{}) {
	l.logger.Println(v...)
}
func (l defaultLogger) Errorln(v ...interface{}) {
	l.logger.Println(v...)
}
func (l defaultLogger) Panicln(v ...interface{}) {
	l.logger.Panic(v...)
}
func (l defaultLogger) Tracef(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}
func (l defaultLogger) Infof(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}
func (l defaultLogger) Warnf(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}
func (l defaultLogger) Errorf(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}
func (l defaultLogger) Panicf(format string, v ...interface{}) {
	l.logger.Panicf(format, v...)
}
