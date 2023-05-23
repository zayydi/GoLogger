package logmon

import (
	"loggerStruct"

	"github.com/fatih/color"
	"github.com/gologger/eslogger"
)

var APP_NAME string

func SetAppName(appName string) {
	APP_NAME = appName
}

func Info(description string) {
	log := loggerStruct.Logger{
		AppName:     APP_NAME,
		Level:       "info",
		Description: description,
	}
	eslogger.LogError(log)
	color.New(color.FgCyan).Println(description)
}

func Warning(description string, trace string) {
	log := loggerStruct.Logger{
		AppName:     APP_NAME,
		Level:       "warning",
		Description: description,
		Trace:       trace,
	}
	eslogger.LogError(log)
	color.New(color.FgYellow).Println(description)
}

func Error(description string, trace string) {
	log := loggerStruct.Logger{
		AppName:     APP_NAME,
		Level:       "error",
		Description: description,
		Trace:       trace,
	}
	eslogger.LogError(log)
	color.New(color.FgRed).Println(description)
}

func Log(level string, description string, trace string) {
	if level != "error" && level != "warning" && level != "info" {
		level = "unknown"
	}

	log := loggerStruct.Logger{
		AppName:     APP_NAME,
		Level:       level,
		Description: description,
		Trace:       trace,
	}

	eslogger.LogError(log)

	if level == "info" {
		color.New(color.FgCyan).Println(description)
	} else if level == "warning" {
		color.New(color.FgYellow).Println(description)
	} else if level == "error" {
		color.New(color.FgRed).Println(description)
	} else {
		color.New(color.FgWhite).Println(description)
	}
}
