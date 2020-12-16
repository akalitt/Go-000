package logger

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github/akalitt/go-errors-example/pkg/file"
	"github/akalitt/go-errors-example/pkg/header"
	"io"
	"os"
	"strings"
	"time"
)

// std represents the logger which outputs to the stdout.
var std = logrus.New()

// Init initializes the global logger.
func Init(level string) {

	var stdFmt logrus.Formatter

	// Create the formatter for stdout output.
	stdFmt = &formatter{
		isStdout:    false,
		serviceName: strings.ToUpper(viper.GetString("db.project")),
	}

	// Std logger.
	LOG_PATH := fmt.Sprintf("%s%s",
		viper.Get("runtimeRootPath"),
		viper.Get("logSavePath"),
	)

	LOG_NAME := fmt.Sprintf("%s%s.%s",
		"log-",
		time.Now().Format("20060102"),
		"log",
	)

	f, err := file.MustOpen(LOG_NAME, LOG_PATH)
	if err != nil {
		fmt.Println("err", err)
	}

	std.Out = io.MultiWriter(f, os.Stdout)
	std.Formatter = stdFmt

	switch strings.ToUpper(level) {
	case "FATAL":
		std.Level = logrus.FatalLevel
	case "ERROR":
		std.Level = logrus.ErrorLevel
	case "WARN":
		std.Level = logrus.WarnLevel
	case "INFO":
		std.Level = logrus.InfoLevel
	case "DEBUG":
		std.Level = logrus.DebugLevel
	default:
		std.Level = logrus.WarnLevel
	}

}

// formatter formats the output format.
type formatter struct {
	isStdout    bool
	serviceName string
}

// Format the input log.
func (f *formatter) Format(e *logrus.Entry) ([]byte, error) {
	// Implode the data to string with `k=v` format.

	dataString := ""
	if len(e.Data) != 0 {
		for k, v := range e.Data {
			dataString += fmt.Sprintf("%s=%+v ", k, v)
		}
		// Trim the trailing whitespace.
		dataString = dataString[0 : len(dataString)-1]
	}
	// Get service name.
	name := f.serviceName
	// Level like: DEBUG, INFO, WARN, ERROR, FATAL.
	level := strings.ToUpper(e.Level.String())

	// Get the time with YYYY-mm-dd H:i:s format.
	time := e.Time.Format("2006-01-02 15:04:05")
	// Get the message.
	msg := e.Message

	stdLevel := ""
	switch level {
	case "DEBUG":
		stdLevel = color.New(color.FgWhite).Sprint(level)
	case "INFO":
		stdLevel = color.New(color.FgCyan).Sprint(" " + level)
	case "WARN":
		stdLevel = color.New(color.FgYellow).Sprint(" " + level)
	case "ERROR":
		stdLevel = color.New(color.FgRed).Sprint(level)
	case "FATAL":
		stdLevel = color.New(color.FgHiRed).Sprint(level)
	}

	body := fmt.Sprintf("[%s] %5s %s %s %s \n", name, stdLevel, time, header.REQUEST_ID, msg)
	data := fmt.Sprintf(" [(%s)]", dataString)

	// Hide the data if there's no data.
	if len(e.Data) == 0 {
		data = ""
	}

	// Mix the body and the data.
	output := fmt.Sprintf("%s%s\n %s\n", body, data,
		"--------------------------------")

	return []byte(output), nil
}

func Debug(msg interface{}) {
	message("Debug", msg)
}
func Info(msg interface{}) {
	message("Info", msg)
}
func Warn(msg interface{}) {
	message("Warn", msg)
}
func Error(msg interface{}) {
	message("Error", fmt.Sprintf("%+v", msg))
}
func Fatal(msg interface{}) {
	message("Fatal", msg)
}

func Debugf(msg string, fds logrus.Fields) {
	fields("Debug", msg, fds)
}
func Infof(msg string, fds logrus.Fields) {
	fields("Info", msg, fds)
}
func Warnf(msg string, fds logrus.Fields) {
	fields("Warn", msg, fds)
}
func Errorf(msg interface{}, fds logrus.Fields) {
	fields("Error", fmt.Sprintf("%+v \n", msg), fds)
}
func Fatalf(msg string, fds logrus.Fields) {
	fields("Fatal", msg, fds)
}

func fields(lvl string, msg string, fds logrus.Fields) {
	s := std.WithFields(fds)

	switch lvl {
	case "Debug":
		s.Debug(msg)
	case "Info":
		s.Info(msg)
	case "Warn":
		s.Warn(msg)
	case "Error":
		s.Error(msg)
	case "Fatal":
		s.Fatal(msg)
	}
}

func message(lvl string, msg interface{}) {
	switch lvl {
	case "Debug":
		std.Debug(msg)
	case "Info":
		std.Info(msg)
	case "Warn":
		std.Warn(msg)
	case "Error":
		std.Error(msg)
	case "Fatal":
		std.Fatal(msg)
	}
}
