package loggerStruct

import (
	"fmt"
	"strings"
	"time"
)

type Logger struct {
	AppName     string    `json:"appName"`
	Level       string    `json:"level"`
	Description string    `json:"description"`
	Trace       string    `json:"trace"`
	Timestamp   time.Time `json:"timestamp"`
}

func (logger *Logger) ValidateName() error {
	level := strings.ToLower(logger.Level)
	if level != "error" && level != "warning" && level != "info" && level != "unknown" {
		return fmt.Errorf("invalid level value: %s\nlevel can only have four possible values:\n1. error\n2. warning\n3. info\n4. unknown", logger.Level)
	}
	logger.Level = level
	logger.Timestamp = time.Now()
	fmt.Println(logger)
	return nil
}
